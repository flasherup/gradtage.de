package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/autocompletesvc/acrpc"
	"os"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/flasherup/gradtage.de/autocompletesvc/impl"
)

const (
	operationFromRemoteToLocal = "frtl"         // from remote to local
	operationFromLocalToRemote = "fltr"         // from local to remote
	operationAutocomplete      = "autocomplete" // autocomplete operation
	operationDefault           = "default"      // default operation
)

const (
	remoteAddr = "212.227.215.17:8109"
	localAddr  = "localhost:8109"
)

func main() {
	operationPrm := flag.String("operation", operationDefault, "Name of the operation")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "autocompleteUtils",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	if *operationPrm == operationDefault {
		logger.Log("msg", "No operation selected")
	}

	switch *operationPrm {
	case operationFromRemoteToLocal:
		fromRemoteToLocal(logger)
	case operationFromLocalToRemote:
		fromLocalToRemote(logger)
	case operationAutocomplete:
		autocomplete(logger)
	}
}

func fromRemoteToLocal(logger log.Logger) {
	remote := impl.NewAutocompleteSCVClient(remoteAddr, logger)
	local := impl.NewAutocompleteSCVClient(localAddr, logger)

	stations, err := remote.GetAllStations()
	if err != nil {
		level.Error(logger).Log("msg", "Failed to get stations", "err", err)
		return
	}

	autocompletes := stationsToAutocomplete(stations)

	err = local.ResetSources(autocompletes)
	if err != nil {
		level.Error(logger).Log("msg", "Failed to reset stations", "err", err)
		return
	}

	level.Info(logger).Log("msg", "Stations transferred")
}

func fromLocalToRemote(logger log.Logger) {
	//remote := impl.NewAutocompleteSCVClient("212.227.215.17:8109", logger)
	//local := impl.NewAutocompleteSCVClient("localhost:8109", logger)
}

func autocomplete(logger log.Logger) {
	local := impl.NewAutocompleteSCVClient(localAddr, logger)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("exit", text) == 0 {
			return
		}

		autocomplete, err := local.GetAutocomplete(text)
		if err != nil {
			level.Error(logger).Log("msg", "Check Autocomplete Error", "error", err.Error())
		}

		for k, v := range autocomplete {
			level.Info(logger).Log("msg", "station", "id", k)
			for _, a := range v {
				level.Info(logger).Log("msg", "autocomplete", "City", a.CityNameEnglish, "Country", a.CountryNameEnglish, "Lat", a.Latitude, "Lon", a.Longitude)
			}
		}
	}
}

func stationsToAutocomplete(stations map[string]*acrpc.Source) []autocompletesvc.Autocomplete {
	res := make([]autocompletesvc.Autocomplete, 0, len(stations))
	for _, v := range stations {
		res = append(res, autocompletesvc.Autocomplete{
			ID:                 v.ID,
			SourceID:           v.SourceID,
			Latitude:           v.Latitude,
			Longitude:          v.Longitude,
			Source:             v.Source,
			Reports:            v.Reports,
			ISO2Country:        v.ISO2Country,
			ISO3Country:        v.ISO3Country,
			Prio:               v.Prio,
			CityNameEnglish:    v.CityNameEnglish,
			CityNameNative:     v.CityNameNative,
			CountryNameEnglish: v.CountryNameEnglish,
			CountryNameNative:  v.CountryNameNative,
			ICAO:               v.ICAO,
			WMO:                v.WMO,
			CWOP:               v.CWOP,
			Maslib:             v.Maslib,
			National_ID:        v.National_ID,
			IATA:               v.IATA,
			USAF_WBAN:          v.USAF_WBAN,
			GHCN:               v.GHCN,
			NWSLI:              v.NWSLI,
			Elevation:          v.Elevation,
		})
	}
	return res
}
