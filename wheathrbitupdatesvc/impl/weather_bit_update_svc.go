package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/wheathrbitupdatesvc/config"
	"github.com/flasherup/gradtage.de/wheathrbitupdatesvc/impl/database"
	"github.com/flasherup/gradtage.de/wheathrbitupdatesvc/impl/parser"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"

	"net/http"
	"time"
)

type WeatherBitUpdateSVC struct {
	stations    stationssvc.Client
	db 			database.WeatherBitDB
	alert 		alertsvc.Client
	logger  	log.Logger
	conf		config.WeatherBitUpdateConfig
	stationList           map[string]string
	dailyRequestCounter   int
	secondsRequestCounter int
	dailyStartTime        time.Time
	secondsStartTime      time.Time
}



func NewWeatherBitUpdateSVC(
	logger 		log.Logger,
	stations 	stationssvc.Client,
	db 			database.WeatherBitDB,
	alert 		alertsvc.Client,
	conf 		config.WeatherBitUpdateConfig,
) (*WeatherBitUpdateSVC, error) {
	wb := WeatherBitUpdateSVC {
		stations:stations,
		db:db,
		alert:alert,
		logger:logger,
		conf:conf,
		dailyRequestCounter:   0,
		secondsRequestCounter: 0,
	}



	//processUpdate(wb, startDate, endDate)

	//go startFetchProcess(&wb)
	return &wb,nil
}

func (wb WeatherBitUpdateSVC) ForceRestart(ctx context.Context) error {
	return nil
}

func startFetchProcess(wb *WeatherBitUpdateSVC) {
	wb.precessStations() //Do it first time
	tick := time.Tick(time.Hour * 24)
	for {
		select {
		case <-tick:
			wb.precessStations()
		}
	}
}


func (wbh *WeatherBitUpdateSVC)precessStations(date time.Time) {
	wbh.dailyStartTime = date
	wbh.secondsStartTime = date
	for k,v := range wbh.stationList {
		level.Info(wbh.logger).Log("msg", "Process station", "innerId", k, "station", v)
		wbh.processRequest(k, v, date)
	}
}

func (wbh *WeatherBitUpdateSVC)processRequest(stID string, st string, end time.Time) error {
	startDate := end
	requestsPerSecond := wbh.conf.Weatherbit.NumberOfRequestPerSecond
	requestsPerDay := wbh.conf.Weatherbit.NumberOfRequestPerDay
	for {
		start := end.AddDate(0, 0, -14)
		sDate := start.Format(common.TimeLayoutWBH)
		eDate := end.Format(common.TimeLayoutWBH)
		if daysCheck(startDate, end, wbh.conf.Weatherbit.NumberOfDays) {
			break
		}
		end = start
		if wbh.checkPeriod(stID, sDate, eDate) {
			continue
		}
		go wbh.processUpdate(stID, st, sDate, eDate)
		wbh.secondsRequestCounter, wbh.secondsStartTime = sleepCheck(requestsPerSecond, wbh.secondsRequestCounter, wbh.secondsStartTime, time.Second)
		wbh.dailyRequestCounter, wbh.dailyStartTime 	= sleepCheck(requestsPerDay, wbh.dailyRequestCounter, wbh.dailyStartTime, time.Hour * 24)
	}
	return nil
}

func (wbh WeatherBitUpdateSVC)processUpdate(stID string, st string, start string, end string) error {
	err := wbh.db.CreateTable(stID)
	if err != nil {
		level.Error(wbh.logger).Log("msg", "table create error", "err", err)
		return err
	}

	url := wbh.conf.Weatherbit.UrlWeatherBit + "/history/hourly?station=" + st + "&key=" + wbh.conf.Weatherbit.KeyWeatherBit + "&start_date=" + start + "&end_date=" + end
	level.Info(wbh.logger).Log("msg", "weather bit request", "url", url)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		level.Error(wbh.logger).Log("msg", "request error", "err", err, "url", url)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		level.Error(wbh.logger).Log("msg", "request error", "err", err, "url", url)
		return err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		level.Error(wbh.logger).Log("msg", "response read error", "err", err, "url", url)
		return err
	}

	result, err := parser.ParseWeatherBit(&contents)
	if err != nil {
		level.Error(wbh.logger).Log("msg", "weather bit data parse error", "err", err, "url", url)
		return err
	}

	err = wbh.db.PushData(stID, result)
	if err != nil {
		level.Error(wbh.logger).Log("msg", "data push error", "err", err, "url", url)
		return err
	}
	return nil
}

func (wbh WeatherBitUpdateSVC)checkPeriod(stID string, start string, end string) bool {
	temps, err := wbh.db.GetPeriod(stID, start, end)
	if err == nil {
		return len(temps) != 0
	}
	return false
}

func sleepCheck(numberOfRequests, counter int, startTime time.Time, duration time.Duration) (int, time.Time) {
	now := time.Now()
	dif := now.Sub(startTime)
	if dif >= duration {
		return 0, now
	}

	counter++
	if counter >= numberOfRequests{
		time.Sleep(duration - dif)
		return 0, time.Now()
	}
	return counter, startTime
}

func daysCheck(start, end time.Time, daysCount int) bool {
	return 	start.Year() - end.Year() >= daysCount &&
		start.Month() > end.Month()
}