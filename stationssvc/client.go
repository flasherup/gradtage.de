package stationssvc

import "github.com/flasherup/gradtage.de/stationssvc/grpc"

type Client interface {
	GetStations(ids []string) *grpc.GetStationsResponse
	GetAllStations() *grpc.GetAllStationsResponse
	AddStations(sts []Station) *grpc.AddStationsResponse
}