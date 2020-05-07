package main

import (
	"flag"
	daily "github.com/flasherup/gradtage.de/dailysvc/impl"
	stations "github.com/flasherup/gradtage.de/stationssvc/impl"
	"github.com/flasherup/gradtage.de/utils/dailyupdateavg/config"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
	"time"
)

func main() {
	configFile := flag.String("config.file", "dwdhistory.yml", "Config file name.")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "apisvc",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	//Config
	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		level.Error(logger).Log("msg", "config loading error", "exit", err.Error())
		return
	}


	stationsService := stations.NewStationsSCVClient(conf.StationsAddr, logger)
	dailyService := daily.NewDailySCVClient(conf.DailyAddr, logger)

	level.Info(logger).Log("msg", "Start Average Update")

	sts, err := stationsService.GetAllStations()
	if err != nil {
		level.Error(logger).Log("msg", "GetStations error", "err", err)
		return
	}

	count := 0
	for k,_ := range sts.Sts {
		level.Info(logger).Log("msg", "Update Average", "sts", k)
		resp, err := dailyService.UpdateAvgForYear(k)
		if err != nil {
			level.Error(logger).Log("msg", "Average Update Error", "id", k, "err", err)
		} else if resp.Err != "nil" {
			level.Error(logger).Log("msg", "Average Update Error", "id", k, "err", resp.Err)
		} else {
			count++
		}
		time.Sleep(time.Millisecond * 10)
	}
	level.Info(logger).Log("msg", "Start Average Update Complete", "count", count)
}