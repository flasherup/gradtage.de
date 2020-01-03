package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/impl/parser"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/impl/source"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type HourlyAggregatorSVC struct {
	stations    stationssvc.Client
	hourly 		hourlysvc.Client
	daily 		dailysvc.Client
	logger  	log.Logger
	counter 	*ktprom.Gauge
	src			source.Hourly
}

func NewHrlAggregatorSVC(logger log.Logger, stations stationssvc.Client, daily dailysvc.Client, src *source.Hourly) (*HourlyAggregatorSVC, error) {
	options := prometheus.Opts{
		Name: "stations_update_count",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := HourlyAggregatorSVC{
		stations: stations,
		daily: 	daily,
		logger: logger,
		counter: guage,
		src: *src,
	}
	go startFetchProcess(&st)
	return &st,nil
}

func (has HourlyAggregatorSVC) GetStatus(ctx context.Context) (temps []dlyaggregatorsvc.Status, err error) {
	level.Info(has.logger).Log("msg", "GetStatus", "ids")
	return temps,err
}


func startFetchProcess(ss *HourlyAggregatorSVC) {
	ss.processUpdate() //Do it first time
	tick := time.Tick(time.Hour)
	for {
		select {
		case <-tick:
			ss.processUpdate()
		}
	}
}


func (has HourlyAggregatorSVC)processUpdate() {
	sts := has.stations.GetAllStations()
	if sts.Err != "nil" {
		level.Error(has.logger).Log("msg", "GetStations error", "err", sts.Err)
		return
	}

	ids := make([]string, len(sts.Sts))
	i := 0
	for k,_ := range sts.Sts {
		ids[i] = k
		i++
	}

	ch := make(chan *parser.StationDaily)
	go has.src.FetchTemperature(ch, ids)

	count := 0.0
	for range ids {
		st := <-ch
		if st != nil {
			resp := has.daily.PushPeriod(st.ID, st.Temps)
			if resp.Err != "nil" {
				level.Error(has.logger).Log("msg", "PushPeriod Error", "err", resp.Err)
			} else {
				count++
			}
		}
	}

	g := has.counter.With("station")
	g.Set(count)
	level.Info(has.logger).Log("msg", "Temperature updated", "stations", count)
}
