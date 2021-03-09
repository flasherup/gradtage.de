package main

import (
	"flag"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/utils/weatherbitHistorical/config"
	"github.com/flasherup/gradtage.de/utils/weatherbitHistorical/database"
	"github.com/flasherup/gradtage.de/utils/weatherbitHistorical/parser"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type weatherHistorical struct {
	db 			database.WeatherBitDB
	logger  	log.Logger
	conf		config.WeatherBitConfig
}

func main() {
	configFile := flag.String("config.file", "config.yml", "Config file name.")
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
	 wbh := weatherHistorical{
	 	 db: db,
		 logger: logger,
		 conf: *conf,
	 }

	 date := time.Now()
	 endDate := date.Format(common.TimeLayoutWBH)
	 sDate := date.AddDate(0, 0, -7)
	 startDate := sDate.Format(common.TimeLayoutWBH)

	 precessStations(wbh, startDate, endDate)
}

func precessStations(wbh weatherHistorical, start string, end string) {
	stations := map[string]string{
		"KNYC": "KNYC",
		"WMO7650": "LFML",
		"KATT": "KATT",
		"EDDH": "EDDH",
		"CYYC": "CYYC",
		"WMO10224": "10224",
		"LEBL": "LEBL",
		"WMO8181": "081810",
		"ESMS":"ESMS",
		"LFBN":"LFBN",
		"D4932":"D4932",
		"W07301399999":"073013-99999",
	}
	for k,v := range stations {
		processUpdate(k, v, start, end, wbh)
	}
}

func processUpdate(stID string, st string, start string, end string, wbh weatherHistorical, ) {
		err := wbh.db.CreateTable(stID)
		if err != nil {
			level.Error(wbh.logger).Log("msg", "table create error", "err", err)
			return
		}

		url := wbh.conf.Sources.UrlWeatherBit + "/history/hourly?station=" + st + "&key=" + wbh.conf.Sources.KeyWeatherBit + "&start_date=" + start + "&end_date=" + end
		level.Info(wbh.logger).Log("msg", "weather bit request", "url", url)

		client := &http.Client{
			Timeout: time.Second * 10,
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			level.Error(wbh.logger).Log("msg", "request error", "err", err)
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			level.Error(wbh.logger).Log("msg", "request error", "err", err)
			return
		}
		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			level.Error(wbh.logger).Log("msg", "response read error", "err", err)
			return
		}

		result, err := parser.ParseWeatherBit(&contents)
		if (err != nil) {
			level.Error(wbh.logger).Log("msg", "weather bit data parse error", "err", err)
			return
		}

		err = wbh.db.PushData(stID, result)
		if err != nil {
			level.Error(wbh.logger).Log("msg", "data push error", "err", err)
			return
		}

}
