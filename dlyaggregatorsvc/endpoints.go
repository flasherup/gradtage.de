package dlyaggregatorsvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetStatusEndpoint  	endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetStatusEndpoint:   	MakeGetStatusEndpoint(s),
	}
}

func MakeGetStatusEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		status, err := s.GetStatus(ctx)
		return GetStatusResponse{status, err}, err
	}
}