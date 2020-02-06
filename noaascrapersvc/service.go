package noaascrapersvc

import (
	"context"
	"github.com/flasherup/gradtage.de/hourlysvc"
)

type Service interface {
	GetPeriod(ctx context.Context, id string, start string, end string) (temps []hourlysvc.Temperature, err error)
	GetUpdateDate(ctx context.Context, ids []string) (dates map[string]string ,err error)
}