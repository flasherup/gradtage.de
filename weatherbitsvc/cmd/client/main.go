package main

import (
	"fmt"
	"github.com/flasherup/gradtage.de/weatherbitsvc/impl"
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
			"svc", "weatherbitclient",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	client := impl.NewWeatherBitSVCClient("localhost:8111",logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	//Just for test
	data, err := client.GetPeriod([]string{"CYYC"}, "2021-03-20T01:00:00", "2021-03-25T23:00:00")
	if err != nil {
		level.Error(logger).Log("msg", "GetPeriod Error", "err", err)
		return
	}

	fmt.Println(*data)
}
