package main

import (
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/stationssvc/impl"
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
			"svc", "stationssvcc",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	client := impl.NewStationsSCVClient("localhost:9090",logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	//Just for test
	client.GetStations([]string{"TEST"})
	client.GetAllStations()
	stations := []stationssvc.Station{
		stationssvc.Station{ID:"TEST", Name:"TestName", Timezone:"CEE" },
		}
	client.AddStations(stations)
}
