package source

import (
	"errors"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/impl/parser"
	"github.com/flasherup/gradtage.de/noaascrapersvc/noaascpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"math"
	"time"

	noaa "github.com/flasherup/gradtage.de/noaascrapersvc/impl"
)

type SourceNOAA struct {
	url 	string
	logger 	log.Logger
}

func NewSourceNOAA(url string, logger log.Logger) *SourceNOAA {
	return &SourceNOAA{
		url: 	url,
		logger:	logger,
	}
}

func (sn SourceNOAA) FetchTemperature(ch chan *parser.ParsedData, daysNumber int,  ids []string) {
	if daysNumber < 1 {
		daysNumber = 72
	}
	start,end := getDates(daysNumber)
	for _,v := range ids {
		go sn.fetchStation(v, start,end, ch)
	}
}


func (sn SourceNOAA)fetchStation(id string, startDate string, endDate string, ch chan *parser.ParsedData) {
	noaa := noaa.NewNoaaScraperSVCClient(sn.url, sn.logger)

	period, err := noaa.GetPeriod(id, startDate, endDate)

	if err != nil {
		level.Error(sn.logger).Log("msg", "Get period error", "start", startDate, "end", endDate, "err", err.Error())
		ch <- &parser.ParsedData{ Success:false, Error:err }
		return
	}

	if period.Err != "nil" {
		level.Error(sn.logger).Log("msg", "Get period error", "start", startDate, "end", endDate, "err", period.Err)
		ch <- &parser.ParsedData{ Success:false, Error: errors.New(period.Err) }
		return
	}

	res := parser.ParsedData{
		Success:   true,
		StationID: id,
		Temps:     processTemperature(period.Temps),
	}

	ch <- &res
}

func getDates(daysNumber int) (string, end string) {
	e := time.Now()
	s := e.AddDate(0,0, -daysNumber)

	return s.Format(common.TimeLayout), e.Format(common.TimeLayout)
}

func processTemperature(src []*noaascpc.Temperature ) []hourlysvc.Temperature {
	res := make([]hourlysvc.Temperature, 0, len(src))
	saved := map[string]int{}
	for _,v := range src {
		d, err := time.Parse(common.TimeLayout, v.Date)
		if err != nil {
			continue
		}
		dp := time.Date(d.Year(), d.Month(), d.Day(), d.Hour(), 0, 0, 0, d.Location())

		t := math.Floor(float64(v.Temperature) * 10) / 10
		temp := hourlysvc.Temperature {
			Date: 			dp.Format(common.TimeLayout),
			Temperature: 	t,
		}

		if index, ok := saved[temp.Date]; ok {
			res[index] = temp
		} else {
			res = append(res, temp)
			saved[temp.Date] = len(res)-1
		}
	}
	return res
}