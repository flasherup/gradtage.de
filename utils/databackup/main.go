package main

import (
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/utils/databackup/config"
	"github.com/flasherup/gradtage.de/utils/databackup/database"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
	"sync"
	"time"
)
//"82.165.119.83:8111"
type WBDataBackup struct {
	StName string
	Data []weatherbitsvc.WBData
}

func main() {
	configFile := flag.String("config.file", "config.yml", "Config file name.")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger)
	}

	//Config
	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		level.Error(logger).Log("msg", "config loading error", "exit", err.Error())
		return
	}

	db, err := database.NewPostgres(conf.Database)
	if err != nil {
		level.Error(logger).Log("msg", "database error", "exit", err.Error())
		return
	}

	source := impl.NewWeatherBitSVCClient(conf.Clients.SRCAddr, logger)

	var wg  sync.WaitGroup
	wg.Add(2)

	level.Info(logger).Log("msg", "Data backup client start")
	defer level.Info(logger).Log("msg", "client end")

	ch  := make(chan WBDataBackup)


	go moveData(db, source, logger, ch, &wg)
	go pushData(db,logger, ch, &wg)

	wg.Wait()
	level.Info(logger).Log("msg", "The program has finished")
}

func moveData(db *database.Postgres, source weatherbitsvc.Client, logger log.Logger, ch chan WBDataBackup, wg *sync.WaitGroup) {
	//fmt.Print(ch)
	level.Info(logger).Log("msg", "Getting data")
	//currentTime := time.Now()
	stations, err := source.GetStationsList()
	if err != nil {
		level.Error(logger).Log("msg", "Getting stations list error", "error", err.Error())
		return
	}

	level.Info(logger).Log("msg", "Getting stations list success", "stations count", len(*stations))

	date := time.Now()
	currentDate := date.Format(common.TimeLayoutWBH)

	for i, stationName := range *stations {

		date2 := time.Now()

		level.Info(logger).Log("msg", "Process station: "+stationName, "#", i)
		data, err := source.GetWBPeriod(stationName, "2000-01-01T00:00:00", currentDate)
		if err != nil {

			level.Error(logger).Log("msg", "Getting station data error", "station", stationName, "error", err.Error())
			continue
		}

		ch <- WBDataBackup{
			stationName,
			  *data,
		}

		err = db.CreateTable(stationName)
		if err != nil {
			level.Error(logger).Log("msg", "table create error", "err", err)
		}
		fmt.Println("GetWBPeriod ", stationName, "complete. Time elapsed:", time.Since(date2).Seconds(), "ms")
	}
}

func pushData(db *database.Postgres, logger log.Logger, ch chan WBDataBackup, wg *sync.WaitGroup) {

	for i := 0; i < 10000; i++ {
		select {
		case wd := <-ch:
			date := time.Now()
			level.Info(logger).Log("msg", "Station: "+wd.StName+" data received")
			err := db.PushWBData(wd.StName, wd.Data)
			if err != nil {
				level.Error(logger).Log("msg", "Saving station data error", "station", wd.StName, "error", err.Error())
			}
			fmt.Println("PushWBPeriod ", wd.StName, "complete. Time elapsed:", time.Since(date).Seconds(), "ms")
		}
	}
	defer wg.Done()
}


