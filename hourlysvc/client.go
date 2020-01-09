package hourlysvc

import "github.com/flasherup/gradtage.de/hourlysvc/hrlgrpc"

type Client interface {
	GetPeriod		(id string, start string, end string) 	(resp *hrlgrpc.GetPeriodResponse, err error)
	PushPeriod		(id string, temps []Temperature) 		(resp *hrlgrpc.PushPeriodResponse, err error)
	GetUpdateDate	(ids []string) 							(resp *hrlgrpc.GetUpdateDateResponse, err error)
	GetLatest		(ids []string) 							(resp *hrlgrpc.GetLatestResponse, err error)
}