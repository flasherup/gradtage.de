package main

import (
	"fmt"
	"github.com/flasherup/gradtage.de/localutils/meteostat/parser"
	"github.com/flasherup/gradtage.de/localutils/meteostat/source"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

func main() {
	updateMeteostat()
}

func updateMeteostat() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "metiostat",
			"ts", log.DefaultTimestampUTC,
			"caller", log.Caller(3),
		)
	}

	mst := source.NewMeteostat("BkgNZdCV", logger)

	ids := []string{
		"01001",
	}

	ch := make(chan *parser.ParsedData)
	go mst.FetchTemperature(ch, ids)

	count := 0.0
	for range ids {
		st := <-ch
		if st != nil && st.Success {
			temps := fmt.Sprintf("tmps: %v", st.Temps)

			level.Info(logger).Log("msg", "Temperature updated form Meteostat", "stations", count, "temp", temps)
			count++
		} else {
			level.Warn(logger).Log("msg", "Station is not updated")
		}
	}
	level.Info(logger).Log("msg", "Temperature updated form CheckWX", "stations", count)
}
