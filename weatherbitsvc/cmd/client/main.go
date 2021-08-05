package main

import (
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl/collectroes"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
	"time"
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

	err := getWBPeriod(client, logger)
	if err != nil {
		level.Error(logger).Log("msg", "GetWBPeriod Error", "err", err)
	}
}

func getPeriod(client *impl.WeatherBitSVCClient, logger log.Logger) error {
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
}

func getWBPeriod(client *impl.WeatherBitSVCClient, logger log.Logger) error {
	//Just for test
	temps, err := client.GetWBPeriod("us_koak", "2021-03-25T00:00:00", "2021-03-25T20:00:00")
	if err != nil {
		return err
	}
	for _,v := range *temps {
		fmt.Println(v)
	}

	return nil
}
