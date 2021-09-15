package main

import (
	"fmt"
	"github.com/flasherup/gradtage.de/autocompletesvc/cmd/client/data"
	"github.com/flasherup/gradtage.de/autocompletesvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
	"strings"
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
	client := impl.NewAutocompleteSCVClient("212.227.215.17:8109",logger)
	//client := impl.NewAutocompleteSCVClient("localhost:8109",logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	/*sources := data.AutocompleteStations

	err := client.ResetSources(sources)
	if err != nil {
		level.Error(logger).Log("msg", "Reset Sources error", "err", err)

	}*/

	/*res, err := client.GetAutocomplete("f3836")
	if err != nil {
		level.Error(logger).Log("msg", "Get Autocomplete error", "err", err)

	}
	fmt.Println("Get Autocomplete response:", res)*/

	checkStationsName(client)
}

func checkStationsName(client *impl.AutocompleteSVCClient ) {
	fmt.Printf("%s;%s;%s\n","id", "source", "name")
	for _,v := range data.StationsList {
		res, err := client.GetAutocomplete(v)
		if err != nil {
			fmt.Println("Error", err)
		}
		for _,auto := range res {
			for _,a := range auto {
				l := strings.ToLower(a.ID)
				if v == l {
					i := strings.Index(a.CityNameEnglish, "'")
					if i >= 0 {
						fmt.Printf("%s;%s;%s\n",a.ID, a.SourceID, a.CityNameEnglish)
					}
					break
				}
			}
		}
	}
}