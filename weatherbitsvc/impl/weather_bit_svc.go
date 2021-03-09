package impl

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/hourlysvc"
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
	db 			database.WeatherBitDB
	alert 		alertsvc.Client
	logger  	log.Logger
	conf		config.WeatherBitConfig
}



func NewWeatherBitSVC(
	logger 		log.Logger,
	db 			database.WeatherBitDB,
	alert 		alertsvc.Client,
	conf 		config.WeatherBitConfig,
) (*WeatherBitSVC, error) {
	wb := WeatherBitSVC {
		db:db,
		alert:alert,
		logger:logger,
		conf:conf,
	}
	go startFetchProcess(&wb)
	return &wb,nil
}

func (wb WeatherBitSVC) GetPeriod(ctx context.Context, ids []string, start string, end string) (temps map[string][]hourlysvc.Temperature, err error) {
	level.Info(wb.logger).Log("msg", "GetPeriod", "ids", fmt.Sprintf("Length:%d, Start:%s End:%s",len(ids), start, end))
	temps = make(map[string][]hourlysvc.Temperature)
	return temps,err
}

func startFetchProcess(wb *WeatherBitSVC) {
	wb.precessStations() //Do it first time
	tick := time.Tick(time.Hour)
	for {
		select {
		case <-tick:
			wb.precessStations()
		}
	}
}

	func (wb WeatherBitSVC)precessStations() {
	stations := map[string]string{
		"KNYC": "KNYC",
		"WMO7650": "LFML",
		"KATT": "KATT",
		"EDDH": "EDDH",
		"CYYC": "CYYC",
		"WMO10224": "10224",
		"LEBL": "LEBL",
		"WMO8181": "081810",
		"ESMS":"ESMS",
		"LFBN":"LFBN",
		"D4932":"D4932",
		"W07301399999":"073013-99999",
	}
	for k,v := range stations {
		wb.processUpdate(k, v)
	}
}

func (wb WeatherBitSVC)processUpdate(stID string, st string) {
	url := wb.conf.Sources.UrlWeatherBit + "/current?station=" + st + "&key=" + wb.conf.Sources.KeyWeatherBit
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		level.Error(wb.logger).Log("msg", "request error", "err", err, "id", stID, "station", st)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		level.Error(wb.logger).Log("msg", "request error", "err", err, "id", stID, "station", st)
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		level.Error(wb.logger).Log("msg", "response read error", "err", err, "id", stID, "station", st)
		return
	}

	result, err := parser.ParseWeatherBit(&contents)
	if err != nil {
		level.Error(wb.logger).Log("msg", "weather bit data parse error", "err", err, "id", stID, "station", st)
		return
	}

	err = wb.db.CreateTable(stID)
	if err != nil {
		level.Error(wb.logger).Log("msg", "table create error", "err", err, "id", stID, "station", st)
		return
	}

	err = wb.db.PushData(stID, result)
	if err != nil {
		level.Error(wb.logger).Log("msg", "data push error", "err", err, "id", stID, "station", st)
		return
	}
}