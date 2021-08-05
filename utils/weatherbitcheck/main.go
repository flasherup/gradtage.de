package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/utils/weatherbitcheck/config"
	"github.com/flasherup/gradtage.de/utils/weatherbitcheck/database"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
	"strconv"
)

func main() {
	configFile := flag.String("config.file", "config.yml", "Config file name. ")
	//csvFile := flag.String("csv.file", "stationsR.csv", "CSV file name. ")
	flag.Parse()

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowDebug())

	//Config
	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		level.Error(logger).Log("msg", "config loading error", "err", err.Error())
		return
	}

	db, err := database.NewPostgres(conf.Database)
	if err != nil {
		level.Error(logger).Log("msg", "database error", "exit", err.Error())
		return
	}

	//getPeriod(db, logger)
	pushPeriod(db, logger);

}

func getPeriod(db *database.Postgres, logger log.Logger) {
	period, err := db.GetWBData("us_kduj", "2020-03-25T00:00:00", "2021-03-25T01:00:00")
	if err != nil {
		fmt.Println("get period error", err)
	} else {
		for _,v := range period {
			fmt.Println("Date:",v.Date,",")
			fmt.Println("Rh:",v.Rh,",")
			fmt.Println("Pod:",v.Pod,",")
			fmt.Println("Pres:",v.Pres,",")
			fmt.Println("Timezone:",v.Timezone,",")
			fmt.Println("CountryCode:",v.CountryCode,",")
			fmt.Println("Clouds:",v.Clouds,",")
			fmt.Println("Vis:",v.Vis,",")
			fmt.Println("SolarRad:",v.SolarRad,",")
			fmt.Println("WindSpd:",v.WindSpd,",")
			fmt.Println("StateCode:",v.StateCode,",")
			fmt.Println("CityName:",v.CityName,",")
			fmt.Println("AppTemp:",v.AppTemp,",")
			fmt.Println("Uv:",v.Uv,",")
			fmt.Println("Lon:",v.Lon,",")
			fmt.Println("Slp:",v.Slp,",")
			fmt.Println("HAngle:",v.HAngle,",")
			fmt.Println("Dewpt:",v.Dewpt,",")
			fmt.Println("Snow:",v.Snow,",")
			fmt.Println("Aqi:",v.Aqi,",")
			fmt.Println("WindDir:",v.WindDir,",")
			fmt.Println("ElevAngle:",v.ElevAngle,",")
			fmt.Println("Ghi:",v.Ghi,",")
			fmt.Println("Lat:",v.Lat,",")
			fmt.Println("Precip:",v.Precip,",")
			fmt.Println("Sunset:",v.Sunset,",")
			fmt.Println("Temp:",v.Temp,",")
			fmt.Println("Station:",v.Station,",")
			fmt.Println("Dni:",v.Dni,",")
			fmt.Println("Sunrise:",v.Sunrise,",")
			return
		}
	}
}

func pushPeriod(db *database.Postgres, logger log.Logger) {
	data := weatherbitsvc.WBData{
		Date: "2020-03-25T00:00:00Z" ,
		Rh: 67 ,
		Pod: "n" ,
		Pres: 952 ,
		Timezone: "America/New_York" ,
		CountryCode: "US" ,
		Clouds: 0 ,
		Vis: 16 ,
		SolarRad: 0 ,
		WindSpd: 6.8 ,
		StateCode: "PA" ,
		CityName: "Allens Mills" ,
		AppTemp: 2 ,
		Uv: 0 ,
		Lon: -78.9 ,
		Slp: 1018.3 ,
		HAngle: 0 ,
		Dewpt: 0.6 ,
		Snow: 0 ,
		Aqi: 0 ,
		WindDir: 110 ,
		ElevAngle: -5.96 ,
		Ghi: 0 ,
		Lat: 41.18 ,
		Precip: 0 ,
		Sunset: "" ,
		Temp: 6.1 ,
		Station: "" ,
		Dni: 0 ,
		Sunrise: "" ,
	}

	stID := "us_kduj";

	err := db.CreateTable(stID)
	if err != nil {
		level.Error(logger).Log("msg", "table create error", "err", err)
	}

	err = db.PushWBData(stID, []weatherbitsvc.WBData{data})
	if err != nil {
		level.Error(logger).Log("msg", "push wb period error", "err", err.Error())
	}
}

func getStationsData(db *database.Postgres, logger log.Logger) {
	level.Info(logger).Log("msg", "Get stations list")
	stationList, err := db.GetListOfTables()
	if err != nil {
		level.Error(logger).Log("msg", "CSV loading error", "err", err.Error())
		return
	}

	level.Info(logger).Log("msg", "Stations list get success", "length", len(stationList))


	//File save logic
	csvFile, err := os.Create("stationsICAO.csv")

	if err != nil {
		fmt.Printf("failed creating file: %s", err.Error())
	}

	csvwriter := csv.NewWriter(csvFile)

	defer csvwriter.Flush()
	defer csvFile.Close()

	for i,v := range stationList {
		count, err := db.CountTableRows(v)
		if err != nil {
			fmt.Println(i, v, "Error:", err)
		} else {
			fmt.Println(i, v, count)
		}

		err = csvwriter.Write([]string{strconv.Itoa(i), v, strconv.Itoa(count) })
		if err != nil {
			fmt.Println(i, v, "Error:", err)
		}
	}
}

