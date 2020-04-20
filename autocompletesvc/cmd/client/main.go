package main

import (
	"github.com/flasherup/gradtage.de/autocompletesvc/impl"
	"github.com/flasherup/gradtage.de/localutils/data"
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
			"svc", "autocompleteclient",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	//client := impl.NewAutocompleteSCVClient("localhost:8109",logger)
	client := impl.NewAutocompleteSCVClient("82.165.18.228:8109",logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	sources := data.AutocompleteStations

	err := client.ResetSources(sources)
	if err != nil {
		level.Error(logger).Log("msg", "Reset Sources error", "err", err)

	}

	/*res, err := client.GetAutocomplete("EDB")
	if err != nil {
		level.Error(logger).Log("msg", "Get Autocomplete error", "err", err)

	}
	fmt.Println("Get Autocomplete response:", res);*/
}