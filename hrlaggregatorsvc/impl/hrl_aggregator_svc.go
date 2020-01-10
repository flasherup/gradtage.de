
package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/hourlysvc/hrlgrpc"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/impl/parser"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/impl/source"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"math"
	"time"
)

type HourlyAggregatorSVC struct {
	stations    stationssvc.Client
	hourly 		hourlysvc.Client
	alert 		alertsvc.Client
	logger  	log.Logger
	counter 	*ktprom.Gauge
	src			source.CheckWX
}

func NewHrlAggregatorSVC(
		logger log.Logger,
		stations stationssvc.Client,
		hourly hourlysvc.Client,
		alert alertsvc.Client,
		src *source.CheckWX,
	) (*HourlyAggregatorSVC, error) {
	options := prometheus.Opts{
		Name: "stations_update_count",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := HourlyAggregatorSVC{
		stations: stations,
		hourly: hourly,
		alert: alert,
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
	sts, err := has.stations.GetAllStations()
	if err != nil {
		level.Error(has.logger).Log("msg", "GetStations error", "err", err)
		has.sendAlert(NewErrorAlert(err))
		return
	}

	ids := make([]string, len(sts.Sts))
	i := 0
	for k,_ := range sts.Sts {
		ids[i] = k
		i++
	}

	latest, err := has.hourly.GetLatest(ids)
	if err != nil {
		level.Error(has.logger).Log("msg", "Get latest error", "err", err)
		has.sendAlert(NewErrorAlert(err))
	}

	ch := make(chan *parser.StationData)
	go has.src.FetchTemperature(ch, ids)

	count := 0.0
	for range ids {
		st := <-ch
		if st != nil {
			temp := stationToTemperature(st)
			has.verifyPlausibility(latest, st.ICAO, temp)
			_, err := has.hourly.PushPeriod(st.ICAO, []hourlysvc.Temperature{temp})
			if err != nil {
				level.Error(has.logger).Log("msg", "PushPeriod Error", "err", err)
				has.sendAlert(NewErrorAlert(err))
			} else {
				count++
			}
		} else {
			level.Warn(has.logger).Log("msg", "Station is not updated")
		}
	}

	g := has.counter.With("stations")
	g.Set(count)
	level.Info(has.logger).Log("msg", "Temperature updated", "stations", count)
}

func (has HourlyAggregatorSVC)verifyPlausibility(latest *hrlgrpc.GetLatestResponse, currentId string, currentTemp hourlysvc.Temperature) {
	if latest == nil {
		level.Warn(has.logger).Log("msg", "Can't check Plausibility", "reason", "no latest temperature data")
		return
	}

	var prevTemp *hrlgrpc.Temperature
	var ok bool
	if prevTemp, ok = latest.Temps[currentId]; !ok {
		level.Warn(has.logger).Log("msg", "Can't check Plausibility", "reason", "station with ID:" + currentId + " is not exist")
		return
	}

	dif := math.Abs(float64(prevTemp.Temperature) - currentTemp.Temperature)
	if dif > 5 {
		has.sendAlert(NewTemperatureChangeAlert(*prevTemp, currentTemp))
	}

	prevD, err := time.Parse(common.TimeLayout, prevTemp.Date)
	if err != nil {
		level.Warn(has.logger).Log("msg", "Plausibility check error", "err", err)
		return
	}

	curD, err := time.Parse(common.TimeLayout, currentTemp.Date)
	if err != nil {
		level.Warn(has.logger).Log("msg", "Plausibility check error", "err", err)
		return
	}

	difT := curD.Sub(prevD)
	if difT > time.Hour * 3 {
		has.sendAlert(NewTemperatureGapAlert(*prevTemp, currentTemp))
	}

	if difT < time.Hour {
		//has.sendAlert(NewTemperatureSameDateAlert(*prevTemp, currentTemp))
		level.Warn(has.logger).Log("msg", "Hourly temperature update period less the hour", "station", currentId)
	}
}

func (has HourlyAggregatorSVC)sendAlert(alert alertsvc.Alert) {
	_, err := has.alert.SendAlert(alert)
	if err != nil {
		level.Error(has.logger).Log("msg", "Send Alert Error", "err", err)
	}
}

func stationToTemperature(st *parser.StationData) hourlysvc.Temperature {
	return  hourlysvc.Temperature{
		Date:st.Observed,
		Temperature:st.Temperature.Celsius,
	}
}


