package common

import (
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

type SilentAlert struct {
	logger log.Logger
}

//NewSilentAlert return initialised SilentAlert
func NewSilentAlert() *SilentAlert {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "silentalert",
			"ts", log.DefaultTimestampUTC,
			"caller", log.Caller(3),
		)
	}
	return &SilentAlert{ logger }
}

//SendAlert log @alert
func (sa SilentAlert) SendAlert(alert alertsvc.Alert) error {
	level.Debug(sa.logger).Log("msg", "Silent Alert",
		"Name", alert.Name, "Desc", alert.Desc);
	return nil
}
