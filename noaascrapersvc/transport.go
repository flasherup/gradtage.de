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
	forceOverrideHourly   	gt.Handler
}

func (s *GRPCServer) ForceOverrideHourly(ctx context.Context, req *noaascpc.ForceOverrideHourlyRequest) (*noaascpc.ForceOverrideHourlyResponse, error) {
	_, resp, err := s.forceOverrideHourly.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*noaascpc.ForceOverrideHourlyResponse), nil
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) noaascpc.NoaaScraperSVCServer{
	server := GRPCServer{
		forceOverrideHourly: gt.NewServer(
			endpoint.ForceOverrideHourlyEndpoint,
			DecodeForceOverrideHourlyRequest,
			EncodeForceOverrideHourlyResponse,
		),
	}
	return &server
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}