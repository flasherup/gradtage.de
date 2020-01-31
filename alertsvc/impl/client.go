package impl

import (
	"context"
	"errors"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/alertsvc/altgrpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	googlerpc "google.golang.org/grpc"
)

type AlertSVCClient struct{
	logger     log.Logger
	host string
}

func NewAlertSCVClient(host string, logger log.Logger) *AlertSVCClient {
	return &AlertSVCClient{
		logger:logger,
		host: host,
	}
}

func (scc AlertSVCClient) SendAlert(alert alertsvc.Alert)  error {
	conn := scc.openConn()
	defer conn.Close()

	client := altgrpc.NewAlertSVCClient(conn)
	resp, err := client.SendAlert(context.Background(), &altgrpc.SendAlertRequest{ Alert: encodeAlert(alert) })

	if err == nil && resp.Err != "nil" {
		err = errors.New(resp.Err)
	}

	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get stations", "err", err)
	}
	return err
}


func (scc AlertSVCClient) openConn() *googlerpc.ClientConn {
	cc, err := googlerpc.Dial(scc.host, googlerpc.WithInsecure())
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to start gRPC connection", "err", err)
	}
	return cc
}

func encodeAlert(src alertsvc.Alert) *altgrpc.Alert {
	return &altgrpc.Alert{
		Name:	src.Name,
		Desc:	src.Desc,
		Params:	src.Params,
	}
}
