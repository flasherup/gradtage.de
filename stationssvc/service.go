package stationssvc

import (
	"context"
)


type Station struct {
	ID 			string `json:"id"`
	Name 		string `json:"name"`
	Timezone 	string `json:timezone`
	SourceType	string `json:source_type`
	SourceID	string `json:source_id`

}

type Service interface {
	GetStations(ctx context.Context, ids []string) (sts map[string]Station, err error)
	GetAllStations(ctx context.Context) (sts map[string]Station, err error)
	GetStationsBySrcType(ctx context.Context, types []string) (sts map[string]Station, err error)
	AddStations(ctx context.Context, sts []Station) (err error)
}