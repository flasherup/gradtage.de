package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/stationssvc/stsgrpc"
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

func (scc StationsSVCClient) GetStations(ids []string) (resp *stsgrpc.GetStationsResponse, err error) {
	conn := scc.openConn()
	defer conn.Close()

	client := stsgrpc.NewStationSVCClient(conn)
	resp, err = client.GetStations(context.Background(), &stsgrpc.GetStationsRequest{ Ids: ids })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get stations", "err", err)

	}
	return resp, err
}

func (scc StationsSVCClient) GetAllStations() (resp *stsgrpc.GetAllStationsResponse, err error) {
	conn := scc.openConn()
	defer conn.Close()

	client := stsgrpc.NewStationSVCClient(conn)
	resp, err = client.GetAllStations(context.Background(), &stsgrpc.GetAllStationsRequest{})
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get all stations", "err", err)

	}
	return resp, err
}

func (scc StationsSVCClient) AddStations(sts []stationssvc.Station) (resp *stsgrpc.AddStationsResponse, err error) {
	conn := scc.openConn()
	defer conn.Close()

	s := toGRPCStations(sts)
	client := stsgrpc.NewStationSVCClient(conn)
	resp, err = client.AddStations(context.Background(), &stsgrpc.AddStationsRequest{ Sts: s })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get stations", "err", err)

	}
	return resp, err
}


func (scc StationsSVCClient) openConn() *googlerpc.ClientConn {
	cc, err := googlerpc.Dial(scc.host, googlerpc.WithInsecure())
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to start gRPC connection", "err", err)
	}
	return cc
}

func toGRPCStations(sts []stationssvc.Station) []*stsgrpc.Station {
	res := make([]*stsgrpc.Station, len(sts))
	for i,v := range sts {
		res[i] = &stsgrpc.Station{
			Id:v.ID,
			Name:v.Name,
			Timezone:v.Timezone,
			SourceType:v.SourceType,
			SourceId:v.SourceID,
		}
	}

	return res
}
