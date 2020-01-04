package main

import (
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dailysvc/impl"
	stations "github.com/flasherup/gradtage.de/stationssvc/impl"
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
			"svc", "dailysvcc",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	//client := impl.NewDailySCVClient("localhost:8104",logger)
	client := impl.NewDailySCVClient("82.165.18.228:8104",logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	_, err := client.PushPeriod("KBOS", period())
	if err != nil {
		level.Error(logger).Log("msg", "PushPeriod Error", "err", err)

	}

	//Just for test
	/*period := client.GetPeriod("KBOS", "2019-12-29T01:00:00", "2019-12-29T15:00:00")
	if period.Err != "nil" {
		level.Error(logger).Log("msg", "GetPeriod Error", "err", period.Err)

	} else {
		for _,v := range period.Temps {
			level.Info(logger).Log("msg", "sts", "id", "KBOS", "date:", v.Date, "temp:", v.Temperature)
		}
	}



	dates := client.GetUpdateDate([]string{"KBOS"})
	if dates.Err != "nil" {
		level.Error(logger).Log("msg", "Get Update Dates Error", "err", dates.Err)

	} else {
		for k,v := range dates.Dates {
			level.Info(logger).Log("msg", "Update Date ", "id", k, "date:", v)
		}
	}*/

	stations := stations.NewStationsSCVClient("82.165.18.228:8102",logger)

	sts, err := stations.GetAllStations()
	if err != nil {
		level.Error(logger).Log("msg", "Get stations err Error", "err", err)
	} else {
		for k := range sts.Sts {
			_, err := client.UpdateAvgForYear(k)
			if err != nil {
				level.Error(logger).Log("msg", "UpdateAvgForYear Error", "err", err)
			}
		}
	}

	_, err = client.UpdateAvgForDOY("KBOS", 1)
	if err != nil {
		level.Error(logger).Log("msg", "UpdateAvgForDOY Error", "err", err)
	}



	respAvg, err := client.GetAvg("KBOS")
	if err != nil {
		level.Error(logger).Log("msg", "GetAvg Error", "err", err)
	} else {
		for _,v := range respAvg.Temps {
			level.Info(logger).Log("msg", "sts", "id", "KBOS", "date:", v.Date, "temp:", v.Temperature)
		}
	}



}


func period() []dailysvc.Temperature {
	return []dailysvc.Temperature {
		dailysvc.Temperature{ Date:"2019-12-29T01:00:00", Temperature:9.0 },
		dailysvc.Temperature{ Date:"2019-12-29T02:00:00", Temperature:10.0 },
		dailysvc.Temperature{ Date:"2019-12-29T03:00:00", Temperature:11.0 },
		dailysvc.Temperature{ Date:"2019-12-29T04:00:00", Temperature:12.0 },
		dailysvc.Temperature{ Date:"2019-12-29T05:00:00", Temperature:13.0 },
		dailysvc.Temperature{ Date:"2019-12-29T06:00:00", Temperature:14.0 },
		dailysvc.Temperature{ Date:"2019-12-29T07:00:00", Temperature:15.0 },
		dailysvc.Temperature{ Date:"2019-12-29T08:00:00", Temperature:16.0 },
		dailysvc.Temperature{ Date:"2019-12-29T09:00:00", Temperature:17.0 },
		dailysvc.Temperature{ Date:"2019-12-29T10:00:00", Temperature:18.0 },
		dailysvc.Temperature{ Date:"2019-12-29T11:00:00", Temperature:19.0 },
		dailysvc.Temperature{ Date:"2019-12-29T12:00:00", Temperature:20.0 },
		dailysvc.Temperature{ Date:"2019-12-29T13:00:00", Temperature:21.0 },
		dailysvc.Temperature{ Date:"2019-12-29T14:00:00", Temperature:22.0 },
		dailysvc.Temperature{ Date:"2019-12-29T15:00:00", Temperature:23.0 },
	}
}
