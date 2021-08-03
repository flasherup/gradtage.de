package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/utils/weatherbitcheck/config"
	"github.com/flasherup/gradtage.de/utils/weatherbitcheck/database"
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

	getPeriod(db, logger)

}

func getPeriod(db *database.Postgres, logger log.Logger) {
	period, err := db.GetPeriod("us_kduj", "2020-03-20T00:00:00", "2021-03-25T20:00:00")
	if err != nil {
		fmt.Println("get period error", err)
	} else {
		for i,v := range period {
			println(i, v.Date, v.Temperature)
		}
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

