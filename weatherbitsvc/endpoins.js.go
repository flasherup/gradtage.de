package weatherbitsvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetPeriodEndpoint       endpoint.Endpoint
	GetWBPeriod             endpoint.Endpoint
	GetUpdateDateEndpoint   endpoint.Endpoint
	GetStationsListEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetPeriodEndpoint:       MakeGetPeriodEndpoint(s),
		GetWBPeriod:             MakeGetWBPeriodEndpoint(s),
		GetUpdateDateEndpoint:   MakeGetUpdateDateEndpoint(s),
		GetStationsListEndpoint: MakeGetStationsListEndpoint(s),
	}
}

func MakeGetPeriodEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPeriodRequest)
		temps, err := s.GetPeriod(ctx, req.IDs, req.Start, req.End)
		return GetPeriodResponse{temps, err}, err
	}
}

func MakeGetWBPeriodEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetWBPeriodRequest)
		temps, err := s.GetWBPeriod(ctx, req.Id, req.Start, req.End)
		return GetWBPeriodResponse{temps, err}, err
	}

}

func MakeGetUpdateDateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUpdateDateRequest)
		dates, err := s.GetUpdateDate(ctx, req.IDs)
		return GetUpdateDateResponse{dates, err}, err
	}
}

func MakeGetStationsListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		dates, err := s.GetStationsList(ctx)
		return GetStationsListResponse{dates, err}, err
	}
}
