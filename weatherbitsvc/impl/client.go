package impl

import (
	"context"
	"errors"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
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

func (wb WeatherBitSVCClient) GetWBPeriod( id string, start string, end string) (resp *[]weatherbitsvc.WBData, err error) {
	conn := wb.openConn()
	defer conn.Close()

	client := weathergrpc.NewWeatherBitScraperSVCClient(conn)
	grpc, err := client.GetWBPeriod(context.Background(), &weathergrpc.GetWBPeriodRequest{ Id:id, Start:start, End:end })
	if err != nil {
		level.Error(wb.logger).Log("msg", "Failed to get WB period", "err", err)
	} else if grpc.Err != common.ErrorNilString {
		err = common.ErrorFromString(grpc.Err);
	} else {
		resp = weatherbitsvc.ToWBData(grpc.Temps)
	}

	return resp, err
}

func (wb WeatherBitSVCClient) PushWBPeriod( id string, data []weatherbitsvc.WBData) (err error) {
	conn := wb.openConn()
	defer conn.Close()
	client := weathergrpc.NewWeatherBitScraperSVCClient(conn)
	grpcData := weatherbitsvc.ToGRPCWBData(data)
	grpc, errRequest := client.PushWBPeriod(context.Background(), &weathergrpc.PushWBPeriodRequest{ Id:id, Data:grpcData })
	if errRequest != nil {
		level.Error(wb.logger).Log("msg", "Failed to get WB period", "err", errRequest)
		err = errRequest
	} else if grpc.Err != common.ErrorNilString {
		err = common.ErrorFromString(grpc.Err);
	}
	return
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

func (wb WeatherBitSVCClient) GetStationsList() (stations *[]string, err error) {
	conn := wb.openConn()
	defer conn.Close()

	client := weathergrpc.NewWeatherBitScraperSVCClient(conn)
	grpc, err := client.GetStationsList(context.Background(), &weathergrpc.GetStationsListRequest{})
	if err != nil {
		level.Error(wb.logger).Log("msg", "Failed to get stations list", "err", err)
	} else if grpc.Err != common.ErrorNilString {
		err = common.ErrorFromString(grpc.Err);
	} else {
		stations = &grpc.List
	}
	return stations, err
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
