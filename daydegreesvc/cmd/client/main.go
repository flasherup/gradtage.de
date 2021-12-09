package main

import (
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
	"github.com/flasherup/gradtage.de/daydegreesvc/impl"
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
			"svc", "daydegreeclient",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	client := impl.NewDayDegreeSVCClient("localhost:8112",logger)
	//client := impl.NewDayDegreeSVCClient("82.165.119.83:8112",logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	params := daydegreesvc.Params{
		Station:   "pl_epmi",
		Start:     "2010-01-01",
		End:       "2021-01-01",
		Breakdown: common.BreakdownDaily,
		Tb:        15,
		Tr:        20,
		Output:    "hdd",
		DayCalc:   common.DayCalcInt,

	}

	err := getDegree(client, params, logger)
	if err != nil {
		level.Error(logger).Log("msg", "GetDegree Error", "err", err)
	}

	err = getAVGDegree(client, params, logger)
	if err != nil {
		level.Error(logger).Log("msg", "getAverageDegree Error", "err", err)
	}
}

func getDegree(client *impl.DayDegreeSVCClient, params daydegreesvc.Params, logger log.Logger) error {
	//Just for test
	data, err := client.GetDegree(params)
	if err != nil {
		return err
	}
	for i,v := range data {
		t, err := time.Parse(common.TimeLayoutDay, v.Date)
		if err == nil && t.Month() == 4 && t.Day() == 1{
			fmt.Println(i, v.Date, v.Temp)
		}
	}
	return nil
}

func getAVGDegree(client *impl.DayDegreeSVCClient, params daydegreesvc.Params, logger log.Logger) error {
	years := 10
	//Just for test
	data, err := client.GetAverageDegree(params, years)
	if err != nil {
		return err
	}
	for i,v := range data {
		t, err := time.Parse(common.TimeLayoutDay, v.Date)
		if err == nil && t.Month() == 4 && t.Day() == 1{
			fmt.Println(i, v.Date, v.Temp)
		}
	}
	return nil
}
