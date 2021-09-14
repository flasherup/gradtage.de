package impl

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/config"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/impl/database"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/impl/utils"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"sync"

	"time"
)

type UpdateStart struct {
	StId string
}

type UpdateResult struct {
	StId    string
	Error   error
	Message string
}

type metrics struct {
	updateCounter   *ktprom.Counter
	requestsCounter *ktprom.Counter
}

type WeatherBitUpdateSVC struct {
	stations              stationssvc.Client
	db                    database.WeatherBitDB
	alert                 alertsvc.Client
	logger                log.Logger
	conf                  config.WeatherBitUpdateConfig
	dailyRequestCounter   int
	secondsRequestCounter int
	dailyStartTime        time.Time
	secondsStartTime      time.Time
	startChanel           chan UpdateStart
	resultChanel          chan UpdateResult
	metrics               metrics
}

const (
	labelStation  = "station"
	labelStatus  = "status"
)

func NewWeatherBitUpdateSVC(
	logger log.Logger,
	stations stationssvc.Client,
	db database.WeatherBitDB,
	alert alertsvc.Client,
	conf config.WeatherBitUpdateConfig,
) (*WeatherBitUpdateSVC, error) {
	wb := WeatherBitUpdateSVC{
		stations:              stations,
		db:                    db,
		alert:                 alert,
		logger:                logger,
		conf:                  conf,
		dailyRequestCounter:   0,
		secondsRequestCounter: 0,
		startChanel:           make(chan UpdateStart),
		resultChanel:          make(chan UpdateResult),
	}

	wb.metrics = *setupMetrics()

	go handleUpdates(&wb)
	go startFetchProcess(&wb)
	return &wb, nil
}

func setupMetrics() *metrics {
	updateCounter := ktprom.NewCounterFrom(
		prometheus.CounterOpts(
			prometheus.Opts{
				Name: "weatherbit_update_counter",
				Help: "The total number of success/error updates",
			},
		),
		[]string{labelStatus},
	)

	requestCounter := ktprom.NewCounterFrom(
		prometheus.CounterOpts(
			prometheus.Opts{
				Name: "weatherbit_request_counter",
				Help: "The total number of requests",
			},
		),
		[]string{labelStation},
	)

	return &metrics{
		updateCounter: updateCounter,
		requestsCounter: requestCounter,
	}
}

func (wbu WeatherBitUpdateSVC) ForceRestart(ctx context.Context) error {
	return nil
}

func handleUpdates(wbu *WeatherBitUpdateSVC) {
	for {
		select {
		case updateResult := <-wbu.resultChanel:
			if updateResult.Error != nil {
				wbu.metrics.updateCounter.With(labelStatus, "error").Add(1)
				level.Error(wbu.logger).Log("msg", updateResult.Message, "err", updateResult.Error)
			} else {
				wbu.metrics.updateCounter.With(labelStatus, "success").Add(1)
			}
		case updateStart := <-wbu.startChanel:
			wbu.metrics.requestsCounter.With(labelStation, updateStart.StId).Add(1)
		}
	}
}

func startFetchProcess(wbu *WeatherBitUpdateSVC) {
	wg := sync.WaitGroup{}
	for {
		date := time.Now()
		resp, err := wbu.stations.GetAllStations()
		if err != nil {
			level.Error(wbu.logger).Log("msg", "Cant get stations list", "error", err)
			time.Sleep(time.Minute)
			continue
		}

		sts := make(map[string]string, len(resp.Sts))
		for _, v := range resp.Sts {
			sts[v.Id] = v.SourceId
		}

		level.Info(wbu.logger).Log("msg", "Start update process", "date", date)
		wbu.precessStations(date, sts, &wg)
		wg.Wait()
	}
}

func (wbu *WeatherBitUpdateSVC) precessStations(date time.Time, sts map[string]string, wg *sync.WaitGroup) {
	wbu.dailyStartTime = date
	wbu.secondsStartTime = date
	for k, v := range sts {
		level.Info(wbu.logger).Log("msg", "Process station", "innerId", k, "station", v)
		wbu.processRequest(k, v, date, wg)
	}
}

func (wbu *WeatherBitUpdateSVC) processRequest(stID string, st string, end time.Time, wg *sync.WaitGroup) {
	endDate := end
	requestsPerSecond := wbu.conf.Weatherbit.NumberOfRequestPerSecond
	requestsPerDay := wbu.conf.Weatherbit.NumberOfRequestPerDay
	for {
		startDate := endDate.AddDate(0, 0, -wbu.conf.Weatherbit.NumberOfDaysPerRequest)
		daysDif := utils.DaysDifference(startDate, end)
		if daysDif > wbu.conf.Weatherbit.NumberOfDays {
			break
		}
		sDate := startDate.Format(common.TimeLayoutWBH)
		eDate := endDate.Format(common.TimeLayoutWBH)
		/*if utils.DaysCheck(startDate, end, wbu.conf.Weatherbit.NumberOfDays) {
			break
		}*/
		endDate = startDate
		if wbu.checkPeriod(stID, sDate, eDate) {
			continue
		}
		wg.Add(1)
		go wbu.processUpdate(stID, st, sDate, eDate, wg)
		wbu.secondsRequestCounter, wbu.secondsStartTime = utils.SleepCheck(requestsPerSecond, wbu.secondsRequestCounter, wbu.secondsStartTime, time.Second)
		wbu.dailyRequestCounter, wbu.dailyStartTime = utils.SleepCheck(requestsPerDay, wbu.dailyRequestCounter, wbu.dailyStartTime, time.Hour*24)
	}
}

func (wbu *WeatherBitUpdateSVC) processUpdate(stID, st, start, end string, wg *sync.WaitGroup) {
	defer wg.Done()

	wbu.startChanel <- UpdateStart{
		StId: stID,
	}
	url := wbu.conf.Weatherbit.UrlWeatherBit + "/history/hourly?station=" + st + "&key=" + wbu.conf.Weatherbit.KeyWeatherBit + "&start_date=" + start + "&end_date=" + end

	err := wbu.db.CreateTable(stID)
	if err != nil {
		wbu.resultChanel <- UpdateResult{
			StId:    stID,
			Message: fmt.Sprintf("database table '%s' create error", stID),
			Error:   err,
		}
		return
	}

	data, message, requestError := utils.MakeWeatherbitRequest(url)
	if requestError != nil {
		wbu.resultChanel <- UpdateResult{
			StId:    stID,
			Message: message,
			Error:   err,
		}
		return
	}

	err = wbu.db.PushData(stID, data)
	if err != nil {
		wbu.resultChanel <- UpdateResult{
			StId:    stID,
			Message: fmt.Sprintf("data push error for url '%s' ", url),
			Error:   err,
		}
		return
	}
	//time.Sleep(time.Second * 30)

	wbu.resultChanel <- UpdateResult{
		StId:    stID,
		Message: fmt.Sprintf("data update success '%s' ", url),
		Error:   nil,
	}
}

func (wbu WeatherBitUpdateSVC) checkPeriod(stID string, start string, end string) bool {
	temps, err := wbu.db.GetPeriod(stID, start, end)
	if err == nil {
		return len(temps) != 0
	}
	return false
}
