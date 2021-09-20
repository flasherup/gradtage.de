package weatherbitupdatesvc

import (
	"context"
	"github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"
)

type Service interface {
	GetMetrics(ctx context.Context, ids []string) ([]*mtrgrpc.Metrics, error)
}
