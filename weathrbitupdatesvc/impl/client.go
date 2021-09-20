package impl

import (
	"context"
	"errors"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/wbugrpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	googlerpc "google.golang.org/grpc"
)

type WeatherBitSVCClient struct{
	logger     log.Logger
	host string
}

func NewWeatherBitUpdateSVCClient(host string, logger log.Logger) *WeatherBitSVCClient {
	logger = log.With(logger,
		"client", "weatherbitupdate",
	)
	return &WeatherBitSVCClient{
		logger:logger,
		host: host,
	}
}

func (wb WeatherBitSVCClient) ForceRestart(ctx context.Context) error{
	conn := wb.openConn()
	defer conn.Close()

	client := wbugrpc.NewWeatherBitUpdateSVCClient(conn)
	resp, err := client.ForceRestart(context.Background(), &wbugrpc.ForceRestartRequest{})
	if err != nil {
		level.Error(wb.logger).Log("msg", "Failed to get period", "err", err)
	}else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}

	return err
}



func (wb WeatherBitSVCClient) openConn() *googlerpc.ClientConn {
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
