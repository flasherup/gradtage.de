package weatherbitsvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetPeriodEndpoint        endpoint.Endpoint
	GetWBPeriod              endpoint.Endpoint
	PushWBPeriod             endpoint.Endpoint
	GetUpdateDateEndpoint    endpoint.Endpoint
	GetStationsListEndpoint  endpoint.Endpoint
	GetAverageEndpoint       endpoint.Endpoint
	GetAverageDegreeEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetPeriodEndpoint:        MakeGetPeriodEndpoint(s),
		PushWBPeriod:             MakePushWBPeriodEndpoint(s),
		GetWBPeriod:              MakeGetWBPeriodEndpoint(s),
		GetUpdateDateEndpoint:    MakeGetUpdateDateEndpoint(s),
		GetStationsListEndpoint:  MakeGetStationsListEndpoint(s),
		GetAverageEndpoint:       MakeGetAverageEndpoint(s),
		GetAverageDegreeEndpoint: MakeGetAverageDegreeEndpoint(s),
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

func MakePushWBPeriodEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PushWBPeriodRequest)
		err := s.PushWBPeriod(ctx, req.Id, req.Data)
		return PushWBPeriodResponse{err}, err
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

func MakeGetAverageEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAverageRequest)
		temps, err := s.GetAverage(ctx, req.Id, req.Years, req.End)
		return GetAverageResponse{temps, err}, err
	}
}

func MakeGetAverageDegreeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAverageDegreeRequest)
		temps, err := s.GetAverageDegree(ctx, req.Params, req.Years)
		return GetAverageDegreeResponse{temps, err}, err
	}
}