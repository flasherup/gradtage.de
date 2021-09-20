package impl

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/metricssvc/config"
	"github.com/flasherup/gradtage.de/metricssvc/impl/database"
	"github.com/flasherup/gradtage.de/metricssvc/impl/utils"
	"github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type UpdateStart struct {
	StId string
}

type UpdateResult struct {
	StId    string
	Error   error
	Message string
}

type metrics struct {
	updateCounter   *ktprom.Counter
	requestsCounter *ktprom.Counter
}

type MetricsSVC struct {
	weatherbit   weatherbitsvc.Client
	db           database.MetricsDB
	alert        alertsvc.Client
	logger       log.Logger
	conf         config.MetricsConfig
	startChanel  chan UpdateStart
	resultChanel chan UpdateResult
	metrics      metrics
}

const (
	labelStation = "station"
	labelStatus  = "status"
)

func NewMetricsSVC(
	logger log.Logger,
	weatherbit weatherbitsvc.Client,
	db database.MetricsDB,
	alert alertsvc.Client,
	conf config.MetricsConfig,
) (*MetricsSVC, error) {
	wb := MetricsSVC{
		weatherbit:            weatherbit,
		db:                    db,
		alert:                 alert,
		logger:                logger,
		conf:                  conf,
		startChanel:           make(chan UpdateStart),
		resultChanel:          make(chan UpdateResult),
	}

	wb.metrics = *setupMetrics()

	go handleUpdates(&wb)
	go startFetchProcess(&wb)
	return &wb, nil
}

func setupMetrics() *metrics {
	updateCounter := ktprom.NewCounterFrom(
		prometheus.CounterOpts(
			prometheus.Opts{
				Name: "metrics_update_counter",
				Help: "The total number of success/error updates",
			},
		),
		[]string{labelStatus},
	)

	requestCounter := ktprom.NewCounterFrom(
		prometheus.CounterOpts(
			prometheus.Opts{
				Name: "metrics_request_counter",
				Help: "The total number of requests",
			},
		),
		[]string{labelStation},
	)

	return &metrics{
		updateCounter:   updateCounter,
		requestsCounter: requestCounter,
	}
}

func (ms MetricsSVC) GetMetrics(ctx context.Context, ids []string) (map[string]*mtrgrpc.Metrics, error) {
	return ms.db.GetMetrics(ids)
}

func handleUpdates(wbu *MetricsSVC) {
	for {
		select {
		case updateResult := <-wbu.resultChanel:
			if updateResult.Error != nil {
				wbu.metrics.updateCounter.With(labelStatus, "error").Add(1)
				level.Error(wbu.logger).Log("msg", updateResult.Message, "err", updateResult.Error)
			} else {
				wbu.metrics.updateCounter.With(labelStatus, "success").Add(1)
			}
		case updateStart := <-wbu.startChanel:
			wbu.metrics.requestsCounter.With(labelStation, updateStart.StId).Add(1)
		}
	}
}

func startFetchProcess(ms *MetricsSVC) {
	err := ms.db.CreateTable()
	if err != nil {
		level.Error(ms.logger).Log("msg", "Cant create table", "error", err)
		return
	}

	for {
		stations, err := ms.weatherbit.GetStationsList()
		if err != nil {
			level.Error(ms.logger).Log("msg", "Cant get stations list", "error", err)
			time.Sleep(time.Minute)
			continue
		}

		level.Info(ms.logger).Log("msg", "Start update process", "Stations", len(*stations))
		ms.precessStations(*stations)
	}
}

func (ms *MetricsSVC) precessStations(sts []string) {
	for i, v := range sts {
		level.Info(ms.logger).Log("msg", "Process station", "index", i, "station", v)
		ms.processUpdate(v)
	}
}

func (ms *MetricsSVC) processUpdate(stID string) {
	ms.startChanel <- UpdateStart{
		StId: stID,
	}

	now := time.Now()
	wbData, err := ms.weatherbit.GetPeriod([]string{stID}, common.TimeVeryFirstWBH , now.Format(common.TimeLayoutWBH))
	if err != nil {
		ms.resultChanel <- UpdateResult{
			StId:    stID,
			Message: fmt.Sprintf("weatherbit data get error for station '%s' ", stID),
			Error:   err,
		}
		return
	}

	temps := (*wbData)[stID]

	metrics, metricsErr := utils.GetWeatherbitMetrics(&temps)
	if metricsErr != nil {
		ms.resultChanel <- UpdateResult{
			StId:    stID,
			Message: fmt.Sprintf("get weatherbit metrics error for station '%s' ", stID),
			Error:   metricsErr,
		}
		return
	}

	err = ms.db.PushMetrics(map[string]*mtrgrpc.Metrics{stID:metrics})
	if err != nil {
		ms.resultChanel <- UpdateResult{
			StId:    stID,
			Message: fmt.Sprintf("push weatherbit metrics error for station '%s' ", stID),
			Error:   err,
		}
		return
	}

	ms.resultChanel <- UpdateResult{
		StId:    stID,
		Message: fmt.Sprintf("metrics update success '%s' ", stID),
		Error:   nil,
	}
}
