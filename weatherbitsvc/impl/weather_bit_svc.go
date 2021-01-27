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
	wb.processUpdate() //Do it first time
	tick := time.Tick(time.Hour)
	for {
		select {
		case <-tick:
			wb.processUpdate()
		}
	}
}

func (wb WeatherBitSVC)processUpdate() {
	url := wb.conf.Sources.UrlWeatherBit + "/current?lat=35.7796&lon=-78.6382&key=" + wb.conf.Sources.KeyWeatherBit + "&include=minutely"
	level.Info(wb.logger).Log("msg", "weather bit request", "url", url)
	//url := "https://api.checkwx.com/metar/" + id + "/decoded"
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		level.Error(wb.logger).Log("msg", "request error", "err", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		level.Error(wb.logger).Log("msg", "request error", "err", err)
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		level.Error(wb.logger).Log("msg", "response read error", "err", err)
		return
	}

	result, err := parser.ParseWeatherBit(&contents)
	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}