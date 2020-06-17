package impl

import (
	"context"
	"errors"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/alertsvc/grpcalt"
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

func NewSilentAlertSCVClient(host string, logger log.Logger) *AlertSVCClient {
	return &AlertSVCClient{
		logger:logger,
		host: host,
	}
}

func (scc AlertSVCClient) SendAlert(alert alertsvc.Alert)  error {
	conn := scc.openConn()
	defer conn.Close()

	client := grpcalt.NewAlertSVCClient(conn)
	resp, err := client.SendAlert(context.Background(), &grpcalt.SendAlertRequest{ Alert: encodeAlert(alert) })

	if err == nil && resp.Err != "nil" {
		err = errors.New(resp.Err)
	}

	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to send alert", "err", err)
	}
	return err
}

func (scc AlertSVCClient) SendEmail(email alertsvc.Email)  error {
	conn := scc.openConn()
	defer conn.Close()

	client := grpcalt.NewAlertSVCClient(conn)
	resp, err := client.SendEmail(context.Background(), &grpcalt.SendEmailRequest{ Email: alertsvc.EncodeEmail(email) })

	if err == nil && resp.Err != "nil" {
		err = errors.New(resp.Err)
	}

	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to send email", "err", err)
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

func encodeAlert(src alertsvc.Alert) *grpcalt.Alert {
	return &grpcalt.Alert{
		Name:	src.Name,
		Desc:	src.Desc,
		Params:	src.Params,
	}
}
