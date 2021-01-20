package weatherbitsvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetPeriodEndpoint  		endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetPeriodEndpoint:   	MakeGetPeriodEndpoint(s),
	}
}


func MakeGetPeriodEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPeriodRequest)
		temps, err := s.GetPeriod(ctx, req.IDs, req.Start, req.End)
		return GetPeriodResponse{temps, err}, err
	}
}