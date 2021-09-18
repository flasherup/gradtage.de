package main

import (
	"fmt"
	//"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl"
	//"github.com/flasherup/gradtage.de/weatherbitsvc/impl/collectroes"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
	//"time"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "weatherbitclient",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	//client := impl.NewWeatherBitSVCClient("212.227.214.163:8111",logger)
	client := impl.NewWeatherBitSVCClient("localhost:8111",logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	/*err := getPeriod(client, logger)
	if err != nil {
		level.Error(logger).Log("msg", "GetPeriod Error", "err", err)
	}*/

	/*err := getWBPeriod(client, logger)
	if err != nil {
		level.Error(logger).Log("msg", "GetWBPeriod Error", "err", err)
	}*/

	/*err := pushWBPeriod(client, logger)
	if err != nil {
		level.Error(logger).Log("msg", "PushWBPeriod Error", "err", err)
	}*/

	err := getStationsList(client, logger)
	if err != nil {
		level.Error(logger).Log("msg", "getStationsList Error", "err", err)
	}


}

/*func getPeriod(client *impl.WeatherBitSVCClient, logger log.Logger) error {
	//Just for test
	data, err := client.GetPeriod([]string{"at_av222"}, "2020-03-20T00:00:00", "2021-03-25T20:00:00")
	if err != nil {
		return err
	}
	fmt.Println(*data)
	daysCollector :=  collectroes.NewDays()
	for _,v := range data.Temps {
		for _,t :=  range v.Temps {
			date, err := time.Parse(common.TimeLayout, t.Date)
			if err != nil {
				return err
			}
			daysCollector.Push(date.YearDay(), date.Hour(), t.Temperature)
		}
	}
	return nil
}*/

func getWBPeriod(client *impl.WeatherBitSVCClient, logger log.Logger) error {
	//Just for test
	temps, err := client.GetWBPeriod("it_liee", "2000-01-01T00:00:00", "2021-08-01T00:00:00")
	if err != nil {
		return err
	}
	for _,v := range *temps {
		fmt.Println(v)
	}

	return nil
}

func pushWBPeriod(client *impl.WeatherBitSVCClient, logger log.Logger) error {
	data := weatherbitsvc.WBData{
		Date: "2020-03-25T00:00:00Z" ,
		Rh: 67 ,
		Pod: "n" ,
		Pres: 952 ,
		Timezone: "America/New_York" ,
		CountryCode: "US" ,
		Clouds: 0 ,
		Vis: 16 ,
		SolarRad: 0 ,
		WindSpd: 6.8 ,
		StateCode: "PA" ,
		CityName: "Allens Mills" ,
		AppTemp: 2 ,
		Uv: 0 ,
		Lon: -78.9 ,
		Slp: 1018.3 ,
		HAngle: 0 ,
		Dewpt: 0.6 ,
		Snow: 0 ,
		Aqi: 0 ,
		WindDir: 110 ,
		ElevAngle: -5.96 ,
		Ghi: 0 ,
		Lat: 41.18 ,
		Precip: 0 ,
		Sunset: "" ,
		Temp: 6.1 ,
		Station: "" ,
		Dni: 0 ,
		Sunrise: "" ,
	}

	stID := "us_kduj";
	//Just for test
	err := client.PushWBPeriod(stID, []weatherbitsvc.WBData{data})
	if err != nil {
		return err
	}

	level.Info(logger).Log("msg", "Push WB data success")
	return nil
}

func getStationsList(client *impl.WeatherBitSVCClient, logger log.Logger) error {
	//Just for test
	stations, err := client.GetStationsList()
	if err != nil {
		return err
	}

	err = testMetrics(client, *stations)
	if err != nil {
		level.Error(logger).Log("msg", "getStationsList Error", "err", err)
	}
 return err
}

func testMetrics(client *impl.WeatherBitSVCClient,stations []string ) error {
	stName := stations
	//stName := []string{"us_k5sm","us_cwzr","us_cwzr","us_koak","us_k5sm","us_imbo","us_k1d7","us_k1cw","us_pfsh","us_kbgd","us_krdd","us_koxc","us_kirs"}
	data, err := client.GetStationsMetrics(stName)
	if err != nil {
		return err
	}
	for i,v := range *data{
		fmt.Println("stations", stations[i], "Lat", v.Lat, "Lon", v.Lon)
	}
	return nil
}
