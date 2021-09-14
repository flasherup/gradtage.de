package impl

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/config"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/impl/database"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/impl/parser"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/impl/utils"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"net/http"
	"sync"

	"time"
)

type UpdateResult struct {
	StId    string
	Error   error
	Message string
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
	resultChanel          chan UpdateResult
	updateCounter         *ktprom.Counter
	waitForRequests       *sync.WaitGroup
}

const (
	labelStation = "station"
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
		resultChanel:          make(chan UpdateResult),
		waitForRequests:       &sync.WaitGroup{},
	}

	wb.updateCounter = ktprom.NewCounterFrom(
		prometheus.CounterOpts(
			prometheus.Opts{
				Name: "weatherbit_update_counter",
				Help: "The total number of success/error updates",
			},
		),
		[]string{labelStation, labelStatus},
	)

	go handleUpdates(&wb)
	go startFetchProcess(&wb)
	return &wb, nil
}

func (wbu WeatherBitUpdateSVC) ForceRestart(ctx context.Context) error {
	return nil
}

func handleUpdates(wbu *WeatherBitUpdateSVC) {
	for {
		select {
		case updateResult := <-wbu.resultChanel:
			if updateResult.Error != nil {
				wbu.updateCounter.With(labelStation, updateResult.StId, labelStatus, "error").Add(1)
				level.Error(wbu.logger).Log("msg", updateResult.Message, "err", updateResult.Error)
			} else {
				wbu.updateCounter.With(labelStation, updateResult.StId, labelStatus, "success").Add(1)
			}
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
	err := wbu.db.CreateTable(stID)
	if err != nil {
		wbu.resultChanel <- UpdateResult{
			StId:    stID,
			Message: fmt.Sprintf("database table '%s' create error", stID),
			Error:   err,
		}
		return
	}

	url := wbu.conf.Weatherbit.UrlWeatherBit + "/history/hourly?station=" + st + "&key=" + wbu.conf.Weatherbit.KeyWeatherBit + "&start_date=" + start + "&end_date=" + end
	level.Info(wbu.logger).Log("msg", "weather bit request", "url", url)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		wbu.resultChanel <- UpdateResult{
			StId:    stID,
			Message: fmt.Sprintf("request create error for url '%s' ", url),
			Error:   err,
		}
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		wbu.resultChanel <- UpdateResult{
			StId:    stID,
			Message: fmt.Sprintf("request error for url '%s' ", url),
			Error:   err,
		}
		return
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		wbu.resultChanel <- UpdateResult{
			StId:    stID,
			Message: fmt.Sprintf("response read error for url '%s' ", url),
			Error:   err,
		}
		return
	}

	err = resp.Body.Close()
	if err != nil {
		wbu.resultChanel <- UpdateResult{
			StId:    stID,
			Message: fmt.Sprintf("response body close error for url '%s' ", url),
			Error:   err,
		}
	}

	result, err := parser.ParseWeatherBit(&contents)
	if err != nil {
		wbu.resultChanel <- UpdateResult{
			StId:    stID,
			Message: fmt.Sprintf("weather bit data parse error for url '%s' ", url),
			Error:   err,
		}
		return
	}

	err = wbu.db.PushData(stID, result)
	if err != nil {
		wbu.resultChanel <- UpdateResult{
			StId:    stID,
			Message: fmt.Sprintf("data push error for url '%s' ", url),
			Error:   err,
		}
		return
	}

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
