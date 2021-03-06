package noaascrapersvc

import (
	"context"
	"github.com/flasherup/gradtage.de/noaascrapersvc/noaascpc"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type GRPCServer struct {
	getPeriod    	gt.Handler
	getUpdateDate   gt.Handler
}

func (s *GRPCServer) GetPeriod(ctx context.Context, req *noaascpc.GetPeriodRequest) (*noaascpc.GetPeriodResponse, error) {
	_, resp, err := s.getPeriod.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*noaascpc.GetPeriodResponse), err
}

func (s *GRPCServer) GetUpdateDate(ctx context.Context, req *noaascpc.GetUpdateDateRequest) (*noaascpc.GetUpdateDateResponse, error) {
	_, resp, err := s.getUpdateDate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*noaascpc.GetUpdateDateResponse), err
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) noaascpc.NoaaScraperSVCServer{
	server := GRPCServer{
		getPeriod: gt.NewServer(
			endpoint.GetPeriodEndpoint,
			DecodeGetPeriodRequest,
			EncodeGetPeriodResponse,
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