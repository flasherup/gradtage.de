package metricssvc

import (
	"context"
	"github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"
)

type Service interface {
	GetMetrics(ctx context.Context, ids []string) (map[string]*mtrgrpc.Metrics, error)
}
