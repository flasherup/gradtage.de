package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dailysvc/grpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	googlerpc "google.golang.org/grpc"
)

type DailySVCClient struct{
	logger     log.Logger
	host string
}

func NewDailySCVClient(host string, logger log.Logger) *DailySVCClient {
	return &DailySVCClient{
		logger:logger,
		host: host,
	}
}

func (scc DailySVCClient) GetPeriod(id string, start string, end string) *grpc.GetPeriodResponse {
	conn := scc.openConn()
	defer conn.Close()

	client := grpc.NewDailySVCClient(conn)
	res, err := client.GetPeriod(context.Background(), &grpc.GetPeriodRequest{ Id:id, Start:start, End:end })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get period", "err", err)

	}
	return res
}

func (scc DailySVCClient) PushPeriod(id string, temps []dailysvc.Temperature) *grpc.PushPeriodResponse {
	conn := scc.openConn()
	defer conn.Close()

	client := grpc.NewDailySVCClient(conn)
	tGRPC := toGRPCTemps(temps)
	res, err := client.PushPeriod(context.Background(), &grpc.PushPeriodRequest{Id:id, Temps:tGRPC})
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to push period", "err", err)

	}
	return res
}

func (scc DailySVCClient) GetUpdateDate(ids []string) *grpc.GetUpdateDateResponse {
	conn := scc.openConn()
	defer conn.Close()

	client := grpc.NewDailySVCClient(conn)
	res, err := client.GetUpdateDate(context.Background(), &grpc.GetUpdateDateRequest{ Ids:ids })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get update date", "err", err)

	}
	return res
}


func (scc DailySVCClient) openConn() *googlerpc.ClientConn {
	cc, err := googlerpc.Dial(scc.host, googlerpc.WithInsecure())
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to start gRPC connection", "err", err)
	}
	return cc
}

func toGRPCTemps(sts []dailysvc.Temperature) []*grpc.Temperature {
	res := make([]*grpc.Temperature, len(sts))
	for i,v := range sts {
		res[i] = &grpc.Temperature{
			Date:			v.Date,
			Temperature:	float32(v.Temperature),
		}
	}

	return res
}
