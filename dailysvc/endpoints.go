package dailysvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetPeriodEndpoint  			endpoint.Endpoint
	PushPeriodEndpoint  		endpoint.Endpoint
	GetUpdateDateEndpoint  		endpoint.Endpoint
	UpdateAvgForYearEndpoint 	endpoint.Endpoint
	UpdateAvgForDOYEndpoint 	endpoint.Endpoint
	GetAvgEndpoint 				endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetPeriodEndpoint:   		MakeGetPeriodEndpoint(s),
		PushPeriodEndpoint:			MakePushPeriodEndpoint(s),
		GetUpdateDateEndpoint:		MakeGetUpdateDateEndpoint(s),
		UpdateAvgForYearEndpoint: 	MakeUpdateAvgForYearEndpoint(s),
		UpdateAvgForDOYEndpoint:	MakeUpdateAvgForDOYEndpoint(s),
		GetAvgEndpoint: 			MakeGetAvgEndpoint(s),
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

func MakeUpdateAvgForYearEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAvgForYearRequest)
		err := s.UpdateAvgForYear(ctx, req.ID)
		return UpdateAvgForYearResponse{ err }, err
	}
}

func MakeUpdateAvgForDOYEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAvgForDOYRequest)
		err := s.UpdateAvgForDOY(ctx, req.ID, req.DOY)
		return UpdateAvgForDOYResponse{ err }, err
	}
}

func MakeGetAvgEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAvgRequest)
		temps, err := s.GetAvg(ctx, req.ID)
		return GetAvgResponse{temps, err}, err
	}
}