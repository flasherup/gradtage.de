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
	//"github.com/flasherup/gradtage.de/stationssvc"
	//stations "github.com/flasherup/gradtage.de/stationssvc/impl"
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

	/*filesList := map[string]string {
		"addon_icao_prioa.csv": "CET",
		"addon_icao_priob.csv": "CET",
		"addon_icao_prioc.csv": "CET",
		"prioa_ch.csv": "CET",
		"prioa_es.csv": "CET",
		"prioa_fr.csv": "CET",
		"prioa_gb.csv": "CET",
		"prioa_it.csv": "CET",
		"prioa_li.csv": "CET",
		"prioa_lu.csv": "CET",
		"prioa_nl.csv": "GMT",
		"Prioa_no.csv": "WET",
		"prioa_pl.csv": "WET",
		"prioa_pt.csv": "WET",
		"prioa_se.csv": "WET",
		"priob_at.csv": "WET",
		"priob_au.csv": "WET",
		"priob_be.csv": "WET",
		"priob_by.csv": "WET",
		"priob_ca.csv": "WET",
		"priob_de.csv": "WET",
		"priob_dk.csv": "WET",
		"priob_fi.csv": "WET",
		"priob_gr.csv": "WET",
		"priob_ie.csv": "WET",
		"prioc_us.csv": "WET",
	}*/

	filesList := map[string]string {
		"PrioD.csv": "CET",
	}

 	//fromCSVListToStations("./data", filesList, logger)
	//fromCSVListToAutocomplete("data", filesList, logger)
	fromCSVToList("./data", filesList, logger)
}


func fromCSVListToStations(path string, filesList map[string]string, logger log.Logger) {
	stationsLocal := stations.NewStationsSCVClient("localhost:8102", logger)

	allStation := make([]stationssvc.Station, 0)
	for fileName,timeZone := range filesList {
		stsl, error := parsers.CSVToStationsList(path + "/" + fileName)
		if error != nil {
			println("Error", error.Error())
			continue
		}

		for i,_ := range stsl {
			stsl[i].Timezone = timeZone;
			stsl[i].SourceType = common.SrcTypeWeatherBit
			//fmt.Println(i,v)
		}

		for i,v := range stsl {
			fmt.Println(i,v)
		}

		allStation = append(allStation, stsl...)
	}

	_,err := stationsLocal.ResetStations(allStation)
	if err != nil {
		level.Error(logger).Log("msg", "AddStations error", "err", err)
	}
}

func fromCSVListToAutocomplete(path string, filesList map[string]string, logger log.Logger) {
	autocompleteLocal := autocomplete.NewAutocompleteSCVClient("localhost:8109", logger)

	allStation := make([]autocompletesvc.Source, 0)
	for fileName,_ := range filesList {
		stsl, error := parsers.CSVToAutocompleteList(path + "/" + fileName)
		if error != nil {
			println("Error", error.Error())
			continue
		}

		for i,v := range stsl {
			fmt.Println(i,v)
		}

		allStation = append(allStation, stsl...)
	}

	err := autocompleteLocal.ResetSources(allStation)
	if err != nil {
		level.Error(logger).Log("msg", "AddAutocomplete error", "err", err)
	}
}

func fromCSVToList(path string, filesList map[string]string, logger log.Logger) {
	//File save logic
	csvFile, err := os.Create("stations.csv")

	if err != nil {
		fmt.Printf("failed creating file: %s", err.Error())
	}

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Comma = ';'

	for fileName,_ := range filesList {
		stsl, error := parsers.CSVToStationsList(path + "/" + fileName)
		if error != nil {
			println("Error", error.Error())
			continue
		}

		for i,_ := range stsl {
			err = csvwriter.Write([]string{stsl[i].ID, stsl[i].SourceID })
			if err != nil {
				fmt.Println(i, stsl[i].ID, "Error:", err)
			}
		}
	}

	csvwriter.Flush()
	csvFile.Close()
}