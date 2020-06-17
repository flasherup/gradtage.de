package impl

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/stationssvc/impl/database"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

type StationsSVC struct {
	logger  	log.Logger
	alert 		alertsvc.Client
	db 			database.StationsDB
	counter 	*ktprom.Gauge
}

func NewStationsSVC(logger log.Logger, db database.StationsDB, alert alertsvc.Client) (*StationsSVC, error) {
	err := db.CreateTable()
	if err != nil {
		return nil, err
	}

	options := prometheus.Opts{
		Name: "stations_count_total",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := StationsSVC{
		logger: logger,
		alert:  alert,
		db:		db,
		counter: guage,
	}
	st.updateStationsMetrics()
	return &st,nil
}

func (ss StationsSVC) GetStations(ctx context.Context, ids []string) (sts map[string]stationssvc.Station, err error) {
	level.Info(ss.logger).Log("msg", "GetStations", "ids", fmt.Sprintf("%+q",ids))
	stations, err := ss.db.GetStations(ids)
	if err != nil {
		level.Error(ss.logger).Log("msg", "GetStations error", "err", err)
		ss.sendAlert(NewErrorAlert(err))
		return nil,err
	}
	sts = make(map[string]stationssvc.Station)
	for _,v := range stations {
		sts[v.ID] = v
	}
	return
}

func (ss StationsSVC) GetAllStations(ctx context.Context) (sts map[string]stationssvc.Station, err error){
	level.Info(ss.logger).Log("msg", "GetAllStations")
	stations, err := ss.db.GetAllStations()
	if err != nil {
		level.Error(ss.logger).Log("msg", "GetAllStations error", "err", err)
		ss.sendAlert(NewErrorAlert(err))
		return nil,err
	}
	sts = make(map[string]stationssvc.Station)
	for _,v := range stations {
		sts[v.ID] = v
	}
	return
}

func (ss StationsSVC) GetStationsBySrcType(ctx context.Context,  types []string) (sts map[string]stationssvc.Station, err error){
	level.Info(ss.logger).Log("msg", "GetStationsBySrcType", "types", fmt.Sprintf("%+q",types))
	stations, err := ss.db.GetStationsBySrcType(types)
	if err != nil {
		level.Error(ss.logger).Log("msg", "GetStationsBySrcType error", "err", err)
		ss.sendAlert(NewErrorAlert(err))
		return nil,err
	}
	sts = make(map[string]stationssvc.Station)
	for _,v := range stations {
		sts[v.ID] = v
	}
	return
}

func (ss *StationsSVC) AddStations(ctx context.Context, sts []stationssvc.Station) (err error) {
	level.Info(ss.logger).Log("msg", "AddStations")
	err = ss.db.AddStations(sts)
	if err != nil {
		level.Error(ss.logger).Log("msg", "AddStations error", "err", err)
		ss.sendAlert(NewErrorAlert(err))
	}
	ss.updateStationsMetrics()
	return
}

func (ss *StationsSVC) ResetStations(ctx context.Context, sts []stationssvc.Station) (err error) {
	level.Info(ss.logger).Log("msg", "ResetStations")
	err = ss.db.RemoveTable()
	if err != nil {
		level.Error(ss.logger).Log("msg", "RemoveTable error", "err", err)
		return
	}
	err = ss.db.CreateTable()
	if err != nil {
		level.Error(ss.logger).Log("msg", "CreateTable error", "err", err)
		return
	}
	err = ss.db.AddStations(sts)
	if err != nil {
		level.Error(ss.logger).Log("msg", "AddStations error", "err", err)
		ss.sendAlert(NewErrorAlert(err))
	}
	ss.updateStationsMetrics()
	return
}

func (ss *StationsSVC)updateStationsMetrics() {
	count, err :=  ss.db.GetCount()
	if err == nil {
		g := ss.counter.With("stations")
		g.Set(float64(count))
		level.Info(ss.logger).Log("msg", "Stations count updated", "count", count)

	} else {
		level.Error(ss.logger).Log("msg", "Stations count update error", "err", err)
		ss.sendAlert(NewErrorAlert(err))
	}
}

func (ss StationsSVC)sendAlert(alert alertsvc.Alert) {
	err := ss.alert.SendAlert(alert)
	if err != nil {
		level.Error(ss.logger).Log("msg", "SendAlert Alert Error", "err", err)
	}
}