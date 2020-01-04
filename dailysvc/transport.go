package dailysvc

import (
	"context"
	"github.com/flasherup/gradtage.de/dailysvc/grpc"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type GRPCServer struct {
	getPeriod    	 gt.Handler
	pushPeriod  	 gt.Handler
	getUpdateDate    gt.Handler
	updateAvgForYear gt.Handler
	updateAvgForDOY  gt.Handler
	getAvg			 gt.Handler

}

func (s *GRPCServer) GetPeriod(ctx context.Context, req *grpc.GetPeriodRequest) (*grpc.GetPeriodResponse, error) {
	_, resp, err := s.getPeriod.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpc.GetPeriodResponse), nil
}

func (s *GRPCServer) PushPeriod(ctx context.Context, req *grpc.PushPeriodRequest) (*grpc.PushPeriodResponse, error) {
	_, resp, err := s.pushPeriod.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpc.PushPeriodResponse), nil
}

func (s *GRPCServer) GetUpdateDate(ctx context.Context, req *grpc.GetUpdateDateRequest) (*grpc.GetUpdateDateResponse, error) {
	_, resp, err := s.getUpdateDate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpc.GetUpdateDateResponse), nil
}

func (s *GRPCServer) UpdateAvgForYear(ctx context.Context, req *grpc.UpdateAvgForYearRequest) (*grpc.UpdateAvgForYearResponse, error) {
	_, resp, err := s.updateAvgForYear.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpc.UpdateAvgForYearResponse), nil
}

func (s *GRPCServer) UpdateAvgForDOY(ctx context.Context, req *grpc.UpdateAvgForDOYRequest) (*grpc.UpdateAvgForDOYResponse, error) {
	_, resp, err := s.updateAvgForDOY.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpc.UpdateAvgForDOYResponse), nil
}

func (s *GRPCServer) GetAvg(ctx context.Context, req *grpc.GetAvgRequest) (*grpc.GetAvgResponse, error) {
	_, resp, err := s.getAvg.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpc.GetAvgResponse), nil
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) grpc.DailySVCServer {
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
		updateAvgForYear: gt.NewServer(
			endpoint.UpdateAvgForYearEndpoint,
			DecodeUpdateAvgForYearRequest,
			EncodeUpdateAvgForYearResponse,
		),
		updateAvgForDOY: gt.NewServer(
			endpoint.UpdateAvgForDOYEndpoint,
			DecodeUpdateAvgForDOYRequest,
			EncodeUpdateAvgForDOYResponse,
		),
		getAvg: gt.NewServer(
			endpoint.GetAvgEndpoint,
			DecodeGetAvgRequest,
			EncodeGetAvgResponse,
		),
	}
	return &server
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}