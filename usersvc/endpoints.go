package usersvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUserEndpoint  		endpoint.Endpoint
	UpdateUserEndpoint  		endpoint.Endpoint
	AddPlanEndpoint  			endpoint.Endpoint
	ValidateSelectionEndpoint  	endpoint.Endpoint
	ValidateKeyEndpoint  		endpoint.Endpoint
	ValidateNameEndpoint  		endpoint.Endpoint
	ValidateStripeEndpoint  	endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUserEndpoint:   		MakeCreateUserEndpoint(s),
		UpdateUserEndpoint:   		MakeUpdateUserEndpoint(s),
		AddPlanEndpoint:   			MakeAddPlanEndpoint(s),
		ValidateSelectionEndpoint:  MakeValidateSelectionEndpoint(s),
		ValidateKeyEndpoint:   		MakeValidateKeyEndpoint(s),
		ValidateNameEndpoint:   	MakeValidateNameEndpoint(s),
		ValidateStripeEndpoint:   	MakeValidateStripeEndpoint(s),
	}
}


func MakeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		key, err := s.CreateUser(ctx, req.UserName,req.Plan, req.Email)
		return CreateUserResponse{key, err}, err
	}
}

func MakeUpdateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		key, err := s.UpdateUser(ctx, req.User, req.Email)
		return UpdateUserResponse{key, err}, err
	}
}

func MakeAddPlanEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddPlanRequest)
		err := s.AddPlan(ctx, req.Plan)
		return AddPlanResponse{err}, err
	}
}

func MakeValidateSelectionEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateSelectionRequest)
		isValid, err := s.ValidateSelection(ctx, req.Selection)
		return ValidateSelectionResponse{ isValid, err}, err
	}
}

func MakeValidateKeyEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateKeyRequest)
		parameters, err := s.ValidateKey(ctx, req.Key)
		return ValidateKeyResponse{ parameters, err}, err
	}
}

func MakeValidateNameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateNameRequest)
		parameters, err := s.ValidateName(ctx, req.Name)
		return ValidateNameResponse{ parameters, err}, err
	}
}

func MakeValidateStripeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateStripeRequest)
		parameters, err := s.ValidateStripe(ctx, req.Stripe)
		return ValidateStripeResponse{ parameters, err}, err
	}
}