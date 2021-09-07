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
	forceRestart       gt.Handler
}

func (s *GRPCServer) ForceRestart(ctx context.Context, req *weathergrpc.GetPeriodRequest) (*weathergrpc.GetPeriodResponse, error) {
	_, resp, err := s.forceRestart.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*weathergrpc.GetPeriodResponse), err
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) *GRPCServer {
	server := GRPCServer{
		forceRestart: gt.NewServer(
			endpoint.ForceRestartEndpoint,
			DecodeForceRestartRequest,
			EncodeForceRestartResponse,
		),
	}
	return &server
}

func NewMetricsTransport(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}
