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
	db 			database.WeatherBitDB
	logger  	log.Logger
	conf		config.WeatherBitConfig
	stationlist	map[string]string
}

func main() {
	configFile := flag.String("config.file", "config.yml", "Config file name. ")
	csvFile := flag.String("csv.file", "Greece.csv", "CSV file name. ")
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
	 	 db: db,
		 logger: logger,
		 conf: *conf,
		 stationlist: *stationList,
	 }

	 date := time.Now()
	 wbh.precessStations(date)
}

func (wbh WeatherHistorical)processRequest(stID string, st string, end time.Time) error {

	dailyRequestCounter := 0
	secondsRequestCounter := 0
	startDate := end
	for {
		start := end.AddDate(0, 0, -14)
		sDate := start.Format(common.TimeLayoutWBH)
		eDate := end.Format(common.TimeLayoutWBH)
		wbh.processUpdate(stID, st, sDate, eDate)
		dailyRequestCounter++
		secondsRequestCounter++
		secondsRequestCounter = sleepCheck(wbh.conf.WeatherBit.NumberOfRequestPerSecond, secondsRequestCounter, time.Second)
		dailyRequestCounter = sleepCheck(wbh.conf.WeatherBit.NumberOfRequestPerDay, dailyRequestCounter, time.Second * 5)
		end = start
		if !yearCheck(startDate, end, wbh.conf.WeatherBit.NumberOfYears) {
			break
		}
	}
	return nil
}

func sleepCheck(numberOfRequests, counter int, duration time.Duration) int {
	if counter >= numberOfRequests{
		time.Sleep(duration)
		return 0
	}
	return counter
}

func (wbh WeatherHistorical)precessStations(date time.Time) {
	time.Now()
	for k,v := range wbh.stationlist {
		level.Info(wbh.logger).Log("msg", "Process station", "innerId", k, "station", v)
		wbh.processRequest(k, v, date)
	}
}

func yearCheck(start, end time.Time, yearsCount int) bool {
	if start.Year() - end.Year() >= yearsCount && start.Month() >= end.Month() && start.Day() >= end.Day() && start.Hour() >= end.Hour(){
		return false
	}
	return true
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
