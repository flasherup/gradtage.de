package usersvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateOrderEndpoint  		endpoint.Endpoint
	UpdateOrderEndpoint  		endpoint.Endpoint
	DeleteOrderEndpoint  		endpoint.Endpoint
	AddPlanEndpoint  			endpoint.Endpoint
	ValidateSelectionEndpoint  	endpoint.Endpoint
	ValidateKeyEndpoint  		endpoint.Endpoint
	ValidateOrderEndpoint  		endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateOrderEndpoint:   		MakeCreateOrderEndpoint(s),
		UpdateOrderEndpoint:   		MakeUpdateOrderEndpoint(s),
		DeleteOrderEndpoint:   		MakeDeleteOrderEndpoint(s),
		AddPlanEndpoint:   			MakeAddPlanEndpoint(s),
		ValidateSelectionEndpoint:  MakeValidateSelectionEndpoint(s),
		ValidateKeyEndpoint:   		MakeValidateKeyEndpoint(s),
		ValidateOrderEndpoint:   	MakeValidateOrderEndpoint(s),
	}
}


func MakeCreateOrderEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateOrderRequest)
		key, err := s.CreateOrder(ctx, req.OrderId, req.Email,req.Plan, req.Key,)
		return CreateOrderResponse{key, err}, err
	}
}

func MakeUpdateOrderEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateOrderRequest)
		key, err := s.UpdateOrder(ctx, req.Order)
		return UpdateOrderResponse{key, err}, err
	}
}

func MakeDeleteOrderEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteOrderRequest)
		err := s.DeleteOrder(ctx, req.OrderId)
		return DeleteOrderResponse{err}, err
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
		err := s.ValidateSelection(ctx, req.Selection)
		return ValidateSelectionResponse{ err}, err
	}
}

func MakeValidateKeyEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateKeyRequest)
		order, plan, err := s.ValidateKey(ctx, req.Key)
		return ValidateKeyResponse{ order, plan, err}, err
	}
}

func MakeValidateOrderEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateOrderRequest)
		order, plan, err := s.ValidateOrder(ctx, req.OrderId)
		return ValidateOrderResponse{ order, plan, err}, err
	}
}