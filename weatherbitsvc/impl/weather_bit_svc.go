package impl

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc/config"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl/database"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"math"
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

func (wb WeatherBitSVC) GetStationsMetrics(ctx context.Context, ids []string, cutDate string) (map[string]weatherbitsvc.StationMetrics, error) {
	stsLen := len(ids)
	var err error
	if stsLen == 0 {
		ids, err = wb.db.GetListOfTables()
		if err != nil {
			return nil, err
		}
	}

	stsLen = len(ids)
	level.Info(wb.logger).Log("msg", "GetStationsMetrics" , "ids length ", stsLen)
	res := make(map[string]weatherbitsvc.StationMetrics)

	if stsLen == 0 {
		return res, nil
	}


	queryLen := 1000
	iterations := int(math.Floor(float64(stsLen/queryLen))) + 1
	for i:=0; i<iterations; i++ {
		s := i * queryLen
		e := s + queryLen
		if e > stsLen {
			e = s + stsLen%queryLen
		}
		fmt.Println("s", s, "e", e)
		currentIds := ids[s:e]
		wbData, err := wb.db.GetLastRecords(currentIds)
		if err != nil {
			return nil, err
		}

		recordsNumber, RNErr := wb.db.GetRecordsNumber(currentIds, cutDate)
		if RNErr != nil {
			return nil, err
		}

		for k,v := range wbData {
			res[k] = weatherbitsvc.StationMetrics{
				StId: k,
				LastUpdate: v.Date,
				Lat: v.Lat,
				Lon: v.Lon,
				RecordsNumber: recordsNumber[k],
			}
		}
	}


	return res, nil
}
