package stationssvc

import "github.com/flasherup/gradtage.de/stationssvc/stsgrpc"

type Client interface {
	GetAutocomplete(text string) (resp *stsgrpc.GetStationsResponse, err error)
}