package database

import "github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"

type MetricsDB interface {
	GetMetrics(ids []string) (map[string]*mtrgrpc.Metrics, error)
	PushMetrics(metrics map[string]*mtrgrpc.Metrics) error
	CreateTable() error
	RemoveTable() error
	Dispose()
}
