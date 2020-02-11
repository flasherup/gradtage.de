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
	forceUpdate   	gt.Handler
}

func (s *GRPCServer) ForceUpdate(ctx context.Context, req *dagrpc.ForceUpdateRequest) (*dagrpc.ForceUpdateResponse, error) {
	_, resp, err := s.forceUpdate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*dagrpc.ForceUpdateResponse), nil
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) dagrpc.DlyAggregatorSVCServer {
	server := GRPCServer{
		forceUpdate: gt.NewServer(
			endpoint.ForceUpdateEndpoint,
			DecodeForceUpdateRequest,
			EncodeForceUpdateResponse,
		),
	}
	return &server
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}