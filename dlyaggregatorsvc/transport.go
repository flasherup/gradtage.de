package dlyaggregatorsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/dagrpc"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type GRPCServer struct {
	getStatus    	gt.Handler
}

func (s *GRPCServer) GetStatus(ctx context.Context, req *dagrpc.GetStatusRequest) (*dagrpc.GetStatusResponse, error) {
	_, resp, err := s.getStatus.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*dagrpc.GetStatusResponse), nil
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) dagrpc.DlyAggregatorSVCServer {
	server := GRPCServer{
		getStatus: gt.NewServer(
			endpoint.GetStatusEndpoint,
			DecodeGetStatusRequest,
			EncodeGetStatusResponse,
		),
	}
	return &server
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}