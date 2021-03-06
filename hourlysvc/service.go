package hourlysvc

import (
	"context"
)


type Temperature struct {
	Date 		string `json:"date"`
	Temperature	float64 `json:"temperature"`
}

type Service interface {
	GetPeriod(ctx context.Context, id string, start string, end string) (temps []Temperature, err error)
	PushPeriod(ctx context.Context, id string, temps []Temperature) (err error)
	GetUpdateDate(ctx context.Context, ids []string) (dates map[string]string ,err error)
	GetLatest(ctx context.Context, ids []string) (dates map[string]Temperature ,err error)
}