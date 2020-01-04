package hrlaggregatorsvc

import (
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/hagrpc"
)

type Client interface {
	GetStatus() (resp *hagrpc.GetStatusResponse, err error)
}