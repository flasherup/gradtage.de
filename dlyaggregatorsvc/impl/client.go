package impl

import (
	"context"
	"errors"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/dagrpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	googlerpc "google.golang.org/grpc"
)

type DailyAggregatorSVCClient struct{
	logger     log.Logger
	host string
}

func NewDailyAggregatorSCVClient(host string, logger log.Logger) *DailyAggregatorSVCClient {
	return &DailyAggregatorSVCClient{
		logger:logger,
		host: host,
	}
}

func (scc DailyAggregatorSVCClient) ForceUpdate(ids []string, start string, end string) (err error)   {
	conn := scc.openConn()
	defer conn.Close()

	client := dagrpc.NewDlyAggregatorSVCClient(conn)
	resp, err := client.ForceUpdate(context.Background(), &dagrpc.ForceUpdateRequest{ Ids: ids, Start:start, End:end })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get period", "err", err)
	} else if resp.Err != "nil" {
		level.Error(scc.logger).Log("msg", "Failed to get period", "err", errors.New(resp.Err))
	}
	return err
}

func (scc DailyAggregatorSVCClient) openConn() *googlerpc.ClientConn {
	cc, err := googlerpc.Dial(scc.host, googlerpc.WithInsecure())
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to start gRPC connection", "err", err)
	}
	return cc
}
