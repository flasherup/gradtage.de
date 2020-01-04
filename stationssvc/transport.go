package stationssvc

import (
	"context"
	"github.com/flasherup/gradtage.de/stationssvc/stsgrpc"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type GRPCServer struct {
	getStations    	gt.Handler
	getAllStations  gt.Handler
	addStations    	gt.Handler
}

func (s *GRPCServer) GetStations(ctx context.Context, req *stsgrpc.GetStationsRequest) (*stsgrpc.GetStationsResponse, error) {
	_, resp, err := s.getStations.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*stsgrpc.GetStationsResponse), nil
}

func (s *GRPCServer) GetAllStations(ctx context.Context, req *stsgrpc.GetAllStationsRequest) (*stsgrpc.GetAllStationsResponse, error) {
	_, resp, err := s.getAllStations.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*stsgrpc.GetAllStationsResponse), nil
}

func (s *GRPCServer) AddStations(ctx context.Context, req *stsgrpc.AddStationsRequest) (*stsgrpc.AddStationsResponse, error) {
	_, resp, err := s.addStations.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*stsgrpc.AddStationsResponse), nil
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) stsgrpc.StationSVCServer {
	return &GRPCServer{
		getStations: gt.NewServer(
			endpoint.GetStationsEndpoint,
			DecodeGetStationsRequest,
			EncodeGetStationsResponse,
		),
		getAllStations: gt.NewServer(
			endpoint.GetAllStationsEndpoint,
			DecodeGeAllStationsRequest,
			EncodeGetAllStationsResponse,
		),
		addStations: gt.NewServer(
			endpoint.AddStationsEndpoint,
			DecodeAddStationsRequest,
			EncodeAddStationsResponse,
		),
	}
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}