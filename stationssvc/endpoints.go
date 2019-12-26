package stationssvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	GetStationsEndpoint  	endpoint.Endpoint
	GetAllStationsEndpoint  endpoint.Endpoint
	AddStationsEndpoint  	endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetStationsEndpoint:   	MakeGetStationsEndpoint(s),
		GetAllStationsEndpoint:	MakeGetAllStationsEndpoint(s),
		AddStationsEndpoint:	MakeAddStationsEndpoint(s),
	}
}


func MakeGetStationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetStationsRequest)
		sts, err := s.GetStations(ctx, req.IDs)
		return GetStationsResponse{sts, err}, nil
	}
}

func MakeGetAllStationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		sts, err := s.GetAllStations(ctx)
		return GetStationsResponse{sts, err}, nil
	}
}

func MakeAddStationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddStationsRequest)
		err := s.AddStations(ctx, req.Stations)
		return AddStationsResponse{err}, nil
	}
}