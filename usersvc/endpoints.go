package usersvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	CreateUserEndpoint  		endpoint.Endpoint
	CreateUserAutoEndpoint  	endpoint.Endpoint
	SetPlanEndpoint  			endpoint.Endpoint
	SetStationsEndpoint  		endpoint.Endpoint
	ValidateKeyEndpoint  		endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUserEndpoint:   		MakeCreateUserEndpoint(s),
		CreateUserAutoEndpoint:   	MakeCreateUserAutoEndpoint(s),
		SetPlanEndpoint:   			MakeSetPlanEndpoint(s),
		SetStationsEndpoint:   		MakeSetStationsEndpoint(s),
		ValidateKeyEndpoint:   		MakeValidateKeyEndpoint(s),
	}
}


func MakeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		err := s.CreateUser(ctx, req.UserName,req.Plan)
		return CreateUserResponse{err}, err
	}
}

func MakeCreateUserAutoEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserAutoRequest)
		err := s.CreateUserAuto(ctx, req.UserName,req.Plan)
		return CreateUserAutoResponse{err}, err
	}
}

func MakeSetPlanEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetPlanRequest)
		err := s.SetPlan(ctx, req.UserName,req.Plan)
		return SetPlanResponse{err}, err
	}
}

func MakeSetStationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SetStationsRequest)
		err := s.SetStations(ctx, req.UserName,req.Station)
		return SetStationsResponse{err}, err
	}
}

func MakeValidateKeyEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateKeyRequest)
		parameters, err := s.ValidateKey(ctx, req.Key)
		return ValidateKeyResponse{ parameters, err}, err
	}
}