package stationssvc

import (
	"context"
)


type Station struct {
	ID 			string `json:"id"`
	Name 		string `json:"name"`
	Timezone 	string `json:timezone`

}

type Service interface {
	GetStations(ctx context.Context, ids []string) (sts map[string]Station, err error)
	GetAllStations(ctx context.Context) (sts map[string]Station, err error)
	AddStations(ctx context.Context, sts []Station) (err error)
}