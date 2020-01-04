package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/impl/parser"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/impl/source"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type HourlyAggregatorSVC struct {
	stations    stationssvc.Client
	src 		*source.Hourly
	daily 		dailysvc.Client
	logger  	log.Logger
	counter 	*ktprom.Gauge
}

func NewDlyAggregatorSVC(logger log.Logger, stations stationssvc.Client, daily dailysvc.Client, src *source.Hourly) (*HourlyAggregatorSVC, error) {
	options := prometheus.Opts{
		Name: "stations_update_count",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := HourlyAggregatorSVC{
		stations: 	stations,
		daily: 		daily,
		src:		src,
		logger: 	logger,
		counter: 	guage,
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
	sts, err := has.stations.GetAllStations()
	if err != nil {
		level.Error(has.logger).Log("msg", "GetStations error", "err", err)
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
			_, err := has.daily.PushPeriod(st.ID, st.Temps)
			if err != nil {
				level.Error(has.logger).Log("msg", "PushPeriod Error", "err", err)
			} else {
				has.updateAverage(st.ID, st.Temps)
				count++
			}
		}
	}

	g := has.counter.With("stations", "all")
	g.Set(count)
	level.Info(has.logger).Log("msg", "Temperature updated", "stations", count)
}

func (has HourlyAggregatorSVC) updateAverage(id string, temps []dailysvc.Temperature) {
	for _,v := range temps {
		doy, err := getDOY(v.Date)
		if err != nil {
			level.Error(has.logger).Log("msg", "Average Update Error", "err", err)
			continue
		}
		_, err = has.daily.UpdateAvgForDOY(id, doy)
		if err != nil {
			level.Error(has.logger).Log("msg", "Average Update Error", "err", err)
		}
	}
}

func getDOY(date string) (int, error){
	t, err := time.Parse(common.TimeLayout, date)
	if err != nil {
		return 0, nil
	}

	return t.YearDay(),nil
}
