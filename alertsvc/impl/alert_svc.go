package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

type AlertSVC struct {
	logger  	log.Logger
	counter 	*ktprom.Gauge
}

func NewAlertSVC(logger log.Logger) (*AlertSVC, error) {
	options := prometheus.Opts{
		Name: "alerts_count_total",
		Help: "The total number of alerts",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := AlertSVC{
		logger: logger,
		counter: guage,
	}
	return &st,nil
}

func (a AlertSVC) SendAlert(ctx context.Context, alert alertsvc.Alert) error {
	level.Info(a.logger).Log("msg", "Send Alert", "Name", alert.Name, "desc", alert.Desc)
	return nil
}

