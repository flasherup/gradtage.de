package impl

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/apisvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dailysvc/dlygrpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"math"
	"time"
)

type APISVC struct {
	logger  	log.Logger
	alert 		alertsvc.Client
	daily		dailysvc.Client
	counter 	ktprom.Gauge
}

const (
	HDDType = "hdd"
	DDType  = "dd"
)

func NewAPISVC(logger log.Logger, daily dailysvc.Client, alert alertsvc.Client) *APISVC {
	options := prometheus.Opts{
		Name: "stations_count_total",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := APISVC{
		logger:  logger,
		daily:	 daily,
		alert:   alert,
		counter: *guage,
	}
	return &st
}

func (as APISVC) GetHDD(ctx context.Context, params apisvc.Params) (data [][]string, err error) {
	level.Info(as.logger).Log("msg", "GetHDD", "station", params.Station)
	temps, err := as.daily.GetPeriod(params.Station, params.Start, params.End)
	if err != nil {
		level.Error(as.logger).Log("msg", "GetHDD error", "err", err)
		as.sendAlert(NewErrorAlert(err))
	}

	avg, err := as.daily.GetAvg(params.Station)
	if err != nil {
		level.Error(as.logger).Log("msg", "GetHDD error", "err", err)
		as.sendAlert(NewErrorAlert(err))
	}

	headerCSV := []string{ "ID","Date","HDD","HDDAverage" }
	csv := as.generateCSV(headerCSV, temps.Temps, avg.Temps, params)
	return csv, err
}

func (as APISVC) GetHDDCSV(ctx context.Context, params apisvc.Params) (data [][]string, fileName string, err error) {
	level.Info(as.logger).Log("msg", "GetHDD", "station", params.Station)
	temps, err := as.daily.GetPeriod(params.Station, params.Start, params.End)
	if err != nil {
		level.Error(as.logger).Log("msg", "GetHDD error", "err", err)
		as.sendAlert(NewErrorAlert(err))
	}

	avg, err := as.daily.GetAvg(params.Station)
	if err != nil {
		level.Error(as.logger).Log("msg", "GetHDD error", "err", err)
		as.sendAlert(NewErrorAlert(err))
	}

	var headerCSV []string
	if params.Output == DDType {
		headerCSV = []string{ "ID", "Date", "DD", "DDAverage" }
	} else {
		headerCSV = []string{ "ID","Date","HDD","HDDAverage" }
	}

	csv := as.generateCSV(headerCSV, temps.Temps, avg.Temps, params)
	fileName = fmt.Sprintf("%s%s%s%g%g.csv", params.Station, params.Start, params.End, params.HL, params.RT)
	return csv,fileName,err
}


func (as APISVC)generateCSV(names []string, temps []*dlygrpc.Temperature, tempsAvg map[int32]*dlygrpc.Temperature, params apisvc.Params) [][]string {
	res := [][]string{ names }
	var line []string
	var degree float64
	var degreeA float64
	for _, v := range temps {
		d, err := time.Parse(common.TimeLayout, v.Date)
		if err != nil {
			level.Error(as.logger).Log("msg", "GetHDD generateCSV error", "err", err)
			as.sendAlert(NewErrorAlert(err))
		}
		doy := int32(d.YearDay())

		aTemperature := tempsAvg[doy].Temperature

		if params.Output 		== HDDType {
			degree 	= calculateHDD(params.HL, float64(v.Temperature))
			degreeA = calculateHDD(params.HL, float64(aTemperature))
		} else if params.Output == DDType {
			degree 	= calculateDD(params.HL, params.RT, float64(v.Temperature))
			degreeA = calculateDD(params.HL, params.RT, float64(aTemperature))
		}

		line = []string{
			params.Station,
			v.Date,
			fmt.Sprintf("%.1f", degree),
			fmt.Sprintf("%.1f", degreeA),
		}

		res = append(res, line)
	}
	return res
}

func ToFixed(x float64) float64 {
	unit := 10.0
	return math.Round(x*unit) / unit
}

func calculateHDD(baseHDD float64, value float64) float64 {
	if value >= baseHDD {
		return 0
	}
	return baseHDD - value
}

func calculateDD(baseHDD float64, baseDD float64, value float64) float64 {
	if value >= baseHDD || value >= baseDD{
		return 0
	}

	return baseDD - value
}

func (as APISVC)sendAlert(alert alertsvc.Alert) {
	_, err := as.alert.SendAlert(alert)
	if err != nil {
		level.Error(as.logger).Log("msg", "Send Alert Error", "err", err)
	}
}

