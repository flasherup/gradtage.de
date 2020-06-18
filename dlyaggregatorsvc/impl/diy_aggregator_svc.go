package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/impl/parser"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/impl/source"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type DailyAggregatorSVC struct {
	stations    stationssvc.Client
	src 		*source.Hourly
	daily 		dailysvc.Client
	alert 		alertsvc.Client
	logger  	log.Logger
	counter 	*ktprom.Gauge
}

func NewDlyAggregatorSVC(
		logger log.Logger,
		stations stationssvc.Client,
		daily dailysvc.Client,
		alert alertsvc.Client,
		src *source.Hourly,
	) (*DailyAggregatorSVC, error) {
	options := prometheus.Opts{
		Name: "stations_update_count",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := DailyAggregatorSVC{
		stations: 	stations,
		daily: 		daily,
		alert:      alert,
		src:		src,
		logger: 	logger,
		counter: 	guage,
	}
	go startFetchProcess(&st)
	return &st,nil
}

func (das DailyAggregatorSVC) ForceUpdate(ctx context.Context, ids []string, start, end string) (err error) {
	das.processPeriodUpdate(ids, start, end)
	return err
}


func startFetchProcess(ss *DailyAggregatorSVC) {
	//ss.processUpdate() //Do it first time
	tick := time.Tick(time.Hour)
	for {
		select {
		case <-tick:
			ss.processUpdate()
		}
	}
}


func (das DailyAggregatorSVC)processUpdate() {
	sts, err := das.stations.GetAllStations()
	if err != nil {
		level.Error(das.logger).Log("msg", "GetStations error", "err", err)
		das.sendAlert(NewErrorAlert(err))
		return
	}

	ids := make([]string, len(sts.Sts))
	i := 0
	for k,_ := range sts.Sts {
		ids[i] = k
		i++
	}

	ch := make(chan *parser.StationDaily)
	go das.src.FetchLatestTemperature(ch, ids)

	count := 0.0
	for range ids {
		st := <-ch
		if st != nil {
			_, err := das.daily.PushPeriod(st.ID, st.Temps)
			if err != nil {
				level.Error(das.logger).Log("msg", "PushPeriod Error", "err", err)
				das.sendAlert(NewErrorAlert(err))
			} else {
				das.updateAverage(st.ID, st.Temps)
				count++
			}
		}
	}

	g := das.counter.With("stations", "all")
	g.Set(count)
	level.Info(das.logger).Log("msg", "Temperature updated", "stations", count)
}


func (das DailyAggregatorSVC)processPeriodUpdate(ids []string, start, end string) {
	ch := make(chan *parser.StationDaily)
	go das.src.FetchPeriodTemperature(ch, ids, start, end)

	count := 0.0
	for range ids {
		st := <-ch
		if st != nil {
			_, err := das.daily.PushPeriod(st.ID, st.Temps)
			if err != nil {
				level.Error(das.logger).Log("msg", "PushPeriod Error", "err", err)
				das.sendAlert(NewErrorAlert(err))
			} else {
				das.updateAverage(st.ID, st.Temps)
				count++
			}
		}
	}
	level.Info(das.logger).Log("msg", "Force temperature updated", "stations", count)
}

func (das DailyAggregatorSVC) updateAverage(id string, temps []dailysvc.Temperature) {
	for _,v := range temps {
		doy, err := getDOY(v.Date)
		if err != nil {
			level.Error(das.logger).Log("msg", "Average Update Error", "err", err)
			das.sendAlert(NewErrorAlert(err))
			continue
		}
		_, err = das.daily.UpdateAvgForDOY(id, doy)
		if err != nil {
			level.Error(das.logger).Log("msg", "Average Update Error", "err", err)
			das.sendAlert(NewErrorAlert(err))
		}
	}
}

func (das DailyAggregatorSVC)sendAlert(alert alertsvc.Alert) {
	err := das.alert.SendAlert(alert)
	if err != nil {
		level.Error(das.logger).Log("msg", "SendAlert Alert Error", "err", err)
		das.sendAlert(NewErrorAlert(err))
	}
}

func getDOY(date string) (int, error){
	t, err := time.Parse(common.TimeLayout, date)
	if err != nil {
		return 0, nil
	}

	return t.YearDay(),nil
}
