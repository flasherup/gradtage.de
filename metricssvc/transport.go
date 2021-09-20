package weatherbitupdatesvc

import (
	"context"
	"github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type GRPCServer struct {
	getMetrics       gt.Handler
}

func (s *GRPCServer) GetMetrics(ctx context.Context, req *mtrgrpc.GetMetricsRequest) (*mtrgrpc.GetMetricsResponse, error) {
	_, resp, err := s.getMetrics.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*mtrgrpc.GetMetricsResponse), err
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) *GRPCServer {
	server := GRPCServer{
		getMetrics: gt.NewServer(
			endpoint.GetMetricsEndpoint,
			DecodeGetMetricsRequest,
			EncodeGetMetricsResponse,
		),
	}
	return &server
}

func NewMetricsTransport(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}
