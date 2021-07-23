package weatherbitsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/hourlysvc"
)

type Service interface {
	GetPeriod(ctx context.Context, ids []string, start string, end string) (temps map[string][]hourlysvc.Temperature, err error)
	GetUpdateDate(ctx context.Context, ids []string) (dates map[string]string ,err error)
}