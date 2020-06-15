package apisvc

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	GetHDDEndpoint  		endpoint.Endpoint
	GetHDDSVEndpoint  		endpoint.Endpoint
	GetSourceDataEndpoint	endpoint.Endpoint
	SearchEndpoint			endpoint.Endpoint
	UserEndpoint			endpoint.Endpoint
	PlanEndpoint			endpoint.Endpoint
	StripeEndpoint			endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetHDDEndpoint:   		MakeGetHDDEndpoint(s),
		GetHDDSVEndpoint: 		MakeGetHDDCSVEndpoint(s),
		GetSourceDataEndpoint:  MakeGetSourceDataEndpoint(s),
		SearchEndpoint:  		MakeSearchEndpoint(s),
		UserEndpoint:  			MakeUserEndpoint(s),
		PlanEndpoint:  			MakePlanEndpoint(s),
		StripeEndpoint:  		MakeStripeEndpoint(s),
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

func MakeGetSourceDataEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSourceDataRequest)
		data, filename, err := s.GetSourceData(ctx, req.Params)
		return GetSourceDataResponse{ data, filename }, err
	}
}

func MakeSearchEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchRequest)
		data, err := s.Search(ctx, req.Params)
		return SearchResponse{ data}, err
	}
}

func MakeUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("MakeUserEndpoint")
		req := request.(UserRequest)
		data, err := s.User(ctx, req.Params)
		return UserResponse{ data}, err
	}
}

func MakePlanEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PlanRequest)
		data, err := s.Plan(ctx, req.Params)
		return PlanResponse{ data}, err
	}
}

func MakeStripeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StripeRequest)
		data, err := s.Stripe(ctx, req.Event)
		return StripeResponse{ data}, err
	}
}