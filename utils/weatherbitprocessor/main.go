package main

import (
	"flag"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/utils/weatherbitprocessor/config"
	"github.com/flasherup/gradtage.de/utils/weatherbitprocessor/csv"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl"
	"github.com/flasherup/gradtage.de/weatherbitsvc/weatherbitgrpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
	"time"
)

func main() {
	configFile := flag.String("config.file", "config.yml", "Config file name.")
	startDate := flag.String("start", "2021-05-21", "Start Date.")
	endDate := flag.String("end", "2021-05-27", "End Date.")
	csvFile := flag.String("csv.file", "Greece.csv", "CSV file name. ")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger)
	}

	//Config
	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		level.Error(logger).Log("msg", "config loading error", "exit", err.Error())
		return
	}

	client := impl.NewWeatherBitSVCClient(conf.Clients.WeatherBitAddr,logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	calculateEntries(client, logger, *csvFile, *startDate, *endDate)


}

func calculateEntries(client *impl.WeatherBitSVCClient, logger log.Logger, stations, startDate, endDate string) {
	//Just for test
	stationList, err := csv.CSVToMap(stations)
	if err != nil {
		level.Error(logger).Log("msg", "CSV loading error", "err", err.Error())
		return
	}

	for innerId, stationId := range *stationList {
		level.Info(logger).Log("Station:", stationId, "innerId", innerId)
		data, err := client.GetPeriod([]string{innerId}, startDate, endDate)
		if err != nil {
			level.Error(logger).Log("Station error:", stationId, "innerId", innerId, "error", err)
			continue
		}

		for _,temperatures := range data.Temps {
			if len(temperatures.Temps) == 0 {
				level.Error(logger).Log("Error", "no entries", "Station:", stationId, "innerId", innerId)
				continue
			}

			counted := countEntriesPerDay(&temperatures.Temps)
			for k,v := range counted {
				//level.Info(logger).Log("date", k, "hours", v)
				if v < 24 {
					level.Error(logger).Log("Error", "less the 24 entries", "date", k, "hours", v, "Station:", stationId, "innerId", innerId)
				}
			}
		}
	}
}

func countEntriesPerDay(temps *[]*weatherbitgrpc.Temperature) map[string]int {
	res := make(map[string]int)
	for _,v := range *temps {
		date,err := time.Parse(common.TimeLayout, v.Date)
		if  err !=  nil {
			continue
		}
		cattedDate := date.Format("2006-01-02")
		if val, ok := res[cattedDate]; ok {
			res[cattedDate] = val+1
		} else {
			res[cattedDate] = 1
		}
	}
	return res
}

