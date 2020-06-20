package main

import (
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/alertsvc/impl"
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
			"svc", "stationssvcc",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	client := impl.NewAlertSCVClient("82.165.18.228:8107",logger)
	//client := impl.NewAlertSCVClient("localhost:8107",logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	/*err := client.SendAlert(alertsvc.Alert{
		Name:"Test",
		Desc:"Service test alert",
		Params:map[string]string{ "client": "alertClient" },
	})
	if err != nil {
		level.Error(logger).Log("msg", "SendAlert Alert error", "err", err)
	}*/


	err := client.SendEmail(alertsvc.Email{
		Name:"Test",
		Email:"flasherup@gmail.com",
		Params:map[string]string{ "key": "test key", "plan": "test plan" },
	})
	if err != nil {
		level.Error(logger).Log("msg", "Send Email error", "err", err)
	}
}