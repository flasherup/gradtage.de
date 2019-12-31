package dailysvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetPeriodEndpoint  	endpoint.Endpoint
	PushPeriodEndpoint  endpoint.Endpoint
	GetUpdateDateEndpoint  	endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetPeriodEndpoint:   	MakeGetPeriodEndpoint(s),
		PushPeriodEndpoint:		MakePushPeriodEndpoint(s),
		GetUpdateDateEndpoint:	MakeGetUpdateDateEndpoint(s),
	}
}


func MakeGetPeriodEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPeriodRequest)
		temps, err := s.GetPeriod(ctx, req.ID, req.Start, req.End)
		return GetPeriodResponse{temps, err}, err
	}
}

func MakePushPeriodEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PushPeriodRequest)
		err := s.PushPeriod(ctx, req.ID, req.Temps)
		return PushPeriodResponse{ err}, err
	}
}

func MakeGetUpdateDateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUpdateDateRequest)
		dates, err := s.GetUpdateDate(ctx, req.IDs)
		return GetUpdateDateResponse{dates, err}, err
	}
}