package noaascrapersvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ForceOverrideHourlyEndpoint  	endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		ForceOverrideHourlyEndpoint:   	MakeForceOverrideHourlyEndpoint(s),
	}
}

func MakeForceOverrideHourlyEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ForceOverrideHourlyRequest)
		err := s.ForceOverrideHourly(ctx, req.Station, req.Start, req.End)
		return ForceOverrideHourlyResponse{err}, err
	}
}