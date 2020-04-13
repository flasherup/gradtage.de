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

func (ds DailySVC) GetUpdateDate(ctx context.Context, ids []string) (dates map[string]string, err error) {
	level.Info(ds.logger).Log("msg", "GetUpdateDate", "idsAmount", len(ids))
	ids = ds.getValidIds(ids)
	dates, err = ds.db.GetUpdateDateList(ids)
	if err != nil {
		level.Error(ds.logger).Log("msg", "Get Update Date error", "err", err)
		ds.sendAlert(NewErrorAlert(err))
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

func (ds DailySVC) getValidIds(ids []string) []string {
	tableNames, err := ds.db.GetTablesList()
	if err != nil {
		level.Error(ds.logger).Log("msg", "Can not validate ids", "err", err)
		ds.sendAlert(NewErrorAlert(err))
		return ids
	}
	valid := make(map[string]bool, 0)
	for _,v := range ids {
		_,ok := tableNames[v]
		if ok {
			valid[v] = true
		} else {
			level.Warn(ds.logger).Log("msg", "Id not found", "id", v)
		}
	}

	res := make([]string, len(valid))
	i := 0
	for id,_ := range valid {
		res[i] = id
		i++
	}

	return res
}

func (ds DailySVC)sendAlert(alert alertsvc.Alert) {
	err := ds.alert.SendAlert(alert)
	if err != nil {
		level.Error(ds.logger).Log("msg", "Send Alert Error", "err", err)
	}
}