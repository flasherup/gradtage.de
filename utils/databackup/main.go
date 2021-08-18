package main

import (
	"flag"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/utils/databackup/database"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()

}

func runBackup() {
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

	level.Info(logger).Log("msg", "Data backup client start")
	defer level.Info(logger).Log("msg", "client end")



	moveData(db, source, logger)
}

func moveData(db *database.Postgres, source weatherbitsvc.Client, logger log.Logger)  {
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
		level.Info(logger).Log("msg", "Process station: "+stationName, "#", i)
		data, err := source.GetWBPeriod(stationName, "2000-01-01T00:00:00", currentDate)
		if err != nil {
			level.Error(logger).Log("msg", "Getting station data error", "station", stationName, "error", err.Error())
			continue
		}

		err = db.CreateTable(stationName)
		if err != nil {
			level.Error(logger).Log("msg", "table create error", "err", err)
		}


		level.Info(logger).Log("msg", "Station: "+stationName+" data received", "count", len(*data))
		err = db.PushWBData(stationName, *data)
		if err != nil {
			level.Error(logger).Log("msg", "Saving station data error", "station", stationName, "error", err.Error())
		}
	}


}
