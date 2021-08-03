package impl

import (
	"context"
	"errors"
	weathergrpc "github.com/flasherup/gradtage.de/weatherbitsvc/weatherbitgrpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	googlerpc "google.golang.org/grpc"
)

type WeatherBitSVCClient struct{
	logger     log.Logger
	host string
}

func NewWeatherBitSVCClient(host string, logger log.Logger) *WeatherBitSVCClient {
	return &WeatherBitSVCClient{
		logger:logger,
		host: host,
	}
}

func (wb WeatherBitSVCClient) GetPeriod(ids []string, start string, end string) (resp *weathergrpc.GetPeriodResponse, err error) {
	conn := wb.openConn()
	defer conn.Close()

	client := weathergrpc.NewWeatherBitScraperSVCClient(conn)
	resp, err = client.GetPeriod(context.Background(), &weathergrpc.GetPeriodRequest{ Ids: ids, Start:start, End:end })
	if err != nil {
		level.Error(wb.logger).Log("msg", "Failed to get period", "err", err)
	}else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}

	return resp, err
}

func (wb WeatherBitSVCClient) GetWBPeriod(ids []string, start string, end string) (resp *weathergrpc.GetPeriodResponse, err error) {
	conn := wb.openConn()
	defer conn.Close()

	client := weathergrpc.NewWeatherBitScraperSVCClient(conn)
	resp, err = client.GetPeriod(context.Background(), &weathergrpc.GetPeriodRequest{ Ids: ids, Start:start, End:end })
	if err != nil {
		level.Error(wb.logger).Log("msg", "Failed to get period", "err", err)
	}else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}

	return resp, err
}

func (wb WeatherBitSVCClient) GetUpdateDate(ids []string) (resp *weathergrpc.GetUpdateDateResponse, err error) {
	conn := wb.openConn()
	defer conn.Close()

	client := weathergrpc.NewWeatherBitScraperSVCClient(conn)
	resp, err = client.GetUpdateDate(context.Background(), &weathergrpc.GetUpdateDateRequest{ Ids: ids })
	if err != nil {
		level.Error(wb.logger).Log("msg", "Failed to get update date", "err", err)
	} else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}
	return resp, err
}

func (wb WeatherBitSVCClient) openConn() *googlerpc.ClientConn {
	cc, err := googlerpc.Dial(wb.host, googlerpc.WithInsecure())
	if err != nil {
		level.Error(wb.logger).Log("msg", "Failed to start gRPC connection", "err", err)
	}
	return cc
}
