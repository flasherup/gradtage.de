package daydegreesvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetDegreeEndpoint       endpoint.Endpoint
	GetAverageDegreeEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetDegreeEndpoint:       MakeGetDegreeEndpoint(s),
		GetAverageDegreeEndpoint: MakeGetAverageDegreeEndpoint(s),
	}
}

func MakeGetDegreeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDegreeRequest)
		temps, err := s.GetDegree(ctx, req.Params)
		return GetDegreeResponse{temps, err}, err
	}
}

func MakeGetAverageDegreeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAverageDegreeRequest)
		temps, err := s.GetAverageDegree(ctx, req.Params, req.Years)
		return GetAverageDegreeResponse{temps, err}, err
	}
}