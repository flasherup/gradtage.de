package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type StationsSVC struct {
	logger  log.Logger
}

func NewStationsSVC(logger log.Logger) *StationsSVC {
	return &StationsSVC{}
}

func (ss StationsSVC) GetStations(ctx context.Context, ids []string) (sts map[string]stationssvc.Station, err error) {
	level.Info(ss.logger).Log("msg", "GetStations")
	sts = make(map[string]stationssvc.Station)
	return
}

func (ss StationsSVC) GetAllStations(ctx context.Context) (sts map[string]stationssvc.Station, err error){
	level.Info(ss.logger).Log("msg", "GetAllStations")
	sts = make(map[string]stationssvc.Station)
	return
}

func (ss StationsSVC) AddStations(ctx context.Context, sts []stationssvc.Station) (err error) {
	level.Info(ss.logger).Log("msg", "AddStations")
	return
}