package impl

import (
	"context"
	"fmt"
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
	db 			database.DailyDB
	avg			*average.Average
	counter 	ktprom.Gauge
}

func NewDailySVC(logger log.Logger, db database.DailyDB, avg *average.Average) (*DailySVC, error) {
	options := prometheus.Opts{
		Name: "stations_count_total",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := DailySVC{
		logger: logger,
		db:		db,
		avg:	avg,
		counter: *guage,
	}
	return &st,nil
}

func (ss DailySVC) GetPeriod(ctx context.Context, id string, start string, end string) (temps []dailysvc.Temperature, err error) {
	level.Info(ss.logger).Log("msg", "GetPeriod", "ids", fmt.Sprintf("%s: %s-%s",id, start, end))
	temps, err = ss.db.GetPeriod(id, start, end)
	if err != nil {
		level.Error(ss.logger).Log("msg", "GetPeriod error", "err", err)
	}
	return temps,err
}

func (ss DailySVC) PushPeriod(ctx context.Context, id string, temps []dailysvc.Temperature) (err error){
	level.Info(ss.logger).Log("msg", "PushPeriod", "stId", id)
	err = ss.db.CreateTable(id)
	if err != nil {
		level.Error(ss.logger).Log("msg", "Station creation error", "err", err)
		return err
	}
	err = ss.db.PushPeriod(id, temps)
	if err != nil {
		level.Error(ss.logger).Log("msg", "PushPeriod error", "err", err)
	}
	return err
}

func (ss *DailySVC) GetUpdateDate(ctx context.Context, ids []string) (dates map[string]string, err error) {
	level.Info(ss.logger).Log("msg", "GetUpdateDate", "ids", fmt.Sprintf("%+q:",ids))
	dates = make(map[string]string)
	for _,v := range ids {
		date, err := ss.db.GetUpdateDate(v)
		if err != nil {
			level.Error(ss.logger).Log("msg", "Get Update Date error", "err", err)
		} else {
			dates[v] = date
		}
	}
	return dates, err
}

func (ss DailySVC) UpdateAvgForYear(ctx context.Context, id string) (err error) {
	level.Info(ss.logger).Log("msg", " UpdateAvgForYear", "id", id)
	err = ss.avg.CalculateAndSaveYearlyAverage(id)
	if err != nil {
		level.Error(ss.logger).Log("msg", "UpdateAvgForYear error", "err", err)
	}
	return err
}

func (ss DailySVC) UpdateAvgForDOY(ctx context.Context, id string, doy int) (err error) {
	level.Info(ss.logger).Log("msg", " UpdateAvgForDOY", "id", id, "doy", doy)
	err = ss.avg.CalculateAndSaveDOYAverage(id, doy)
	if err != nil {
		level.Error(ss.logger).Log("msg", "UpdateAvgForDOY error", "err", err)
	}
	return err
}

func (ss DailySVC) GetAvg(ctx context.Context, id string) (temps []dailysvc.Temperature, err error) {
	level.Info(ss.logger).Log("msg", "GetAvg", "id", id)
	temps, err = ss.avg.GetAll(id)
	if err != nil {
		level.Error(ss.logger).Log("msg", "GetAvg error", "err", err)
	}
	return temps,err
}