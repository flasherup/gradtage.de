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
		"EDG_Stationlist_Masterfile.csv",
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

 	fromCSVListToStations("./data", filesList, logger)
	//fromCSVListToAutocomplete("data", filesList, logger)
	//fromCSVToList("./data", filesList, logger)
}


func fromCSVListToStations(path string, filesList []string, logger log.Logger) {
	stationsLocal := stations.NewStationsSCVClient("localhost:8102", logger)

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


			/*fmt.Println("=========================", i)
			fmt.Println("ID",v.ID)
			fmt.Println("SourceID",v.SourceID)
			fmt.Println("Latitude",v.Latitude)
			fmt.Println("Longitude",v.Longitude)
			fmt.Println("Source",v.Source)
			fmt.Println("Reports",v.Reports)
			fmt.Println("ISO2Country",v.ISO2Country)
			fmt.Println("ISO3Country",v.ISO3Country)
			fmt.Println("Prio",v.Prio)
			fmt.Println("CityNameEnglish",v.CityNameEnglish)
			fmt.Println("CityNameNative",v.CityNameNative)
			fmt.Println("CountryNameEnglish",v.CountryNameEnglish)
			fmt.Println("CountryNameNative",v.CountryNameNative)
			fmt.Println("ICAO",v.ICAO)
			fmt.Println("WMO",v.WMO)
			fmt.Println("CWOP",v.CWOP)
			fmt.Println("Maslib",v.Maslib)
			fmt.Println("National_ID",v.National_ID)
			fmt.Println("IATA",v.IATA)
			fmt.Println("USAF_WBAN",v.USAF_WBAN)
			fmt.Println("GHCN",v.GHCN)
			fmt.Println("NWSLI",v.NWSLI)
			fmt.Println("Elevation",v.Elevation)*/
		}
		_,err = stationsLocal.AddStations(sts)
		if err != nil {
			level.Error(logger).Log("msg", "AddStations error", "err", err)
		}

		//allStation = append(allStation, sts...)
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