package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/noaascrapersvc/config"
	"github.com/flasherup/gradtage.de/noaascrapersvc/impl/database"
	"github.com/flasherup/gradtage.de/noaascrapersvc/impl/parser"
	"github.com/flasherup/gradtage.de/noaascrapersvc/impl/source"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
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

func NewNOAAScraperSVC(
		logger 		log.Logger,
		stations 	stationssvc.Client,
		db 			database.NoaaDB,
		alert 		alertsvc.Client,
		conf config.NOAAScraperConfig,
	) (*NOAAScraperSVC, error) {
	options := prometheus.Opts{
		Name: "noaa_stations_update_count",
		Help: "The total number of NOAA stations that was updated",
	}


	src := 	source.NewNOAA(conf.Sources.UrlNoaa, logger)
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
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

func (das NOAAScraperSVC) ForceOverrideHourly(ctx context.Context, station string, start string, end string) error {
	level.Info(das.logger).Log("msg", "ForceOverrideHourly", "station", station, "start", start, "end", end)
	return nil
}


func startFetchProcess(ss *NOAAScraperSVC) {
	ss.processUpdate(1) //Do it first time
	tick := time.Tick(time.Hour)
	for {
		select {
		case <-tick:
			ss.processUpdate(1)
		}
	}
}


func (das NOAAScraperSVC)processUpdate(rowsNumber int) {
	/*sts, err := das.stations.GetStationsBySrcType([]string{ common.SrcTypeNOAA })
	if err != nil {
		level.Error(das.logger).Log("msg", "Get DWD Stations error", "err", err)
		das.sendAlert(NewErrorAlert(err))
		return
	}*/

	/*ids := make(map[string]string)
	for k,v := range sts.Sts {
		ids[k] = v.SourceId
	}*/

	ids := map[string]string{"ebbr":"ebbr"}

	/*latest, err := has.hourly.GetLatest(ids)
	if err != nil {
		level.Error(has.logger).Log("msg", "Get latest error", "err", err)
		has.sendAlert(NewErrorAlert(err))
	}*/

	ch := make(chan *parser.ParsedData)
	go das.src.FetchTemperature(ch, ids)

	count := 0.0
	for range ids {
		pd := <-ch
		if pd != nil && pd.Success {
			//has.verifyPlausibility(latest, pd.StationID, pd.Temps)
			rowsToUpdate := pd.Temps
			/*if rowsNumber > 0 {
				rowsToUpdate = rowsToUpdate[len(rowsToUpdate)-rowsNumber:]
			}*/
			err := das.db.PushPeriod(pd.StationID, rowsToUpdate)
			if err != nil {
				level.Error(das.logger).Log("msg", "PushPeriod Error", "err", err)
				das.sendAlert(NewErrorAlert(err))
			} else {
				count++
			}
		} else {
			if pd != nil {
				level.Error(das.logger).Log("msg", "Station update error", "err", pd.Error)
			}
			level.Warn(das.logger).Log("msg", "Station is not updated")
		}
	}

	g := das.counter.With("stations")
	g.Set(count)
	level.Info(das.logger).Log("msg", "Temperature updated from NOAA", "stations", count)
}


func (das NOAAScraperSVC)sendAlert(alert alertsvc.Alert) {
	err := das.alert.SendAlert(alert)
	if err != nil {
		level.Error(das.logger).Log("msg", "Send Alert Error", "err", err)
	}
}


