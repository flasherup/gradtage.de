package weatherbitsvc

import (
	"context"
	weathergrpc "github.com/flasherup/gradtage.de/weatherbitsvc/weatherbitgrpc"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type GRPCServer struct {
	getPeriod        gt.Handler
	getWbPeriod      gt.Handler
	pushWbPeriod     gt.Handler
	getUpdateDate    gt.Handler
	getStationsList  gt.Handler
	getAverage       gt.Handler
}

func (s *GRPCServer) GetPeriod(ctx context.Context, req *weathergrpc.GetPeriodRequest) (*weathergrpc.GetPeriodResponse, error) {
	_, resp, err := s.getPeriod.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*weathergrpc.GetPeriodResponse), err
}

func (s *GRPCServer) GetWBPeriod(ctx context.Context, req *weathergrpc.GetWBPeriodRequest) (request *weathergrpc.GetWBPeriodResponse, err error) {
	_, resp, err := s.getWbPeriod.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*weathergrpc.GetWBPeriodResponse), err
}

func (s *GRPCServer) PushWBPeriod(ctx context.Context, req *weathergrpc.PushWBPeriodRequest) (request *weathergrpc.PushWBPeriodResponse, err error) {
	_, resp, err := s.pushWbPeriod.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*weathergrpc.PushWBPeriodResponse), err
}

func (s *GRPCServer) GetUpdateDate(ctx context.Context, req *weathergrpc.GetUpdateDateRequest) (*weathergrpc.GetUpdateDateResponse, error) {
	_, resp, err := s.getUpdateDate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*weathergrpc.GetUpdateDateResponse), err
}

func (s *GRPCServer) GetStationsList(ctx context.Context, req *weathergrpc.GetStationsListRequest) (*weathergrpc.GetStationsListResponse, error) {
	_, resp, err := s.getStationsList.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*weathergrpc.GetStationsListResponse), err
}

func (s *GRPCServer) GetAverage(ctx context.Context, req *weathergrpc.GetAverageRequest) (*weathergrpc.GetAverageResponse, error) {
	_, resp, err := s.getAverage.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*weathergrpc.GetAverageResponse), err
}

func NewGRPCServer(_ context.Context, endpoint Endpoints) *GRPCServer {
	server := GRPCServer{
		getPeriod: gt.NewServer(
			endpoint.GetPeriodEndpoint,
			DecodeGetPeriodRequest,
			EncodeGetPeriodResponse,
		),
		getWbPeriod: gt.NewServer(
			endpoint.GetWBPeriod,
			DecodeGetWBPeriodRequest,
			EncodeGetWBPeriodResponse,
		),
		pushWbPeriod: gt.NewServer(
			endpoint.PushWBPeriod,
			DecodePushWBPeriodRequest,
			EncodePushWBPeriodResponse,
		),
		getUpdateDate: gt.NewServer(
			endpoint.GetUpdateDateEndpoint,
			DecodeGetUpdateDateRequest,
			EncodeGetUpdateDateResponse,
		),
		getStationsList: gt.NewServer(
			endpoint.GetStationsListEndpoint,
			DecodeGetStationsListRequest,
			EncodeGetStationsListResponse,
		),
		getAverage: gt.NewServer(
			endpoint.GetAverageEndpoint,
			DecodeGetAverageRequest,
			EncodeGetAverageResponse,
		),
	}
	return &server
}

func NewMetricsTransport(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}
