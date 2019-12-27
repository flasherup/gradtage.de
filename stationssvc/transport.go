package stationssvc

import (
	"context"
	"github.com/flasherup/gradtage.de/stationssvc/grpc"
	gt "github.com/go-kit/kit/transport/grpc"
)

type GRPCServer struct {
	getStations    	gt.Handler
	getAllStations  gt.Handler
	addStations    	gt.Handler
}

func (s *GRPCServer) GetStations(ctx context.Context, req *grpc.GetStationsRequest) (*grpc.GetStationsResponse, error) {
	_, resp, err := s.getStations.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpc.GetStationsResponse), nil
}

func (s *GRPCServer) GetAllStations(ctx context.Context, req *grpc.GetAllStationsRequest) (*grpc.GetAllStationsResponse, error) {
	_, resp, err := s.getAllStations.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpc.GetAllStationsResponse), nil
}

func (s *GRPCServer) AddStations(ctx context.Context, req *grpc.AddStationsRequest) (*grpc.AddStationsResponse, error) {
	_, resp, err := s.addStations.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpc.AddStationsResponse), nil
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) grpc.StationSVCServer {
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