package weatherbitsvc

import (
	"context"
	weathergrpc "github.com/flasherup/gradtage.de/weatherbitsvc/weatherbitgrpc"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

type GRPCServer struct {
	getPeriod    	gt.Handler
}

func (s *GRPCServer) GetPeriod(ctx context.Context, req *weathergrpc.GetPeriodRequest) (*weathergrpc.GetPeriodResponse, error) {
	_, resp, err := s.getPeriod.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*weathergrpc.GetPeriodResponse), err
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) weathergrpc.WeatherBitScraperSVCServer{
	server := GRPCServer{
		getPeriod: gt.NewServer(
			endpoint.GetPeriodEndpoint,
			DecodeGetPeriodRequest,
			EncodeGetPeriodResponse,
		),
	}
	return &server
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}