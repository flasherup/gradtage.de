package main

import (
	"github.com/flasherup/gradtage.de/noaascrapersvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "noaascrapersvc",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	client := impl.NewNoaaScraperSVCClient("localhost:8108",logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	//Just for test
	period, err := client.GetPeriod("KBOS", "2020-02-03T01:00:00", "2020-02-04T23:00:00")
	if err != nil {
		level.Error(logger).Log("msg", "GetPeriod Error", "err", err)
	} else {
		for _,v := range period.Temps {
			level.Info(logger).Log("msg", "sts", "id", "KBOS", "date:", v.Date, "temp:", v.Temperature)
		}
	}

	dates, err := client.GetUpdateDate([]string{"KBOS"})
	if err != nil {
		level.Error(logger).Log("msg", "Get Update Dates Error", "err", dates.Err)
	} else {
		for k,v := range dates.Dates {
			level.Info(logger).Log("msg", "Update Date ", "id", k, "date:", v)
		}
	}
}
