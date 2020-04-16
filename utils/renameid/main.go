package main

import (
	"flag"
	"github.com/flasherup/gradtage.de/dailysvc"
	daily "github.com/flasherup/gradtage.de/dailysvc/impl"
	"github.com/flasherup/gradtage.de/hourlysvc"
	hourly "github.com/flasherup/gradtage.de/hourlysvc/impl"
	"github.com/flasherup/gradtage.de/stationssvc"
	stations "github.com/flasherup/gradtage.de/stationssvc/impl"
	"github.com/flasherup/gradtage.de/utils/renameid/config"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
	"time"
)

const veryFirstTime = "2000-01-02T01:01:01.00Z"

func main() {
	configFile := flag.String("config.file", "config.yml", "Config file name.")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "Rename ID Util",
			"ts", log.DefaultTimestampUTC,
			"caller", log.Caller(3),
		)
	}

	level.Info(logger).Log("msg", "Start Renaming IDs")

	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		level.Error(logger).Log("msg", "Config Loading error", "err", err)
		return
	}

	renameHourly(logger, conf.Services.HourlyUrl, conf.Renames)
	renameDaily(logger, conf.Services.DailyUrl, conf.Renames)
	renameStations(logger, conf.Services.StationsUrl, conf.Renames)
}

func renameHourly(logger log.Logger, url string, ids []config.Rename) {
	level.Info(logger).Log("msg", "Rename Hourly IDs")
	hourlyService := hourly.NewHourlySCVClient(url, logger)
	currentTime := time.Now().Format(veryFirstTime)
	count := 0
	for _,v := range ids {
		level.Info(logger).Log("msg", "Rename Hourly ID", "current", v.CurrentID, "new", v.NewID)
		temps, err := hourlyService.GetPeriod(v.CurrentID, veryFirstTime, currentTime)
		if err != nil {
			level.Error(logger).Log("msg", "Hourly Get Period error", "err", err)
		} else if temps.Err != "nil" {

			level.Error(logger).Log("msg", "Hourly Get Period error", "err", temps.Err)
		} else {
			if len(temps.Temps) > 0 {
				level.Info(logger).Log("msg", "Rename Hourly Period", "start", temps.Temps[0].Date, "end", temps.Temps[len(temps.Temps)-1].Date)
			} else {
				level.Info(logger).Log("msg", "Rename Hourly Period is empty", "start")
			}
			t := hourlysvc.DecodeTemperature(temps.Temps)
			hourlyService.PushPeriod(v.NewID, t)
			count++
		}
	}
	level.Info(logger).Log("msg", "Rename Hourly IDs Complete", "updated", count)
}

func renameDaily(logger log.Logger, url string, ids []config.Rename) {
	level.Info(logger).Log("msg", "Rename Daily IDs")
	dailyService := daily.NewDailySCVClient(url, logger)
	currentTime := time.Now().Format(veryFirstTime)
	count := 0
	for _,v := range ids {
		level.Info(logger).Log("msg", "Rename Daily ID", "current", v.CurrentID, "new", v.NewID)
		temps, err := dailyService.GetPeriod(v.CurrentID, veryFirstTime, currentTime)
		if err != nil {
			level.Error(logger).Log("msg", "Daily Get Period error", "err", err)
		} else if temps.Err != "nil" {

			level.Error(logger).Log("msg", "Daily Get Period error", "err", temps.Err)
		} else {
			if len(temps.Temps) > 0 {
				level.Info(logger).Log("msg", "Rename Daily Period", "start", temps.Temps[0].Date, "end", temps.Temps[len(temps.Temps)-1].Date)
			} else {
				level.Info(logger).Log("msg", "Rename Daily Period is empty", "start")
			}
			t := dailysvc.DecodeTemperature(temps.Temps)
			dailyService.PushPeriod(v.NewID, t)
			count++
		}
	}
	level.Info(logger).Log("msg", "Rename Daily IDs Complete", "updated", count)
}

func renameStations(logger log.Logger, url string, ids []config.Rename) {
	level.Info(logger).Log("msg", "Rename Stations IDs")
	stationsService := stations.NewStationsSCVClient(url, logger)

	sts, err := stationsService.GetAllStations()
	if err != nil {
		level.Error(logger).Log("msg", "GetStations error", "err", err)
		return
	}

	stationsToRename := make(map[string]string)

	for _,v := range ids {
		stationsToRename[v.CurrentID] = v.NewID
	}

	stsl := make([]stationssvc.Station, len(sts.Sts))
	i := 0
	count := 0
	for k,v := range sts.Sts {
		if id, ok:= stationsToRename[k]; ok {
			count++
			level.Info(logger).Log("msg", "Rename Station ID", "current", k, "new", id)
			stsl[i] = stationssvc.Station{
				ID:id,
				Name:v.Name,
				Timezone:v.Timezone,
				SourceType:v.SourceType,
				SourceID:id,
			}

		} else {
			stsl[i] = stationssvc.Station{
				ID:k,
				Name:v.Name,
				Timezone:v.Timezone,
				SourceType:v.SourceType,
				SourceID:v.SourceId,
			}
		}
		i++
	}

	stationsLocal := stations.NewStationsSCVClient("localhost:8102", logger)
	_,err = stationsLocal.ResetStations(stsl)
	if err != nil {
		level.Error(logger).Log("msg", "AddStations error", "err", err)
	} else {

		level.Info(logger).Log("msg", "Rename Stations IDs Complete", "updated", count)
	}
}