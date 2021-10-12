package impl

import (
	"context"
	"errors"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	googlerpc "google.golang.org/grpc"
)

type MetricsSVCClient struct{
	logger     log.Logger
	host string
}

func NewWeatherBitSVCClient(host string, logger log.Logger) *MetricsSVCClient {
	logger = log.With(logger,
		"client", "metrics",
	)
	return &MetricsSVCClient{
		logger:logger,
		host: host,
	}
}

func (wb MetricsSVCClient) GetMetrics(ids []string) (map[string]*mtrgrpc.Metrics, error) {
	conn := wb.openConn()
	defer conn.Close()

	client := mtrgrpc.NewMetricsSVCClient(conn)
	resp, err := client.GetMetrics(context.Background(), &mtrgrpc.GetMetricsRequest{ Ids: ids })
	if err != nil {
		level.Error(wb.logger).Log("msg", "Failed to get period", "err", err)
	}else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}

	return resp.Metrics, err
}


func (wb MetricsSVCClient) openConn() *googlerpc.ClientConn {
	options := googlerpc.WithDefaultCallOptions(
			googlerpc.MaxCallRecvMsgSize(common.MaxMessageReceiveSize),
			googlerpc.MaxCallSendMsgSize(common.MaxMessageSendSize),
		)
	cc, err := googlerpc.Dial(wb.host, googlerpc.WithInsecure(), options)
	if err != nil {
		level.Error(wb.logger).Log("msg", "Failed to start gRPC connection", "err", err)
	}
	return cc
}
