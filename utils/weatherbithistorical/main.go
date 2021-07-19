package main

import (
	"flag"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/utils/weatherbithistorical/config"
	"github.com/flasherup/gradtage.de/utils/weatherbithistorical/csv"
	"github.com/flasherup/gradtage.de/utils/weatherbithistorical/database"
	"github.com/flasherup/gradtage.de/utils/weatherbithistorical/parser"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type WeatherHistorical struct {
	db                    database.WeatherBitDB
	logger                log.Logger
	conf                  config.WeatherBitConfig
	stationList           map[string]string
	dailyRequestCounter   int
	secondsRequestCounter int
	dailyStartTime        time.Time
	secondsStartTime      time.Time
}

func main() {
	configFile := flag.String("config.file", "config.yml", "Config file name. ")
	csvFile := flag.String("csv.file", "stationsR.csv", "CSV file name. ")
	flag.Parse()

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowDebug())

	//Config
	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		level.Error(logger).Log("msg", "config loading error", "err", err.Error())
		return
	}
	level.Info(logger).Log("msg", "config", "value", conf.Sources.UrlWeatherBit)


	db, err := database.NewPostgres(conf.Database)
	if err != nil {
		level.Error(logger).Log("msg", "database error", "exit", err.Error())
		return
	}

	stationList, err := csv.CSVToMap(*csvFile)
	if err != nil {
		level.Error(logger).Log("msg", "CSV loading error", "err", err.Error())
		return
	}

	 wbh := WeatherHistorical{
	 	 db:                    db,
		 logger:                logger,
		 conf:                  *conf,
		 stationList:           *stationList,
		 dailyRequestCounter:   0,
		 secondsRequestCounter: 0,
	 }

	 wbh.precessStations(time.Now())
}

func (wbh WeatherHistorical)precessStations(date time.Time) {
	wbh.dailyStartTime = date
	wbh.secondsStartTime = date
	for k,v := range wbh.stationList {
		level.Info(wbh.logger).Log("msg", "Process station", "innerId", k, "station", v)
		wbh.processRequest(k, v, date)
	}
}

func (wbh WeatherHistorical)processRequest(stID string, st string, end time.Time) error {
	startDate := end
	requestsPerSecond := wbh.conf.WeatherBit.NumberOfRequestPerSecond
	requestsPerDay := wbh.conf.WeatherBit.NumberOfRequestPerDay
	for {
		start := end.AddDate(0, 0, -14)
		sDate := start.Format(common.TimeLayoutWBH)
		eDate := end.Format(common.TimeLayoutWBH)
		if yearCheck(startDate, end, wbh.conf.WeatherBit.NumberOfYears) {
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

func (wbh WeatherHistorical)processUpdate(stID string, st string, start string, end string) error {
	err := wbh.db.CreateTable(stID)
	if err != nil {
		level.Error(wbh.logger).Log("msg", "table create error", "err", err)
		return err
	}

	url := wbh.conf.Sources.UrlWeatherBit + "/history/hourly?station=" + st + "&key=" + wbh.conf.Sources.KeyWeatherBit + "&start_date=" + start + "&end_date=" + end
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

func (wbh WeatherHistorical)checkPeriod(stID string, start string, end string) bool {
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

func yearCheck(start, end time.Time, yearsCount int) bool {
	return 	start.Year() - end.Year() >= yearsCount &&
		start.Month() > end.Month()
}

