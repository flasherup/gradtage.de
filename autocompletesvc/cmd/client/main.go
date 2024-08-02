package main

import (
	"github.com/flasherup/gradtage.de/autocompletesvc/impl"
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
	//client := impl.NewAutocompleteSCVClient("212.227.215.17:8109",logger)
	client := impl.NewAutocompleteSCVClient("localhost:8109", logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	//checkStationsName(client)
	//getAllStations(client, logger)
	checkAutocomplete(client, logger)
}

func checkAutocomplete(client *impl.AutocompleteSVCClient, logger log.Logger) {
	level.Info(logger).Log("msg", "Check Autocomplete")
	autocomplete, err := client.GetAutocomplete("berlin")
	if err != nil {
		level.Error(logger).Log("msg", "Check Autocomplete Error", "error", err.Error())
	}

	level.Info(logger).Log("msg", "Check Autocomplete Success", "length", len(autocomplete))

	for k, v := range autocomplete {
		level.Info(logger).Log("msg", "station", "id", k)
		for _, a := range v {
			level.Info(logger).Log("msg", "autocomplete", "name", a.CityNameEnglish)
		}
	}
}

func getAllStations(client *impl.AutocompleteSVCClient, logger log.Logger) {
	level.Info(logger).Log("msg", "Get All Stations")
	stations, err := client.GetAllStations()
	if err != nil {
		level.Error(logger).Log("msg", "Get All Stations Error", "error", err.Error())
	}

	level.Info(logger).Log("msg", "Get All Stations Success", "length", len(stations))

	for k, v := range stations {
		level.Info(logger).Log("msg", "station", "id", k, "lat", v.Longitude)
	}
}
