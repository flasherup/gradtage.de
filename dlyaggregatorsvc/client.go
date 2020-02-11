package dlyaggregatorsvc

import "github.com/flasherup/gradtage.de/dlyaggregatorsvc/dagrpc"

type Client interface {
	ForceUpdate() (resp *dagrpc.ForceUpdateResponse, err error)

}