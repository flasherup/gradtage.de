package impl

import (
	"context"
	"errors"
	"github.com/flasherup/gradtage.de/noaascrapersvc/noaascpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	googlerpc "google.golang.org/grpc"
)

type NoaaScraperSVCClient struct{
	logger     log.Logger
	host string
}

func NewNoaaScraperSVCClient(host string, logger log.Logger) *NoaaScraperSVCClient {
	return &NoaaScraperSVCClient{
		logger:logger,
		host: host,
	}
}

func (scc NoaaScraperSVCClient) GetPeriod(id string, start string, end string) (resp *noaascpc.GetPeriodResponse, err error) {
	conn := scc.openConn()
	defer conn.Close()

	client := noaascpc.NewNoaaScraperSVCClient(conn)
	resp, err = client.GetPeriod(context.Background(), &noaascpc.GetPeriodRequest{ Id: id, Start:start, End:end })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get period", "err", err)
	}else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}

	return resp, err
}

func (scc NoaaScraperSVCClient) GetUpdateDate(ids []string) (resp *noaascpc.GetUpdateDateResponse, err error) {
	conn := scc.openConn()
	defer conn.Close()

	client := noaascpc.NewNoaaScraperSVCClient(conn)
	resp, err = client.GetUpdateDate(context.Background(), &noaascpc.GetUpdateDateRequest{ Ids: ids })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get update date", "err", err)
	} else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}
	return resp, err
}

func (scc NoaaScraperSVCClient) openConn() *googlerpc.ClientConn {
	cc, err := googlerpc.Dial(scc.host, googlerpc.WithInsecure())
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to start gRPC connection", "err", err)
	}
	return cc
}
