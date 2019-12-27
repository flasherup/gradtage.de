package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/stationssvc/grpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	googlerpc "google.golang.org/grpc"
)

type StationsSVCClient struct{
	logger     log.Logger
	host string
}

func NewStationsSCVClient(host string, logger log.Logger) *StationsSVCClient {
	return &StationsSVCClient{
		logger:logger,
		host: host,
	}
}

func (scc StationsSVCClient) GetStations(ids []string) *grpc.GetStationsResponse {
	conn := scc.openConn()
	defer conn.Close()

	client := grpc.NewStationSVCClient(conn)
	res, err := client.GetStations(context.Background(), &grpc.GetStationsRequest{ Ids:ids })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get stations", "err", err)

	}
	return res
}

func (scc StationsSVCClient) GetAllStations() *grpc.GetAllStationsResponse {
	conn := scc.openConn()
	defer conn.Close()

	client := grpc.NewStationSVCClient(conn)
	res, err := client.GetAllStations(context.Background(), &grpc.GetAllStationsRequest{})
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get all stations", "err", err)

	}
	return res
}

func (scc StationsSVCClient) AddStations(sts []stationssvc.Station) *grpc.AddStationsResponse {
	conn := scc.openConn()
	defer conn.Close()

	s := toGRPCStations(sts)
	client := grpc.NewStationSVCClient(conn)
	res, err := client.AddStations(context.Background(), &grpc.AddStationsRequest{ Sts:s })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get stations", "err", err)

	}
	return res
}


func (scc StationsSVCClient) openConn() *googlerpc.ClientConn {
	cc, err := googlerpc.Dial(scc.host, googlerpc.WithInsecure())
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to start gRPC connection", "err", err)
	}
	return cc
}

func toGRPCStations(sts []stationssvc.Station) []*grpc.Station {
	res := make([]*grpc.Station, len(sts))
	for i,v := range sts {
		res[i] = &grpc.Station{
			Id:v.ID,
			Name:v.Name,
			Timezone:v.Timezone,
		}
	}

	return res
}
