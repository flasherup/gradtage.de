package impl

import (
	"context"
	"errors"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/hourlysvc/hrlgrpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	googlerpc "google.golang.org/grpc"
)

type HourlySVCClient struct{
	logger     log.Logger
	host string
}

func NewHourlySCVClient(host string, logger log.Logger) *HourlySVCClient {
	return &HourlySVCClient{
		logger:logger,
		host: host,
	}
}

func (scc HourlySVCClient) GetPeriod(id string, start string, end string) (resp *hrlgrpc.GetPeriodResponse, err error) {
	conn := scc.openConn()
	defer conn.Close()

	client := hrlgrpc.NewHourlySVCClient(conn)
	resp, err = client.GetPeriod(context.Background(), &hrlgrpc.GetPeriodRequest{ Id: id, Start:start, End:end })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get period", "err", err)
	}else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}

	return resp, err
}

func (scc HourlySVCClient) PushPeriod(id string, temps []hourlysvc.Temperature) (resp *hrlgrpc.PushPeriodResponse, err error) {
	conn := scc.openConn()
	defer conn.Close()

	client := hrlgrpc.NewHourlySVCClient(conn)
	tGRPC := toGRPCTemps(temps)
	resp, err = client.PushPeriod(context.Background(), &hrlgrpc.PushPeriodRequest{Id: id, Temps:tGRPC})
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to push period", "err", err)
	}else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}
	return resp, err
}

func (scc HourlySVCClient) GetUpdateDate(ids []string) (resp *hrlgrpc.GetUpdateDateResponse, err error) {
	conn := scc.openConn()
	defer conn.Close()

	client := hrlgrpc.NewHourlySVCClient(conn)
	resp, err = client.GetUpdateDate(context.Background(), &hrlgrpc.GetUpdateDateRequest{ Ids: ids })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get update date", "err", err)
	} else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}
	return resp, err
}

func (scc HourlySVCClient) GetLatest(ids []string) (resp *hrlgrpc.GetLatestResponse, err error) {
	conn := scc.openConn()
	defer conn.Close()

	client := hrlgrpc.NewHourlySVCClient(conn)
	resp, err = client.GetLatest(context.Background(), &hrlgrpc.GetLatestRequest{ Ids: ids })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get latest", "err", err)
	} else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}
	return resp, err
}


func (scc HourlySVCClient) openConn() *googlerpc.ClientConn {
	cc, err := googlerpc.Dial(scc.host, googlerpc.WithInsecure())
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to start gRPC connection", "err", err)
	}
	return cc
}

func toGRPCTemps(sts []hourlysvc.Temperature) []*hrlgrpc.Temperature {
	res := make([]*hrlgrpc.Temperature, len(sts))
	for i,v := range sts {
		res[i] = &hrlgrpc.Temperature{
			Date:			v.Date,
			Temperature:	float32(v.Temperature),
		}
	}

	return res
}
