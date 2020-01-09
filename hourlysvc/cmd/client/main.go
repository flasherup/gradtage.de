package main

import (
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/hourlysvc/impl"
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
			"svc", "hourlysvcc",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	client := impl.NewHourlySCVClient("82.165.18.228:8103",logger)
	//client := impl.NewHourlySCVClient("localhost:8103",logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	/*_, err := client.PushPeriod("KBOG", period())
	if err != nil {
		level.Error(logger).Log("msg", "PushPeriod Error", "err", err)

	}

	//Just for test
	period, err := client.GetPeriod("KBOS", "2019-12-29T01:00:00", "2019-12-29T15:00:00")
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
	}*/

	temps, err := client.GetLatest([]string{"KBOS"})
	if err != nil {
		level.Error(logger).Log("msg", "Get Latest Error", "err", err)
	} else {
		for k,v := range temps.Temps {

			level.Info(logger).Log("msg", "Latest ", "id", k, "date:", v.Date, "temp", v.Temperature)
		}
	}
}


func period() []hourlysvc.Temperature {
	return []hourlysvc.Temperature {
		hourlysvc.Temperature{ Date:"2019-12-29T01:00:00", Temperature:9.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T02:00:00", Temperature:10.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T03:00:00", Temperature:11.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T04:00:00", Temperature:12.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T05:00:00", Temperature:13.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T06:00:00", Temperature:14.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T07:00:00", Temperature:15.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T08:00:00", Temperature:16.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T09:00:00", Temperature:17.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T10:00:00", Temperature:18.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T11:00:00", Temperature:19.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T12:00:00", Temperature:20.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T13:00:00", Temperature:21.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T14:00:00", Temperature:22.0 },
		hourlysvc.Temperature{ Date:"2019-12-29T15:00:00", Temperature:23.0 },
	}
}
