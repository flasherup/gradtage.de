package source

import (
	"fmt"
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/impl/parser"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/hourlysvc/hrlgrpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"time"
)

const veryFirstTime = "2000-01-02T01:00:00.000000000Z "

type Hourly struct {
	hrlService 	hourlysvc.Client
	dlyService 	dailysvc.Client
	logger 		log.Logger
}

func NewHourly(logger log.Logger, hourly hourlysvc.Client, daily dailysvc.Client) *Hourly {
	return &Hourly{
		hrlService:hourly,
		dlyService:daily,
		logger:logger,
	}
}

func (h Hourly) FetchTemperature(ch chan *parser.StationDaily, ids []string) {
	huorlyUpdates, err := h.hrlService.GetUpdateDate(ids)
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
		hrlUpdate = veryFirstTime

		if date, ok := huorlyUpdates.Dates[v]; ok {
			hrlUpdate = date
		}

		if date, ok := dailyUpdates.Dates[v]; ok {
			dlyUpdate = date
		}

		go h.fetchStation(v, ch, dlyUpdate, hrlUpdate)
	}
}

func (h Hourly)fetchStation(id string, ch chan *parser.StationDaily, start string, end string) {
	period, err := h.hrlService.GetPeriod(id, start, end)
	if err != nil {
		level.Error(h.logger).Log("msg", "GetPeriod hourly error", "err", err)
		ch <- nil
		return
	}

	dPeriod := h.hourlyToDaily(period.Temps)
	if len(dPeriod) == 0 {
		level.Warn(h.logger).Log("msg", "Daily update warning", "warn", "nothing to update", "station", id)
		ch <- nil
		return
	}

	ch <- &parser.StationDaily{ ID:id, Temps:dPeriod }
}


func (h Hourly)hourlyToDaily(src []*hrlgrpc.Temperature) []dailysvc.Temperature {
	res := make([]dailysvc.Temperature , 0)
	day := 0
	var latest, current time.Time
	var err error
	dayCount := 0
	sum := 0.0
	for _,v := range src {
		current, err =  time.Parse(veryFirstTime, v.Date)
		if err != nil {

		}
		if day != current.Day() {
			if dayCount < 20 {
				if dayCount > 0 {
					level.Error(h.logger).Log("msg", "Hourly to daily error", "err", fmt.Sprint("Not enough data: %d hours, for: %s", dayCount, latest.Format(veryFirstTime)))
				}
			} else {
				date := time.Date(latest.Year(), latest.Month(), latest.Day(), 0, 0, 0, 0, latest.Location() )
				ave := sum/float64(dayCount)
				res = append(res, dailysvc.Temperature{Date:date.Format(veryFirstTime), Temperature:ave})
			}
			dayCount = 0
			sum = 0.0
		}
		latest = current
		day = current.Day()
		dayCount++
		sum += float64(v.Temperature)
	}

	return res
}
