package weatherbitsvc

import (
	weathergrpc "github.com/flasherup/gradtage.de/weatherbitsvc/weatherbitgrpc"
)

type Client interface {
	GetPeriod		(ids []string, start string, end string) 		(resp *weathergrpc.GetPeriodResponse, err error)
	GetWBPeriod		(id string, start string, end string) 			(resp *[]WBData, err error)
	GetUpdateDate	(ids []string) 									(resp *weathergrpc.GetUpdateDateResponse, err error)
}