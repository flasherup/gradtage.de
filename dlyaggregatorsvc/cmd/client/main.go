package main

import (
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"


	stations "github.com/flasherup/gradtage.de/stationssvc/impl"
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
	client := impl.NewDailyAggregatorSCVClient("localhost:8106",logger)

	stationsService := stations.NewStationsSCVClient("localhost:8102", logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")


	sts, err := stationsService.GetAllStations()
	if err != nil {
		level.Error(logger).Log("msg", "GetStations error", "err", err)
		return
	}

	ids := make([]string, len(sts.Sts))
	i := 0
	for k,_ := range sts.Sts {
		ids[i] = k
		i++
	}


	//Just for test
	err = client.ForceUpdate(ids, "2018-01-29T01:00:00Z", "2020-06-23T15:00:00Z")
	if err != nil {
		level.Error(logger).Log("msg", "GetPeriod Error", "err", err)

	}
}
