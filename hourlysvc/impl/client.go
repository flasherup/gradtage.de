package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/hourlysvc/grpc"
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

func (scc HourlySVCClient) GetPeriod(id string, start string, end string) *grpc.GetPeriodResponse {
	conn := scc.openConn()
	defer conn.Close()

	client := grpc.NewHourlySVCClient(conn)
	res, err := client.GetPeriod(context.Background(), &grpc.GetPeriodRequest{ Id:id, Start:start, End:end })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get period", "err", err)

	}
	return res
}

func (scc HourlySVCClient) PushPeriod(id string, temps []hourlysvc.Temperature) *grpc.PushPeriodResponse {
	conn := scc.openConn()
	defer conn.Close()

	client := grpc.NewHourlySVCClient(conn)
	tGRPC := toGRPCTemps(temps)
	res, err := client.PushPeriod(context.Background(), &grpc.PushPeriodRequest{Id:id, Temps:tGRPC})
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to push period", "err", err)

	}
	return res
}

func (scc HourlySVCClient) GetUpdateDate(ids []string) *grpc.GetUpdateDateResponse {
	conn := scc.openConn()
	defer conn.Close()

	client := grpc.NewHourlySVCClient(conn)
	res, err := client.GetUpdateDate(context.Background(), &grpc.GetUpdateDateRequest{ Ids:ids })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get update date", "err", err)

	}
	return res
}


func (scc HourlySVCClient) openConn() *googlerpc.ClientConn {
	cc, err := googlerpc.Dial(scc.host, googlerpc.WithInsecure())
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to start gRPC connection", "err", err)
	}
	return cc
}

func toGRPCTemps(sts []hourlysvc.Temperature) []*grpc.Temperature {
	res := make([]*grpc.Temperature, len(sts))
	for i,v := range sts {
		res[i] = &grpc.Temperature{
			Date:			v.Date,
			Temperature:	v.Temperature,
		}
	}

	return res
}
