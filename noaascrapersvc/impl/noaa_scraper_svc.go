package impl

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/noaascrapersvc/config"
	"github.com/flasherup/gradtage.de/noaascrapersvc/impl/database"
	"github.com/flasherup/gradtage.de/noaascrapersvc/impl/parser"
	"github.com/flasherup/gradtage.de/noaascrapersvc/impl/source"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
	"time"
)

type NOAAScraperSVC struct {
	stations    stationssvc.Client
	db 			database.NoaaDB
	alert 		alertsvc.Client
	logger  	log.Logger
	counter 	*ktprom.Gauge
	src			source.SourceNOAA
}

const (
	labelStations = "stations"
	labelTemperature = "temperature"
	labelStationError = "station_error"
)

func NewNOAAScraperSVC(
		logger 		log.Logger,
		stations 	stationssvc.Client,
		db 			database.NoaaDB,
		alert 		alertsvc.Client,
		conf 		config.NOAAScraperConfig,
	) (*NOAAScraperSVC, error) {
	options := prometheus.Opts{
		Name: "noaa_stations_update_count",
		Help: "The total number of NOAA stations that was updated",
	}

	src := 	source.NewNOAA(conf.Sources.UrlNoaa, logger)
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ labelStations, labelStationError })
	st := NOAAScraperSVC{
		stations: stations,
		db: db,
		alert: alert,
		logger: logger,
		counter: guage,
		src: *src,
	}
	go startFetchProcess(&st)
	return &st,nil
}

func (hs NOAAScraperSVC) GetPeriod(ctx context.Context, id string, start string, end string) (temps []hourlysvc.Temperature, err error) {
	level.Info(hs.logger).Log("msg", "GetPeriod", "ids", fmt.Sprintf("%s: %s-%s",id, start, end))
	temps, err = hs.db.GetPeriod(id, start, end)
	if err != nil {
		level.Error(hs.logger).Log("msg", "GetPeriod error", "err", err)
		hs.sendAlert(NewErrorAlert(err))
	}
	return temps,err
}

func (hs *NOAAScraperSVC) GetUpdateDate(ctx context.Context, ids []string) (dates map[string]string, err error) {
	level.Info(hs.logger).Log("msg", "GetUpdateDate", "ids", fmt.Sprintf("%+q:",ids))
	dates = make(map[string]string)
	for _,v := range ids {
		date, err := hs.db.GetUpdateDate(v)
		if err != nil {
			level.Error(hs.logger).Log("msg", "Get Update Date error", "err", err)
			hs.sendAlert(NewErrorAlert(err))
		} else {
			dates[v] = date
		}
	}

	return dates, err
}


func startFetchProcess(ss *NOAAScraperSVC) {
	ss.processUpdate(-1) //Do it first time
	tick := time.Tick(time.Hour)
	for {
		select {
		case <-tick:
			ss.processUpdate(3)
		}
	}
}


func (das NOAAScraperSVC)processUpdate(rowsNumber int) {
	sts, err := das.stations.GetStationsBySrcType([]string{ common.SrcTypeNOAA })
	if err != nil {
		level.Error(das.logger).Log("msg", "Get NOAA Stations error", "err", err)
		das.sendAlert(NewErrorAlert(err))
		return
	}

	ids := make(map[string]string)
	for k,v := range sts.Sts {
		ids[k] = strings.ToUpper(v.SourceId)
	}

	ch := make(chan *parser.ParsedData)
	go das.src.FetchTemperature(ch, ids)

	count := 0.0
	for range ids {
		pd := <-ch
		sErr := das.counter.With(labelStations, pd.StationID, labelStationError, "update_fail")
		sErr.Set(0)
		if pd != nil && pd.Success && len(pd.Temps) > 0 {
			rowsToUpdate := pd.Temps
			if rowsNumber > 0 {
				rowsToUpdate = rowsToUpdate[len(rowsToUpdate)-rowsNumber:]
			}
			err := das.db.CreateTable(pd.StationID)
			if err != nil {
				level.Error(das.logger).Log("msg", "Create station Error", "err", err)
				das.sendAlert(NewErrorAlert(err))
				sErr.Set(1)
				continue
			}

			err = das.db.PushPeriod(pd.StationID, rowsToUpdate)
			if err != nil {
				level.Error(das.logger).Log("msg", "PushPeriod Error", "err", err)
				das.sendAlert(NewErrorAlert(err))
				sErr.Set(1)
				continue
			}

			if len(pd.Temps) > 0 {
				g := das.counter.With(labelStations, pd.StationID, labelStationError, "")
				g.Set(pd.Temps[len(pd.Temps)-1].Temperature)
			}

			level.Info(das.logger).Log("msg", "Station updated", "id", pd.StationID, "temp", fmt.Sprintf("%+q:",pd.Temps))
			count++

		} else {
			sErr.Set(1)
			if pd != nil {
				level.Error(das.logger).Log("msg", "Station update error", "err", pd.Error)
			}
			level.Warn(das.logger).Log("msg", "Station is not updated", "reason", pd.Error)
		}
	}

	g := das.counter.With(labelStations, "all", labelStationError, "")
	g.Set(count)
	level.Info(das.logger).Log("msg", "Temperature updated from NOAA", "stations", count)
}


func (das NOAAScraperSVC)sendAlert(alert alertsvc.Alert) {
	err := das.alert.SendAlert(alert)
	if err != nil {
		level.Error(das.logger).Log("msg", "Send Alert Error", "err", err)
	}
}


