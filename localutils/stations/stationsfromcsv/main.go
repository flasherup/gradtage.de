package main

import (
	"encoding/csv"
	"fmt"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	autocomplete "github.com/flasherup/gradtage.de/autocompletesvc/impl"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/localutils/stations/stationsfromcsv/parsers"
	"github.com/flasherup/gradtage.de/stationssvc"
	stations "github.com/flasherup/gradtage.de/stationssvc/impl"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "stations_upgrade",
			"ts", log.DefaultTimestampUTC,
			"caller", log.Caller(3),
		)
	}

	filesList := []string{
		"EDG_Stationlist_Masterfile.csv",
	}

	s := stations.NewStationsSCVClient("212.227.215.17:8102", logger)
	//s := stations.NewStationsSCVClient("localhost:8102", logger)
	fromCSVListToStations("data", filesList, s, logger, nil)

	//a := autocomplete.NewAutocompleteSCVClient("localhost:8109", logger)
	a := autocomplete.NewAutocompleteSCVClient("212.227.215.17:8109", logger)
	fromCSVListToAutocomplete("data", filesList, a, logger, nil)
	//fromCSVToList("./data", filesList, logger)

}

func fromCSVListToStations(path string, filesList []string, stationsLocal *stations.StationsSVCClient, logger log.Logger, filter map[string]bool) {

	//allStation := make([]stationssvc.Station, 0)
	for _, fileName := range filesList {
		fmt.Println("Process", fileName)
		stsl, error := parsers.ParseStationsCSV(path + "/" + fileName)
		if error != nil {
			println("Error", error.Error())
			continue
		}

		sts := make([]stationssvc.Station, 0)

		_, err := stationsLocal.ResetStations([]stationssvc.Station{})
		if err != nil {
			level.Error(logger).Log("msg", "AddStations error", "err", err)
		}

		for _, v := range stsl {
			tz, err := common.GetTimezoneFormLatLon(v.Latitude, v.Longitude)
			if err != nil {
				fmt.Println("Get timezone error", err)
			}

			st := stationssvc.Station{
				ID:         v.ID,
				Name:       v.CityNameEnglish,
				Timezone:   tz,
				SourceType: common.SrcTypeWeatherBit,
				SourceID:   v.SourceID,
			}

			sts = append(sts, st)
		}
		fmt.Println(len(stsl), len(sts))
		_, err = stationsLocal.AddStations(sts)
		if err != nil {
			level.Error(logger).Log("msg", "AddStations error", "err", err)
		}
	}
}

func fromCSVListToAutocomplete(path string, filesList []string, autocompleteLocal *autocomplete.AutocompleteSVCClient, logger log.Logger, filter map[string]bool) {
	for _, fileName := range filesList {
		stsl, err := parsers.ParseStationsCSV(path + "/" + fileName)
		if err != nil {
			println("Error", err.Error())
			continue
		}

		sts := make([]autocompletesvc.Autocomplete, 0, len(stsl))

		for _, v := range stsl {
			sts = append(sts, autocompletesvc.Autocomplete{
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
		err = autocompleteLocal.ResetSources(sts)
		if err != nil {
			level.Error(logger).Log("msg", "ResetSources error", "err", err)
		}
	}
}

func fromCSVToList(path string, filesList []string, logger log.Logger) {
	//File save logic
	csvFile, err := os.Create("stations.csv")

	if err != nil {
		fmt.Printf("failed creating file: %s", err.Error())
	}

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Comma = ';'

	for _, fileName := range filesList {
		stsl, error := parsers.CSVToStationsList(path + "/" + fileName)
		if error != nil {
			println("Error", error.Error())
			continue
		}

		for i, _ := range stsl {
			err = csvwriter.Write([]string{stsl[i].ID, stsl[i].SourceID})
			if err != nil {
				fmt.Println(i, stsl[i].ID, "Error:", err)
			}
		}
	}

	csvwriter.Flush()
	csvFile.Close()
}
