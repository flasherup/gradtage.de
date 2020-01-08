package apisvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	GetHDDEndpoint  endpoint.Endpoint
	GetHDDSVEndpoint  endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetHDDEndpoint:   MakeGetHDDEndpoint(s),
		GetHDDSVEndpoint: MakeGetHDDCSVEndpoint(s),
	}
}

func MakeGetHDDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetHDDRequest)
		data, err := s.GetHDD(ctx, req.Params)
		return GetHDDResponse{ data}, err
	}
}

func MakeGetHDDCSVEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetHDDCSVRequest)
		data, filename, err := s.GetHDDCSV(ctx, req.Params)
		return GetHDDCSVResponse{ data, filename }, err
	}
}