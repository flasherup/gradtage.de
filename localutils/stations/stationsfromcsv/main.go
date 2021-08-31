package main

import (
	"encoding/csv"
	"fmt"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	autocomplete "github.com/flasherup/gradtage.de/autocompletesvc/impl"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/localutils/stations/stationsfromcsv/parsers"
	"github.com/flasherup/gradtage.de/stationssvc"

	//"github.com/flasherup/gradtage.de/stationssvc"
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

	filesList := []string {
		"stations_V2_20210831.csv",
		//"EDG_Stationlist_Masterfile.csv",
		/*"PrioD.csv",
		"addon_icao_prioa.csv",
		"addon_icao_priob.csv",
		"addon_icao_prioc.csv",
		"prioa_ch.csv",
		"prioa_es.csv",
		"prioa_fr.csv",
		"prioa_gb.csv",
		"prioa_it.csv",
		"prioa_li.csv",
		"prioa_lu.csv",
		"prioa_nl.csv",
		"Prioa_no.csv",
		"prioa_pl.csv",
		"prioa_pt.csv",
		"prioa_se.csv",
		"priob_at.csv",
		"priob_au.csv",
		"priob_be.csv",
		"priob_by.csv",
		"priob_ca.csv",
		"priob_de.csv",
		"priob_dk.csv",
		"priob_fi.csv",
		"priob_gr.csv",
		"priob_ie.csv",
		"prioc_us.csv",*/
	}

 	//fromCSVListToStations("./data", filesList, logger)
	//fromCSVListToAutocomplete("./data", filesList, logger)
	fromCSVToList("./data", filesList, logger)
}


func fromCSVListToStations(path string, filesList []string, logger log.Logger) {
	stationsLocal := stations.NewStationsSCVClient("212.227.214.163:8102", logger)

	//allStation := make([]stationssvc.Station, 0)
	for _,fileName := range filesList {
		fmt.Println("Process", fileName)
		stsl, error := parsers.ParseStationsCSV(path + "/" + fileName)
		if error != nil {
			println("Error", error.Error())
			continue
		}

		sts := make([]stationssvc.Station, len(stsl))

		_,err := stationsLocal.ResetStations([]stationssvc.Station{})
		if err != nil {
			level.Error(logger).Log("msg", "AddStations error", "err", err)
		}

		for i,v := range stsl {

			tz,err := common.GetTimezoneFormLatLon(v.Latitude, v.Longitude)
			if err != nil {
				fmt.Println("Get timezone error", err)
			}

			st := stationssvc.Station{
				ID: v.ID,
				Name: v.CityNameEnglish,
				Timezone: tz,
				SourceType: common.SrcTypeWeatherBit,
				SourceID: v.SourceID,
			}

			sts[i] = st
		}
		_,err = stationsLocal.AddStations(sts)
		if err != nil {
			level.Error(logger).Log("msg", "AddStations error", "err", err)
		}

		//allStation = append(allStation, sts...)
	}
}

func fromCSVListToAutocomplete(path string, filesList []string, logger log.Logger) {
	//autocompleteLocal := autocomplete.NewAutocompleteSCVClient("212.227.214.163:8109", logger)
	autocompleteLocal := autocomplete.NewAutocompleteSCVClient("localhost:8109", logger)

	for _,fileName := range filesList {
		fmt.Println("Process", fileName)
		stsl, error := parsers.ParseStationsCSV(path + "/" + fileName)
		if error != nil {
			println("Error", error.Error())
			continue
		}

		sts := make([]autocompletesvc.Autocomplete, len(stsl))

		err := autocompleteLocal.ResetSources([]autocompletesvc.Autocomplete{})
		if err != nil {
			level.Error(logger).Log("msg", "Seset Autocomplete error", "err", err)
		}

		for i,v := range stsl {
			sts[i] = autocompletesvc.Autocomplete{
				ID: v.ID,
				SourceID: v.SourceID,
				Latitude: v.Latitude,
				Longitude: v.Longitude,
				Source: v.Source,
				Reports: v.Reports,
				ISO2Country: v.ISO2Country,
				ISO3Country: v.ISO3Country,
				Prio: v.Prio,
				CityNameEnglish: v.CityNameEnglish,
				CityNameNative: v.CityNameNative,
				CountryNameEnglish: v.CountryNameEnglish,
				CountryNameNative: v.CountryNameNative,
				ICAO: v.ICAO,
				WMO: v.WMO,
				CWOP: v.CWOP,
				Maslib: v.Maslib,
				National_ID: v.National_ID,
				IATA: v.IATA,
				USAF_WBAN: v.USAF_WBAN,
				GHCN: v.GHCN,
				NWSLI: v.NWSLI,
				Elevation: v.Elevation,
			}
		}
		err = autocompleteLocal.AddSources(sts)
		if err != nil {
			level.Error(logger).Log("msg", "AddAutocomplete error", "err", err)
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

	for _,fileName := range filesList {
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