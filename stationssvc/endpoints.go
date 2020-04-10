package stationssvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	GetStationsEndpoint  			endpoint.Endpoint
	GetAllStationsEndpoint  		endpoint.Endpoint
	GetStationsBySrcTypeEndpoint  	endpoint.Endpoint
	AddStationsEndpoint  			endpoint.Endpoint
	ResetStationsEndpoint  			endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetStationsEndpoint:   			MakeGetStationsEndpoint(s),
		GetAllStationsEndpoint:			MakeGetAllStationsEndpoint(s),
		GetStationsBySrcTypeEndpoint:	MakeGetStationsBySrcTypeEndpoint(s),
		AddStationsEndpoint:			MakeAddStationsEndpoint(s),
		ResetStationsEndpoint:			MakeResetStationsEndpoint(s),
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
		return GetAllStationsResponse{sts, err}, nil
	}
}

func MakeGetStationsBySrcTypeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetStationsBySrcTypeRequest)
		sts, err := s.GetStationsBySrcType(ctx, req.Types)
		return GetStationsBySrcTypeResponse{sts, err}, nil
	}
}

func MakeAddStationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddStationsRequest)
		err := s.AddStations(ctx, req.Stations)
		return AddStationsResponse{err}, nil
	}
}


func MakeResetStationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ResetStationsRequest)
		err := s.ResetStations(ctx, req.Stations)
		return ResetStationsResponse{err}, nil
	}
}