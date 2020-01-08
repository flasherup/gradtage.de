package dailysvc

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
	UpdateAvgForYear(ctx context.Context, id string) (err error)
	UpdateAvgForDOY(ctx context.Context, id string, doy int) (err error)
	GetAvg(ctx context.Context, id string) (temps map[int]Temperature, err error)
}