
package impl

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/hourlysvc/impl/database"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

type HourlySVC struct {
	logger  	log.Logger
	alert 		alertsvc.Client
	db 			database.HourlyDB
	counter 	*ktprom.Gauge
}

func NewHourlySVC(logger log.Logger, db database.HourlyDB, alert alertsvc.Client) (*HourlySVC, error) {
	options := prometheus.Opts{
		Name: "stations_count_total",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := HourlySVC{
		logger:  logger,
		alert:   alert,
		db:		 db,
		counter: guage,
	}
	return &st,nil
}

func (hs HourlySVC) GetPeriod(ctx context.Context, id string, start string, end string) (temps []hourlysvc.Temperature, err error) {
	level.Info(hs.logger).Log("msg", "GetPeriod", "ids", fmt.Sprintf("%s: %s-%s",id, start, end))
	temps, err = hs.db.GetPeriod(id, start, end)
	if err != nil {
		level.Error(hs.logger).Log("msg", "GetPeriod error", "err", err)
		hs.sendAlert(NewErrorAlert(err))
	}
	return temps,err
}

func (hs HourlySVC) PushPeriod(ctx context.Context, id string, temps []hourlysvc.Temperature) (err error){
	level.Info(hs.logger).Log("msg", "PushPeriod", "stId", id)
	err = hs.db.CreateTable(id)
	if err != nil {
		level.Error(hs.logger).Log("msg", "Station creation error", "err", err)
		hs.sendAlert(NewErrorAlert(err))
		return err
	}
	err = hs.db.PushPeriod(id, temps)
	if err != nil {
		level.Error(hs.logger).Log("msg", "PushPeriod error", "err", err)
		hs.sendAlert(NewErrorAlert(err))
	}
	return err
}

func (hs *HourlySVC) GetUpdateDate(ctx context.Context, ids []string) (dates map[string]string, err error) {
	level.Info(hs.logger).Log("msg", "GetUpdateDate", "ids", fmt.Sprintf("%+q:",ids))
	dates, err = hs.db.GetUpdateDateList(ids)
	if err != nil {
		level.Error(hs.logger).Log("msg", "Get Update Date List error", "err", err)
		hs.sendAlert(NewErrorAlert(err))
	}
	return dates, err
}

func (hs *HourlySVC) GetLatest(ctx context.Context, ids []string) (temps map[string]hourlysvc.Temperature, err error) {
	level.Info(hs.logger).Log("msg", "Get Latest", "ids", fmt.Sprintf("%+q:",ids))
	temps, err = hs.db.GetLatestList(ids)
	if err != nil {
		level.Error(hs.logger).Log("msg", "Get Latest List error", "err", err)
		hs.sendAlert(NewErrorAlert(err))
	}
	return temps, err
}

func (hs HourlySVC)sendAlert(alert alertsvc.Alert) {
	err := hs.alert.SendAlert(alert)
	if err != nil {
		level.Error(hs.logger).Log("msg", "Send Alert Error", "err", err)
	}
}