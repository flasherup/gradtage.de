package main

import (
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/utils/databackup/config"
	"github.com/flasherup/gradtage.de/utils/databackup/database"
	"github.com/flasherup/gradtage.de/utils/databackup/ui"
	"github.com/flasherup/gradtage.de/utils/databackup/ui/mainwindow"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
	"time"
)

type DataBackup struct {
	app ui.Application
	logger log.Logger
	config config.DataBackupConfig
	postgresDB database.Postgres
	stations []string
}

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger)
	}

	dBack := DataBackup{
		app: *ui.NewApplication(),
		logger: logger,
	}

	dBack.initConfig()

	db, err := database.NewPostgres(dBack.config.Database)
	if err != nil {
		level.Error(logger).Log("msg", "database error", "exit", err.Error())
		return
	}

	dBack.postgresDB = *db

	dBack.app.WMain.Bus.Subscribe(mainwindow.OnStationsLoad, dBack.getStations)
	dBack.app.WMain.Bus.Subscribe(mainwindow.OnBackupStart, dBack.startBackupData)

	dBack.app.WMain.Window.Show()
	dBack.app.App.Run()
}

func (dBack *DataBackup)initConfig() {
	//Config
	conf, err := config.LoadConfig("config.yml")
	if err != nil {
		level.Error(dBack.logger).Log("msg", "config loading error", "exit", err.Error())
		return
	}

	dBack.config = *conf
}

func (dBack *DataBackup) getStations() {
	source := impl.NewWeatherBitSVCClient(dBack.config.Clients.SRCAddr, dBack.logger)

	stations, err := source.GetStationsList()
	if err != nil {
		level.Error(dBack.logger).Log("msg", "Getting stations list error", "error", err.Error())
		return
	}

	level.Info(dBack.logger).Log("msg", "Getting stations list success", "stations count", len(*stations))
	dBack.stations = *stations
	dBack.app.WMain.SetStationsLength(len(*stations))
}


func (dBack *DataBackup) startBackupData() {
	source := impl.NewWeatherBitSVCClient(dBack.config.Clients.SRCAddr, dBack.logger)

	date := time.Now()
	currentDate := date.Format(common.TimeLayoutWBH)

	for i, stationName := range dBack.stations {
		level.Info(dBack.logger).Log("msg", "Process station: "+stationName, "#", i)
		data, err := source.GetWBPeriod(stationName, "2000-01-01T00:00:00", currentDate)
		if err != nil {
			level.Error(dBack.logger).Log("msg", "Getting station data error", "station", stationName, "error", err.Error())
			continue
		}

		err = dBack.postgresDB.CreateTable(stationName)
		if err != nil {
			level.Error(dBack.logger).Log("msg", "table create error", "err", err)
		}


		level.Info(dBack.logger).Log("msg", "Station: "+stationName+" data received", "count", len(*data))
		err = dBack.postgresDB.PushWBData(stationName, *data)
		if err != nil {
			level.Error(dBack.logger).Log("msg", "Saving station data error", "station", stationName, "error", err.Error())
		}
	}
}
