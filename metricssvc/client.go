package weatherbitupdatesvc

import (
	"github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"
)

type Client interface {
	GetMetrics(ids []string) ([]*mtrgrpc.Metrics, error)
}
