package source

import (
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/impl/parser"
	weathergrpc "github.com/flasherup/gradtage.de/weatherbitsvc/weatherbitgrpc"

	//"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"math"
	"time"
)
const veryFirstTime = "2000-01-02T01:01:01.00Z"

type Hourly struct {
	hrlService 	weatherbitsvc.Client
	dlyService 	dailysvc.Client
	logger 		log.Logger
}

func NewHourly(logger log.Logger, hourly *impl.WeatherBitSVCClient, daily dailysvc.Client) *Hourly {
	return &Hourly{
		hrlService:hourly,
		dlyService:daily,
		logger:logger,
	}
}

func (h Hourly) FetchLatestTemperature(ch chan *parser.StationDaily, ids []string) {
	hourlyUpdates, err := h.hrlService.GetUpdateDate(ids)
	if err != nil {
		level.Error(h.logger).Log("msg", "GetUpdateDate hourly error", "err", err)
		close(ch)
		return
	}
	dailyUpdates, err := h.dlyService.GetUpdateDate(ids)
	if err != nil {
		level.Error(h.logger).Log("msg", "GetUpdateDate daily error", "err", err)
		close(ch)
		return
	}


	currentTime := time.Now().Format(veryFirstTime)

	var (
		hrlUpdate string
		dlyUpdate string
	)
	for _,v := range ids {
		hrlUpdate = currentTime
		dlyUpdate = veryFirstTime

		if date, ok := hourlyUpdates.Dates[v]; ok {
			hrlUpdate = date
		} else {
			level.Warn(h.logger).Log("msg", "Daily update warning", "warn", "Station is not presented in hourly db", "station", v)
		}

		if date, ok := dailyUpdates.Dates[v]; ok {
			dlyUpdate = date
		} else {
			level.Warn(h.logger).Log("msg", "Daily update warning", "warn", "Station is not presented in daily db", "station", v)
		}

		h.fetchStation(v, ch, dlyUpdate, hrlUpdate)
	}
}

func (h Hourly) FetchPeriodTemperature(ch chan *parser.StationDaily, ids []string, start, end string) {
	for _,v := range ids {
		h.fetchStation(v, ch, start, end)
	}
}

func (h Hourly)fetchStation(id string, ch chan *parser.StationDaily, start string, end string) {
	level.Info(h.logger).Log("msg", "fetchStation", "id", id, "start", start, "end", end)
	period, err := h.hrlService.GetPeriod([]string{id}, start, end)
	if err != nil {
		level.Error(h.logger).Log("msg", "GetPeriod hourly error", "err", err, "station", id)
		ch <- nil
		return
	}

	dPeriod := h.hourlyToDaily(period.Temps[id].Temps, id)
	if len(dPeriod) == 0 {
		level.Warn(h.logger).Log("msg", "Daily update warning", "warn", "nothing to update", "station", id)
		ch <- nil
		return
	}

	ch <- &parser.StationDaily{ ID:id, Temps:dPeriod }
}

func (h Hourly)hourlyToDaily(src []*weathergrpc.Temperature, stId string) []dailysvc.Temperature {
	res := make([]dailysvc.Temperature , 0)
	day := 0
	var latest, current time.Time
	var err error
	dayCount := 0
	sum := 0.0
	for _,v := range src {
		current, err =  time.Parse(common.TimeLayout, v.Date)
		if err != nil {
			level.Error(h.logger).Log("msg", "Hourly to daily error", "err", err, "station", stId)
		}
		if day != current.Day() {
			if dayCount < 20 {
				if dayCount > 0 {
					es := fmt.Sprintf("Not enough data: %d hours, for: %s, station: %s", dayCount, latest.Format(common.TimeLayout), stId)
					level.Error(h.logger).Log("msg", "Hourly to daily error", "err", es)
				}
			} else {
				date := time.Date(latest.Year(), latest.Month(), latest.Day(), 0, 0, 0, 0, latest.Location() )
				ave := sum/float64(dayCount)
				res = append(res, dailysvc.Temperature{Date:date.Format(common.TimeLayout), Temperature:math.Round(ave*10)/10 })
			}
			dayCount = 0
			sum = 0.0
		}
		latest = current
		day = current.Day()
		dayCount++
		sum += v.Temperature
	}

	return res
}
