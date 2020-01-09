package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/alertsvc/impl/alertsys"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

type AlertSVC struct {
	logger  	log.Logger
	alertSys    alertsys.AlertSystem
	counter 	*ktprom.Gauge
}

func NewAlertSVC(logger log.Logger, alertSystem alertsys.AlertSystem) (*AlertSVC, error) {
	options := prometheus.Opts{
		Name: "alerts_count_total",
		Help: "The total number of alerts",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "name" })
	st := AlertSVC{
		logger: logger,
		alertSys: alertSystem,
		counter: guage,
	}
	return &st,nil
}

func (a AlertSVC) SendAlert(ctx context.Context, alert alertsvc.Alert) error {
	level.Info(a.logger).Log("msg", "Send Alert", "Name", alert.Name, "desc", alert.Desc)
	err := a.alertSys.Send(alert)
	if err != nil {
		level.Error(a.logger).Log("msg", "Send Alert error", "err", err)

	}

	g := a.counter.With("name", alert.Name)
	g.Add(1.0)
	return err
}

