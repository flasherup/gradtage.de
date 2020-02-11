
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
	stations    stationssvc.Client
	hourly      hourlysvc.Client
	alert       alertsvc.Client
	logger      log.Logger
	counterCWX  *ktprom.Gauge
	counterDWD  *ktprom.Gauge
	counterNOAA *ktprom.Gauge
	checkWX     source.CheckWX
	dwd         source.SourceDWD
	noaa	    source.SourceNOAA
}


func NewHrlAggregatorSVC(
		logger log.Logger,
		stations stationssvc.Client,
		hourly hourlysvc.Client,
		alert alertsvc.Client,
		conf config.HrlAggregatorConfig,
	) (*HourlyAggregatorSVC, error) {

	optionsCWX := prometheus.Opts{
		Name: "stations_update_count_checkwx",
		Help: "The number of stations updated form CheckWX",
	}
	guageCWX := ktprom.NewGaugeFrom(prometheus.GaugeOpts(optionsCWX), []string{ common.SrcTypeCheckWX})

	optionsDWD := prometheus.Opts{
		Name: "stations_update_count_dwd",
		Help: "The number of stations updated form DWD",
	}
	guageDWD := ktprom.NewGaugeFrom(prometheus.GaugeOpts(optionsDWD), []string{ common.SrcTypeDWD})

	optionsNOAA := prometheus.Opts{
		Name: "stations_update_count_noaa",
		Help: "The number of stations updated form NOAA",
	}
	guageNOAA := ktprom.NewGaugeFrom(prometheus.GaugeOpts(optionsNOAA), []string{ common.SrcTypeNOAA})

	checkWX := source.NewCheckWX(conf.Sources.CheckwxKey, logger)
	dwd := source.NewDWD(conf.Sources.UrlDWD, logger)
	noaa := source.NewSourceNOAA(conf.Clients.NoaaAddr, logger)
	st := HourlyAggregatorSVC{
		stations:   	stations,
		hourly:     	hourly,
		alert:      	alert,
		logger:     	logger,
		counterCWX: 	guageCWX,
		counterDWD: 	guageDWD,
		counterNOAA: 	guageNOAA,
		checkWX:    	*checkWX,
		dwd:        	*dwd,
		noaa:			*noaa,
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
	ss.updateCheckWX()
	ss.updateDWD(-1)
	ss.updateNOAA(-1)

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
				ss.updateNOAA(24)
			}
		}
	}
}


func (has HourlyAggregatorSVC) updateCheckWX() {
	sts, err := has.stations.GetStationsBySrcType([]string{ common.SrcTypeCheckWX, common.SrcTypeNOAA })
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

	ch := make(chan *parser.ParsedData)
	go has.checkWX.FetchTemperature(ch, ids)

	count := 0.0
	for range ids {
		st := <-ch
		if st != nil && st.Success {
			has.verifyPlausibility(latest, st.StationID, st.Temps[0])
			_, err := has.hourly.PushPeriod(st.StationID, st.Temps)
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

	g := has.counterCWX.With(common.SrcTypeCheckWX)
	g.Set(count)
	level.Info(has.logger).Log("msg", "Temperature updated form CheckWX", "stations", count)
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

	g := has.counterDWD.With(common.SrcTypeDWD)
	g.Set(count)
	level.Info(has.logger).Log("msg", "Temperature updated from DWD", "stations", count)
}

func (has HourlyAggregatorSVC) updateNOAA(daysNumber int) {
	sts, err := has.stations.GetStationsBySrcType([]string{ common.SrcTypeNOAA })
	if err != nil {
		level.Error(has.logger).Log("msg", "Get NOAA Stations error", "err", err)
		has.sendAlert(NewErrorAlert(err))
		return
	}

	ids := make([]string, len(sts.Sts))
	i := 0
	for k,_ := range sts.Sts {
		ids[i] = k
		i++
	}

	ch := make(chan *parser.ParsedData)
	go has.noaa.FetchTemperature(ch, daysNumber, ids)

	count := 0.0
	for range ids {
		pd := <-ch
		if pd != nil && pd.Success {
			_, err := has.hourly.PushPeriod(pd.StationID, pd.Temps)
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

	g := has.counterNOAA.With(common.SrcTypeNOAA)
	g.Set(count)
	level.Info(has.logger).Log("msg", "Temperature updated from NOAA", "stations", count)
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


