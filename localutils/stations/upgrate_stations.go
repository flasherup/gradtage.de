package main

import (
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"

	stations "github.com/flasherup/gradtage.de/stationssvc/impl"
)

func main() {
	//addNewToLocal(data.MeteastatStation)
	//fromLocalToRemote()

	fromRemoteToLocal()

	//filterLocal()

	//fixDwdIDs()
	//fixAutocompleteIDs()
}

func addNewToLocal(sts []stationssvc.Station) {
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

	stationsLocal := stations.NewStationsSCVClient("localhost:8102", logger)
	_, err := stationsLocal.AddStations(sts)
	if err != nil {
		level.Error(logger).Log("msg", "ResetStations error", "err", err)
	}
}

func filterLocal() {
	filterIds := []string{
		"DE06256",
		"DE06245",
		"WMO10500",
		"DE06242",
		"DE06247",
		"DE00398",
		"DE00397",
		"DE06246",
		"DE06244",
		"DE00396",
		"DE06243",
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "stations_remove",
			"ts", log.DefaultTimestampUTC,
			"caller", log.Caller(3),
		)
	}

	stations := stations.NewStationsSCVClient("localhost:8102", logger)

	sts, err := stations.GetAllStations()
	if err != nil {
		level.Error(logger).Log("msg", "GetStations error", "err", err)
		return
	}

	level.Info(logger).Log("msg", "Current number of stations", "num", len(sts.Sts))

	nsts := make([]stationssvc.Station, 0)

	for k, v := range sts.Sts {
		if indexOf(k, filterIds) > -1 {
			level.Info(logger).Log("msg", "Filtered station", "id", k)
			continue
		}
		nsts = append(nsts,
			stationssvc.Station{
				ID:         k,
				Name:       v.Name,
				Timezone:   v.Timezone,
				SourceType: v.SourceType,
				SourceID:   v.SourceId,
			})
	}

	level.Info(logger).Log("msg", "Filtered number of stations", "num", len(nsts))

	_, err = stations.ResetStations(nsts)
	if err != nil {
		level.Error(logger).Log("msg", "ResetStations error", "err", err)
	}
}

func fromRemoteToLocal() {
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

	stationsRemout := stations.NewStationsSCVClient("212.227.215.17:8102", logger)

	sts, err := stationsRemout.GetAllStations()
	if err != nil {
		level.Error(logger).Log("msg", "GetStations error", "err", err)
		return
	}

	stsl := make([]stationssvc.Station, len(sts.Sts))

	i := 0
	for k, v := range sts.Sts {
		stsl[i] = stationssvc.Station{
			ID:         k,
			Name:       v.Name,
			Timezone:   v.Timezone,
			SourceType: v.SourceType,
			SourceID:   v.SourceId,
		}
		i++
	}

	for k, v := range stsl {
		fmt.Println(k, v)
	}

	stationsLocal := stations.NewStationsSCVClient("localhost:8102", logger)
	_, err = stationsLocal.ResetStations(stsl)
	if err != nil {
		level.Error(logger).Log("msg", "AddStations error", "err", err)
	}

	level.Info(logger).Log("msg", "Stations updated")
}

func fromLocalToRemote() {
	stationsDWD := map[string]string{}

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

	stationsRemout := stations.NewStationsSCVClient("localhost:8102", logger)

	sts, err := stationsRemout.GetAllStations()
	if err != nil {
		level.Error(logger).Log("msg", "GetStations error", "err", err)
		return
	}

	stsl := make([]stationssvc.Station, len(sts.Sts))

	i := 0
	for k, v := range sts.Sts {
		fmt.Println(k, v)
		if id, ok := stationsDWD[k]; ok {
			stsl[i] = stationssvc.Station{
				ID:         k,
				Name:       v.Name,
				Timezone:   v.Timezone,
				SourceType: common.SrcTypeDWD,
				SourceID:   id,
			}

		} else {
			stsl[i] = stationssvc.Station{
				ID:         k,
				Name:       v.Name,
				Timezone:   v.Timezone,
				SourceType: v.SourceType,
				SourceID:   v.SourceId,
			}
		}
		i++
	}

	for k, v := range stsl {
		fmt.Println(k, v)
	}

	stationsLocal := stations.NewStationsSCVClient("82.165.18.228:8102", logger)
	_, err = stationsLocal.ResetStations(stsl)
	if err != nil {
		level.Error(logger).Log("msg", "AddStations error", "err", err)
	}
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}

	return -1
}
