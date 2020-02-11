package main

import (
	"flag"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"


	stations "github.com/flasherup/gradtage.de/stationssvc/impl"
)

func main() {
	aggregatorURL := flag.String("aggregator.url", "localhost:8106", "Url of daily aggregator service")
	stationsURL := flag.String("stations.url", "localhost:8102", "Url of stations service")
	start := flag.String("start", "2018-01-29T01:00:00Z", "Start date")
	end := flag.String("end", "2020-03-01T15:00:00Z", "End date")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "force update client",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	client := impl.NewDailyAggregatorSCVClient(*aggregatorURL,logger)

	stationsService := stations.NewStationsSCVClient(*stationsURL, logger)

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
	err = client.ForceUpdate(ids, *start, *end)
	if err != nil {
		level.Error(logger).Log("msg", "Force Update Error", "err", err)

	} else {
		level.Info(logger).Log("msg", "Force Update  Complete", "err", err)
	}
}
