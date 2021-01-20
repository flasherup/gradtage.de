package weatherbitsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/hourlysvc"
)
type Service interface {
	GetPeriod(ctx context.Context, ids []string, start string, end string) (temps map[string][]hourlysvc.Temperature, err error)
}