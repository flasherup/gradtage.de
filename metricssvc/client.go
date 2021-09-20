package metricssvc

import (
	"github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"
)

type Client interface {
	GetMetrics(ids []string) (map[string]*mtrgrpc.Metrics, error)
}
