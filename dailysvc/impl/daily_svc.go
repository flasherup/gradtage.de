package impl

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dailysvc/impl/average"
	"github.com/flasherup/gradtage.de/dailysvc/impl/database"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

type DailySVC struct {
	logger  	log.Logger
	alert 		alertsvc.Client
	db 			database.DailyDB
	avg			*average.Average
	counter 	ktprom.Gauge
}

func NewDailySVC(logger log.Logger, db database.DailyDB, avg *average.Average, alert alertsvc.Client) (*DailySVC, error) {
	options := prometheus.Opts{
		Name: "stations_count_total",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := DailySVC{
		logger: logger,
		alert: alert,
		db:		db,
		avg:	avg,
		counter: *guage,
	}
	return &st,nil
}

func (ds DailySVC) GetPeriod(ctx context.Context, id string, start string, end string) (temps []dailysvc.Temperature, err error) {
	level.Info(ds.logger).Log("msg", "GetPeriod", "ids", fmt.Sprintf("%s: %s-%s",id, start, end))
	temps, err = ds.db.GetPeriod(id, start, end)
	if err != nil {
		level.Error(ds.logger).Log("msg", "GetPeriod error", "err", err)
		ds.sendAlert(NewErrorAlert(err))
	}
	return temps,err
}

func (ds DailySVC) PushPeriod(ctx context.Context, id string, temps []dailysvc.Temperature) (err error){
	level.Info(ds.logger).Log("msg", "PushPeriod", "stId", id)
	err = ds.db.CreateTable(id)
	if err != nil {
		level.Error(ds.logger).Log("msg", "Station creation error", "err", err)
		ds.sendAlert(NewErrorAlert(err))
		return err
	}
	err = ds.db.PushPeriod(id, temps)
	if err != nil {
		level.Error(ds.logger).Log("msg", "PushPeriod error", "err", err)
		ds.sendAlert(NewErrorAlert(err))
	}
	return err
}

func (ds *DailySVC) GetUpdateDate(ctx context.Context, ids []string) (dates map[string]string, err error) {
	level.Info(ds.logger).Log("msg", "GetUpdateDate", "ids", fmt.Sprintf("%+q:",ids))
	dates = make(map[string]string)
	for _,v := range ids {
		date, err := ds.db.GetUpdateDate(v)
		if err != nil {
			level.Error(ds.logger).Log("msg", "Get Update Date error", "err", err)
			ds.sendAlert(NewErrorAlert(err))
		} else {
			dates[v] = date
		}
	}
	return dates, err
}

func (ds DailySVC) UpdateAvgForYear(ctx context.Context, id string) (err error) {
	level.Info(ds.logger).Log("msg", " UpdateAvgForYear", "id", id)
	err = ds.avg.CalculateAndSaveYearlyAverage(id)
	if err != nil {
		level.Error(ds.logger).Log("msg", "UpdateAvgForYear error", "err", err)
		ds.sendAlert(NewErrorAlert(err))
	}
	return err
}

func (ds DailySVC) UpdateAvgForDOY(ctx context.Context, id string, doy int) (err error) {
	level.Info(ds.logger).Log("msg", " UpdateAvgForDOY", "id", id, "doy", doy)
	err = ds.avg.CalculateAndSaveDOYAverage(id, doy)
	if err != nil {
		level.Error(ds.logger).Log("msg", "UpdateAvgForDOY error", "err", err)
		ds.sendAlert(NewErrorAlert(err))
	}
	return err
}

func (ds DailySVC) GetAvg(ctx context.Context, id string) (temps map[int]dailysvc.Temperature, err error) {
	level.Info(ds.logger).Log("msg", "GetAvg", "id", id)
	temps, err = ds.avg.GetAll(id)
	if err != nil {
		level.Error(ds.logger).Log("msg", "GetAvg error", "err", err)
		ds.sendAlert(NewErrorAlert(err))
	}
	return temps,err
}

func (ds DailySVC)sendAlert(alert alertsvc.Alert) {
	_, err := ds.alert.SendAlert(alert)
	if err != nil {
		level.Error(ds.logger).Log("msg", "Send Alert Error", "err", err)
	}
}