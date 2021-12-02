package main

import (
	"github.com/flasherup/gradtage.de/metricssvc/impl"
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
			"svc", "metricsclient",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	client := impl.NewMetricsSVCClient("82.165.119.83:8114",logger)
	//client := impl.NewMetricsSVCClient("localhost:8114",logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	testGetMetrics(client, logger)
}

func testGetMetrics(client *impl.MetricsSVCClient, logger log.Logger) {
	level.Info(logger).Log("msg", "satart get metrics test")

	metrics, err := client.GetMetrics([]string{"us_kmis", "us_kmgc"})
	if err != nil {
		level.Error(logger).Log("msg", "Get metrics error", "error", err)
		return
	}

	level.Info(logger).Log("msg", "Get metrics test complete", "Response Length", len(metrics))

	for k,v := range metrics {
		level.Info(logger).Log("msg", "Metric", "id", k, "date", v.Date, "LastUpdate", v.LastUpdate, "FirstUpdate", v.FirstUpdate, "RecordsAll", v.RecordsAll, "RecordsClean", v.RecordsClean)
	}
}

