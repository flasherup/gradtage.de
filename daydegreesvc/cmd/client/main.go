package main

import (
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
	"github.com/flasherup/gradtage.de/daydegreesvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "daydegreeclient",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	//client := impl.NewDayDegreeSVCClient("localhost:8112",logger)
	client := impl.NewDayDegreeSVCClient("82.165.119.83:8112",logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	err := getDegree(client, logger)
	if err != nil {
		level.Error(logger).Log("msg", "GetDegree Error", "err", err)
	}
}

func getDegree(client *impl.DayDegreeSVCClient, logger log.Logger) error {
	params := daydegreesvc.Params{
		Station: "us_koak",
		Start: "2020-01-01",
		End: "2020-12-31",
		Breakdown: common.BreakdownDaily,
		Tb: 15,
		Tr: 20,
		Method: "hdd",
		DayCalc: common.DayCalcMean,

	}
	//Just for test
	data, err := client.GetDegree(params)
	if err != nil {
		return err
	}
	for i,v := range data {
		fmt.Println(i, v.Date, v.Temp)
	}
	return nil
}
