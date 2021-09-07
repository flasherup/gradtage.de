package impl

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc/config"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl/parser"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"

	"net/http"
	"time"
)

type WeatherBitUpdateSVC struct {
	weatherbit    weatherbitsvc.Client
	alert 		alertsvc.Client
	logger  	log.Logger
	conf		config.WeatherBitConfig
	Context 	context.Context
	dailyRequestCounter   int
	secondsRequestCounter int
	dailyStartTime        time.Time
	secondsStartTime      time.Time
}



func NewWeatherBitUpdateSVC(
	logger 		log.Logger,
	weatherbit 	weatherbitsvc.Client,
	alert 		alertsvc.Client,
	conf 		config.WeatherBitConfig,
) (*WeatherBitUpdateSVC, error) {
	wbu := WeatherBitUpdateSVC {
		weatherbit:weatherbit,
		alert:alert,
		logger:logger,
		conf:conf,
		Context: context.Background(),
		dailyRequestCounter:   0,
		secondsRequestCounter: 0,
	}
	go wbu.runInfiniteUpdate(wbu.Context)
	return &wbu,nil
}

func (wbu *WeatherBitUpdateSVC) ForceRestart(ctx context.Context, ids []string, start string, end string)  error {
	level.Info(wbu.logger).Log("msg", "GetPeriod", "ids", fmt.Sprintf("Length:%d, Start:%s End:%s",len(ids), start, end))
	return nil
}

func (wbu *WeatherBitUpdateSVC) runInfiniteUpdate(ctx context.Context) error{
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
		date := time.Now()
		wbu.precessStations(date)
	}
}

/*func (wbu *WeatherBitUpdateSVC)precessStations() {
	sts, err := wbu.weatherbit.GetStationsList()

	if err != nil {
		level.Error(wbu.logger).Log("msg", "WeatherBitUpdate GetStationsList error", "err", err)
		return
	}

	for _ , station := range *sts {
		wb.processUpdate(station.Id, station.SourceId)
	}

}*/


func (wbu *WeatherBitUpdateSVC)precessStations(date time.Time) {
	wbu.dailyStartTime = date
	wbu.secondsStartTime = date
	sts, err := wbu.weatherbit.GetStationsList()
	if err != nil {
		level.Error(wbu.logger).Log("msg", "WeatherBitUpdate GetStationsList error", "err", err)
		return
	}
	for _,v := range *sts {
		level.Info(wbu.logger).Log("msg", "Process station", "innerId", k, "station", v)
		wbu.processRequest(v, date)
	}
}

func (wbu *WeatherBitUpdateSVC)processRequest(stID string, end time.Time) error {
	startDate := end
	requestsPerSecond := wbu.conf.WeatherBit.NumberOfRequestPerSecond
	requestsPerDay := wbu.conf.WeatherBit.NumberOfRequestPerDay
	for {
		start := end.AddDate(0, 0, -14)
		sDate := start.Format(common.TimeLayoutWBH)
		eDate := end.Format(common.TimeLayoutWBH)
		if yearCheck(startDate, end, wbh.conf.WeatherBit.NumberOfYears) {
			break
		}
		end = start
		if wbh.checkPeriod(stID, sDate, eDate) {
			continue
		}
		go wbh.processUpdate(stID, st, sDate, eDate)
		wbh.secondsRequestCounter, wbh.secondsStartTime = sleepCheck(requestsPerSecond, wbh.secondsRequestCounter, wbh.secondsStartTime, time.Second)
		wbh.dailyRequestCounter, wbh.dailyStartTime 	= sleepCheck(requestsPerDay, wbh.dailyRequestCounter, wbh.dailyStartTime, time.Hour * 24)
	}
	return nil
}

func (wbh WeatherHistorical)processUpdate(stID string, st string, start string, end string) error {
	err := wbh.db.CreateTable(stID)
	if err != nil {
		level.Error(wbh.logger).Log("msg", "table create error", "err", err)
		return err
	}

	url := wbh.conf.Sources.UrlWeatherBit + "/history/hourly?station=" + st + "&key=" + wbh.conf.Sources.KeyWeatherBit + "&start_date=" + start + "&end_date=" + end
	level.Info(wbh.logger).Log("msg", "weather bit request", "url", url)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		level.Error(wbh.logger).Log("msg", "request error", "err", err, "url", url)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		level.Error(wbh.logger).Log("msg", "request error", "err", err, "url", url)
		return err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		level.Error(wbh.logger).Log("msg", "response read error", "err", err, "url", url)
		return err
	}

	result, err := parser.ParseWeatherBit(&contents)
	if err != nil {
		level.Error(wbh.logger).Log("msg", "weather bit data parse error", "err", err, "url", url)
		return err
	}

	err = wbh.db.PushData(stID, result)
	if err != nil {
		level.Error(wbh.logger).Log("msg", "data push error", "err", err, "url", url)
		return err
	}
	return nil
}

func (wbh WeatherHistorical)checkPeriod(stID string, start string, end string) bool {
	temps, err := wbh.db.GetPeriod(stID, start, end)
	if err == nil {
		return len(temps) != 0
	}
	return false
}

func sleepCheck(numberOfRequests, counter int, startTime time.Time, duration time.Duration) (int, time.Time) {
	now := time.Now()
	dif := now.Sub(startTime)
	if dif >= duration {
		return 0, now
	}

	counter++
	if counter >= numberOfRequests{
		time.Sleep(duration - dif)
		return 0, time.Now()
	}
	return counter, startTime
}

func yearCheck(start, end time.Time, yearsCount int) bool {
	return 	start.Year() - end.Year() >= yearsCount &&
		start.Month() > end.Month()
}