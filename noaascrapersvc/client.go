package noaascrapersvc

import (
	"github.com/flasherup/gradtage.de/noaascrapersvc/noaascpc"
)

type Client interface {
	GetPeriod		(id string, start string, end string) 		(resp *noaascpc.GetPeriodResponse, err error)
	GetUpdateDate	(ids []string) 								(resp *noaascpc.GetUpdateDateResponse, err error)
}