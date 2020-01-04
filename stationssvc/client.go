package stationssvc

import "github.com/flasherup/gradtage.de/stationssvc/stsgrpc"

type Client interface {
	GetStations(ids []string) 	(resp *stsgrpc.GetStationsResponse, err error)
	GetAllStations() 			(resp *stsgrpc.GetAllStationsResponse, err error)
	AddStations(sts []Station) 	(resp *stsgrpc.AddStationsResponse, err error)
}