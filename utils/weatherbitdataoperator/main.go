package main

import (
	"flag"
	"github.com/flasherup/gradtage.de/utils/weatherbitdataoperator/config"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

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

	source := impl.NewWeatherBitSVCClient(conf.Clients.SRCAddr,logger)
	receiver := impl.NewWeatherBitSVCClient(conf.Clients.RCVRAddr,logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	moveData(source, receiver, logger)



}

func moveData(source, receiver weatherbitsvc.Client, logger log.Logger) {
	level.Info(logger).Log("msg", "Getting stations list")

	stations, err := source.GetStationsList()
	if err != nil {
		level.Error(logger).Log("msg", "Getting stations list error", "error", err.Error())
		return
	}

	level.Info(logger).Log("msg", "Getting stations list success", "stations count", len(*stations))

	for i,stName := range *stations {
		level.Info(logger).Log("msg", "Process station: " + stName, "#", i)
		data, err := source.GetWBPeriod(stName, "2000-01-01T00:00:00", "2021-08-01T00:00:00")
		if err != nil {
			level.Error(logger).Log("msg", "Getting station data error","station", stName, "error", err.Error())
			continue
		}
		level.Info(logger).Log("msg", "Station: " + stName + " data received", "count", len(*data))
		err = receiver.PushWBPeriod(stName, *data)
		if err != nil{
			level.Error(logger).Log("msg", "Saving station data error","station", stName, "error", err.Error())
		}
	}
}


