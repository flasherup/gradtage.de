package dlyaggregatorsvc

import "github.com/flasherup/gradtage.de/dlyaggregatorsvc/dagrpc"

type Client interface {
	GetStatus() (resp *dagrpc.GetStatusResponse, err error)

}