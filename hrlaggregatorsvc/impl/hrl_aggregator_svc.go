
package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/hourlysvc/hrlgrpc"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/config"
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
	stations stationssvc.Client
	hourly   hourlysvc.Client
	alert    alertsvc.Client
	logger   log.Logger
	counter  *ktprom.Gauge
	checkWX  source.CheckWX
	dwd 	 source.SourceDWD
}

func NewHrlAggregatorSVC(
		logger log.Logger,
		stations stationssvc.Client,
		hourly hourlysvc.Client,
		alert alertsvc.Client,
		conf config.HrlAggregatorConfig,
	) (*HourlyAggregatorSVC, error) {
	options := prometheus.Opts{
		Name: "stations_update_count",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	checkWX := source.NewCheckWX(conf.Sources.CheckwxKey, logger)
	dwd := source.NewDWD(conf.Sources.UrlDWD, logger)
	st := HourlyAggregatorSVC{
		stations: stations,
		hourly:   hourly,
		alert:    alert,
		logger:   logger,
		counter:  guage,
		checkWX:  *checkWX,
		dwd: 	  *dwd,
	}
	go startFetchProcess(&st)
	return &st,nil
}

func (has HourlyAggregatorSVC) GetStatus(ctx context.Context) (temps []hrlaggregatorsvc.Status, err error) {
	level.Info(has.logger).Log("msg", "GetStatus", "ids")
	return temps,err
}

const (
	updateCheckWX 	= 1
	updateDWD 		= 2
)


func startFetchProcess(ss *HourlyAggregatorSVC) {
	ss.updateCheckWX() //Do it first time
	ss.updateDWD(-1) //Do it first time


	chTimer := make(chan bool)
	chAlarm := make(chan bool)

	go runTimer(chTimer, time.Hour)
	go runAlarm(chAlarm, 23, 59)

	for {
		select {
		case timer := <- chTimer:
			if timer {
				ss.updateCheckWX()
			}

		case alarm := <- chAlarm:
			if alarm {
				ss.updateDWD(24)
			}
		}
	}
}


func (has HourlyAggregatorSVC) updateCheckWX() {
	sts, err := has.stations.GetStationsBySrcType([]string{ common.SrcTypeCheckWX })
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

	ch := make(chan *parser.StationDataCheckWX)
	go has.checkWX.FetchTemperature(ch, ids)

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

func (has HourlyAggregatorSVC) updateDWD(rowsNumber int) {
	sts, err := has.stations.GetStationsBySrcType([]string{ common.SrcTypeDWD })
	if err != nil {
		level.Error(has.logger).Log("msg", "Get DWD Stations error", "err", err)
		has.sendAlert(NewErrorAlert(err))
		return
	}

	ids := make(map[string]string)
	for k,v := range sts.Sts {
		ids[k] = v.SourceId
	}

	/*latest, err := has.hourly.GetLatest(ids)
	if err != nil {
		level.Error(has.logger).Log("msg", "Get latest error", "err", err)
		has.sendAlert(NewErrorAlert(err))
	}*/

	ch := make(chan *parser.ParsedData)
	go has.dwd.FetchTemperature(ch, ids)

	count := 0.0
	for range ids {
		pd := <-ch
		if pd != nil && pd.Success {
			//has.verifyPlausibility(latest, pd.StationID, pd.Temps)
			rowsToUpdate := pd.Temps
			if rowsNumber > 0 {
				rowsToUpdate = rowsToUpdate[len(rowsToUpdate)-rowsNumber:]
			}
			_, err := has.hourly.PushPeriod(pd.StationID, rowsToUpdate)
			if err != nil {
				level.Error(has.logger).Log("msg", "PushPeriod Error", "err", err)
				has.sendAlert(NewErrorAlert(err))
			} else {
				count++
			}
		} else {
			if pd != nil {
				level.Error(has.logger).Log("msg", "Station update error", "err", pd.Error)
			}
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
	if dif > 10 {
		has.sendAlert(NewTemperatureChangeAlert(*prevTemp, currentTemp, currentId))
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
		has.sendAlert(NewTemperatureGapAlert(*prevTemp, currentTemp, currentId))
	}

	if difT < time.Hour {
		//has.sendAlert(NewTemperatureSameDateAlert(*prevTemp, currentTemp))
		level.Warn(has.logger).Log("msg", "Hourly temperature update period less the hour", "station", currentId)
	}
}

func (has HourlyAggregatorSVC)sendAlert(alert alertsvc.Alert) {
	err := has.alert.SendAlert(alert)
	if err != nil {
		level.Error(has.logger).Log("msg", "Send Alert Error", "err", err)
	}
}

func stationToTemperature(st *parser.StationDataCheckWX) hourlysvc.Temperature {
	return  hourlysvc.Temperature{
		Date:st.Observed,
		Temperature:st.Temperature.Celsius,
	}
}


//Timer

func runTimer(ch chan bool, period time.Duration) {
	defer close(ch)
	for {
		time.Sleep(period)
		ch <- true
	}
}


func runAlarm(ch chan bool, hours int, minutes int) {
	defer close(ch)
	loc, err := time.LoadLocation("CET")
	if err != nil {
		ch <- false
		return
	}

	for {
		current := time.Now()
		current = current.In(loc)
		period := getAlarmDuration(hours, minutes, current)
		time.Sleep(period)
		ch <- true
	}
}


func getAlarmDuration(hours int, minutes int, current time.Time) time.Duration {
	alarm := time.Date(current.Year(), current.Month(), current.Day(), hours, minutes, 0, 0, current.Location())
	if alarm.Before(current) || alarm.Equal(current){
		alarm = alarm.Add(time.Hour * 24)
	}
	return alarm.Sub(current)
}


