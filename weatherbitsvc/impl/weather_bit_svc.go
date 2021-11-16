package impl

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc/config"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl/database"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl/parser"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"net/http"
	"time"
)

type WeatherBitSVC struct {
	stations stationssvc.Client
	db       database.WeatherBitDB
	alert    alertsvc.Client
	logger   log.Logger
	conf     config.WeatherBitConfig
}

func NewWeatherBitSVC(
	logger log.Logger,
	stations stationssvc.Client,
	db database.WeatherBitDB,
	alert alertsvc.Client,
	conf config.WeatherBitConfig,
) (*WeatherBitSVC, error) {
	wb := WeatherBitSVC{
		stations: stations,
		db:       db,
		alert:    alert,
		logger:   logger,
		conf:     conf,
	}

	//processUpdate(wb, startDate, endDate)

	//go startFetchProcess(&wb)
	return &wb, nil
}

func (wb WeatherBitSVC) GetPeriod(ctx context.Context, ids []string, start string, end string) (temps map[string][]common.Temperature, err error) {
	level.Info(wb.logger).Log("msg", "GetPeriod", "ids", fmt.Sprintf("Length:%d, Start:%s End:%s", len(ids), start, end))
	temps = make(map[string][]common.Temperature)

	for _, id := range ids {
		t, err := wb.db.GetPeriod(id, start, end)
		if err != nil {
			return temps, err
		}
		temps[id] = t
	}
	return temps, err
}

func (wb WeatherBitSVC) GetWBPeriod(ctx context.Context, id string, start string, end string) (temps []weatherbitsvc.WBData, err error) {
	level.Info(wb.logger).Log("msg", "GetWBPeriod", "id", id, "start", start, "end", end)

	temps, err = wb.db.GetWBData(id, start, end)
	if err != nil {
		return temps, err
	}
	return
}

func (wb WeatherBitSVC) PushWBPeriod(ctx context.Context, id string, data []weatherbitsvc.WBData) (err error) {
	level.Info(wb.logger).Log("msg", "PushWBPeriod", "id", id, "length", len(data))

	err = wb.db.CreateTable(id)
	if err != nil {
		level.Error(wb.logger).Log("msg", "PushWBPeriod table create error", "err", err)
	}

	err = wb.db.PushWBData(id, data)
	if err != nil {
		return err
	}
	return
}

func startFetchProcess(wb *WeatherBitSVC) {
	wb.precessStations() //Do it first time
	tick := time.Tick(time.Hour * 24)
	for {
		select {
		case <-tick:
			wb.precessStations()
		}
	}
}

func (wb WeatherBitSVC) precessStations() {
	sts, err := wb.stations.GetAllStations()

	if err != nil {
		level.Error(wb.logger).Log("msg", "WeatherBit GetStations error", "err", err)
		return
	}

	for _, station := range sts.Sts {
		wb.processUpdate(station.Id, station.SourceId)
	}

}

func (wb *WeatherBitSVC) GetUpdateDate(ctx context.Context, ids []string) (dates map[string]string, err error) {
	level.Info(wb.logger).Log("msg", "GetUpdateDate", "ids", fmt.Sprintf("%+q:", ids))
	dates, err = wb.db.GetUpdateDateList(ids)
	if err != nil {
		level.Error(wb.logger).Log("msg", "Get Update Date List error", "err", err)
		//hs.sendAlert(NewErrorAlert(err))
	}
	return dates, err
}

func (wb *WeatherBitSVC) GetStationsList(ctx context.Context) (stations []string, err error) {
	level.Info(wb.logger).Log("msg", "GetStationsList")
	stations, err = wb.db.GetListOfTables()
	if err != nil {
		level.Error(wb.logger).Log("msg", "Get list fo stations error", "err", err)
	}
	return stations, err
}

func (wb *WeatherBitSVC) GetAverage(ctx context.Context, id string, years int, end string) ([]common.Temperature, error) {
	level.Info(wb.logger).Log("msg", "GetAverage", "id", id, "years", years, "end", end)


	if years < 1 {
		years = 1
	}
	start, err := getStartDate(end, years)
	if err == nil{

	}

	data, err := wb.db.GetPeriod(id, start, end)
	if err != nil {
		level.Error(wb.logger).Log("msg", "Get WB Data error", "err", err)
		return []common.Temperature{}, err
	}

	temps := make([]common.Temperature, len(data)/years)
	//allData := make(map[int][]common.Temperature)
	forGAFloat :=  make([]float64, 25)
	for _, v := range data {
		d, err := time.Parse(common.TimeLayout, v.Date)
		if err == nil{
			level.Error(wb.logger).Log("msg", "Time Parse error", "err", err)
		}
		//key := fmt.Sprintf("%d - %d - %d", d.Month(), d.Day(), d.Hour())

		//getPeriodData := v.Date
		//getPeriodTemps := v.Temp
		//allData[key] = data
	
		//_, exist := allData[key]
		t, err := time.Parse(common.TimeLayout, v.Date)
		if err == nil{
			level.Error(wb.logger).Log("msg", "Time Parse error", "err", err)
		}
		fmt.Println(t, d)

		//if !exist {}//va := make([]common.Temperature, len(data))}
		for i, v := range data{
			if t.Month() == d.Month() && t.Day() == d.Day() && d.Hour() == 1{
				//temp := v.Temp
				//allData[d.Hour()] = v
				fmt.Println(v.Date, v.Temp)
				fmt.Println(t.Hour())
			}
			if i % 24 == 0 {
				toAverage :=common.GetAverageFloat64(forGAFloat)
				result := common.ToFixedFloat64(toAverage, 2)
				fmt.Println(result)
			}


		}


		//fmt.Println(getPeriodTemps)
		//toAverage :=common.GetAverageFloat64(forGAFloat[i])


		//fmt.Println(toAverage)


	//

		//temps[i] = common.Temperature{key, result}




//если длинна даты < 8734 вернуть за 1 год
	}

	return temps, err
}

func (wb *WeatherBitSVC) GetAverageDegree(ctx context.Context, params weatherbitsvc.Params, years int)  ([]weatherbitsvc.Degree, error) {
	id := "us_koak"
		level.Info(wb.logger).Log("msg", "GetDegree", "Station", params.Station, "Start", params.Start, "End", params.End)
		_, err := wb.db.GetPeriod(id, params.Start, params.End)
		if err != nil {
			level.Error(wb.logger).Log("msg", "GetPeriod error", "err", err)
			return []weatherbitsvc.Degree{}, err
		}
	/*
		var degrees *[]common.Temperature
		t := (*temps)[params.Station]
		if params.Output == common.HDDType {
			degrees = common.CalculateHDDDegree(t, params.Tb, params.Breakdown, params.DayCalc)
		} else if params.Output == common.DDType {
			degrees = common.CalculateDDegree(t, params.Tb, params.Tr, params.Breakdown, params.DayCalc)
		} else if params.Output == common.CDDType {
			degrees = common.CalculateCDDegree(t, params.Tb, params.Breakdown, params.DayCalc)
		}

		res := toDegree(degrees)
		return *res, nil
	}

	func toDegree(temps *[]common.Temperature) *[]daydegreesvc.Degree {
		if temps == nil {
			return &[]daydegreesvc.Degree{}
		}
		res := make([]daydegreesvc.Degree, len(*temps))
		for i,v := range *temps {
			res[i] =  daydegreesvc.Degree{
				Date: v.Date,
				Temp: v.Temp,
			}
		}
		return &res*/
	return []weatherbitsvc.Degree{}, err
}

func (wb WeatherBitSVC) processUpdate(stID string, st string) {
	date := time.Now()
	endDate := date.Format(common.TimeLayoutWBH)
	sDate := date.AddDate(0, 0, -2)
	startDate := sDate.Format(common.TimeLayoutWBH)

	url := wb.conf.Sources.UrlWeatherBit + "/history/hourly?station=" + st + "&key=" + wb.conf.Sources.KeyWeatherBit + "&start_date=" + startDate + "&end_date=" + endDate
	//url := wb.conf.Sources.UrlWeatherBit + "/current?station=" + st + "&key=" + wb.conf.Sources.KeyWeatherBit
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		level.Error(wb.logger).Log("msg", "request error", "err", err, "id", stID, "station", st, "url", url)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		level.Error(wb.logger).Log("msg", "request error", "err", err, "id", stID, "station", st, "url", url)
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		level.Error(wb.logger).Log("msg", "response read error", "err", err, "id", stID, "station", st, "url", url)
		return
	}

	result, err := parser.ParseWeatherBit(&contents)
	if err != nil {
		level.Error(wb.logger).Log("msg", "weather bit data parse error", "err", err, "id", stID, "station", st, "url", url)
		return
	}

	err = wb.db.CreateTable(stID)
	if err != nil {
		level.Error(wb.logger).Log("msg", "table create error", "err", err, "id", stID, "station", st, "url", url)
		return
	}

	err = wb.db.PushData(stID, result)
	if err != nil {
		level.Error(wb.logger).Log("msg", "data push error", "err", err, "id", stID, "station", st, "url", url)
		return
	}

}

func getStartDate(end string, years int) (string, error) {
	t, err := time.Parse(common.TimeLayoutWBH, end)
	if err != nil{
		return end, err
	}
	layout := common.TimeLayoutWBH
	time.Parse(layout, end)
	start := t.AddDate(-years, 0, 0)
	return start.Format(common.TimeLayoutWBH), nil
}

func toAverageDegree(temps *[]common.Temperature) *[]daydegreesvc.Degree {
	if temps == nil {
		return &[]daydegreesvc.Degree{}
	}
	res := make([]daydegreesvc.Degree, len(*temps))
	for i,v := range *temps {
		res[i] =  daydegreesvc.Degree{
			Date: v.Date,
			Temp: v.Temp,
		}
	}
	return &res
}
