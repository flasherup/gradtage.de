package hourlysvc

import (
	"context"
	"github.com/flasherup/gradtage.de/hourlysvc/hrlgrpc"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type GRPCServer struct {
	getPeriod    	gt.Handler
	pushPeriod  	gt.Handler
	getUpdateDate   gt.Handler
}

func (s *GRPCServer) GetPeriod(ctx context.Context, req *hrlgrpc.GetPeriodRequest) (*hrlgrpc.GetPeriodResponse, error) {
	_, resp, err := s.getPeriod.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*hrlgrpc.GetPeriodResponse), nil
}

func (s *GRPCServer) PushPeriod(ctx context.Context, req *hrlgrpc.PushPeriodRequest) (*hrlgrpc.PushPeriodResponse, error) {
	_, resp, err := s.pushPeriod.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*hrlgrpc.PushPeriodResponse), nil
}

func (s *GRPCServer) GetUpdateDate(ctx context.Context, req *hrlgrpc.GetUpdateDateRequest) (*hrlgrpc.GetUpdateDateResponse, error) {
	_, resp, err := s.getUpdateDate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*hrlgrpc.GetUpdateDateResponse), nil
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) hrlgrpc.HourlySVCServer {
	server := GRPCServer{
		getPeriod: gt.NewServer(
			endpoint.GetPeriodEndpoint,
			DecodeGetPeriodRequest,
			EncodeGetPeriodResponse,
		),
		pushPeriod: gt.NewServer(
			endpoint.PushPeriodEndpoint,
			DecodePushPeriodRequest,
			EncodePushPeriodResponse,
		),
		getUpdateDate: gt.NewServer(
			endpoint.GetUpdateDateEndpoint,
			DecodeGetUpdateDateRequest,
			EncodeGetUpdateDateResponse,
		),
	}
	return &server
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}