package weatherbitsvc

import (
	"github.com/flasherup/gradtage.de/common"
	weathergrpc "github.com/flasherup/gradtage.de/weatherbitsvc/weatherbitgrpc"
)

type Client interface {
	GetPeriod(ids []string, start string, end string) (temps *map[string][]common.Temperature, err error)
	GetWBPeriod(id string, start string, end string) (resp *[]WBData, err error)
	PushWBPeriod(id string, data []WBData) (err error)
	GetUpdateDate(ids []string) (resp *weathergrpc.GetUpdateDateResponse, err error)
	GetStationsList() (resp *[]string, err error)
}
