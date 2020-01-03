
package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/impl/parser"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/impl/source"
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
	logger  	log.Logger
	counter 	*ktprom.Gauge
	src			source.CheckWX
}

func NewHrlAggregatorSVC(logger log.Logger, stations stationssvc.Client, hourly hourlysvc.Client, src *source.CheckWX) (*HourlyAggregatorSVC, error) {
	options := prometheus.Opts{
		Name: "stations_update_count",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := HourlyAggregatorSVC{
		stations: stations,
		hourly: hourly,
		logger: logger,
		counter: guage,
		src: *src,
	}
	go startFetchProcess(&st)
	return &st,nil
}

func (has HourlyAggregatorSVC) GetStatus(ctx context.Context) (temps []hrlaggregatorsvc.Status, err error) {
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

	ch := make(chan *parser.StationData)
	go has.src.FetchTemperature(ch, ids)

	count := 0.0
	for range ids {
		st := <-ch
		if st != nil {
			resp := has.hourly.PushPeriod(st.ICAO, []hourlysvc.Temperature{stationToTemperature(st)})
			if resp.Err != "nil" {
				level.Error(has.logger).Log("msg", "PushPeriod Error", "err", resp.Err)
			} else {
				count++
			}
		}
	}

	g := has.counter.With("stations")
	g.Set(count)
	level.Info(has.logger).Log("msg", "Temperature updated", "stations", count)
}

func stationToTemperature(st *parser.StationData) hourlysvc.Temperature {
	return  hourlysvc.Temperature{
		Date:st.Observed,
		Temperature:st.Temperature.Celsius,
	}
}
