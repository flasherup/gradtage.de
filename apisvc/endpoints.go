package apisvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeGetDataEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDataRequest)
		data, format, err := s.GetData(ctx, req.Params)
		return GetDataResponse{data, format, err}, err
	}
}

func MakeGetSourceDataEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSourceDataRequest)
		data, filename, err := s.GetSourceData(ctx, req.Params)
		return GetSourceDataResponse{data, filename}, err
	}
}

func MakeSearchEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchRequest)
		data, err := s.Search(ctx, req.Params)
		return SearchResponse{data}, err
	}
}

func MakeUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserRequest)
		data, err := s.User(ctx, req.Params)
		return UserResponse{data}, err
	}
}

func MakeWoocommerceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WoocommerceRequest)
		data, err := s.Woocommerce(ctx, req.Event)
		return WoocommerceResponse{data}, err
	}
}

func MakeServiceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ServiceRequest)
		data, err := s.Service(ctx, req.Name, req.Params)
		return ServiceResponse{data}, err
	}
}

func MakeOrderCreateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OrderCreateRequest)
		err := s.OrderCreate(ctx, req.OrderID, req.Email, req.Plan, req.Key)
		if err == nil {
			return OrderCreateResponse{Status: "success"}, err
		}

		return OrderCreateResponse{
			Status: "error",
			Error:  err,
		}, err
	}
}

func MakeOrderUpdateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OrderUpdateRequest)
		err := s.OrderUpdate(ctx, req.OrderID, req.Email, req.Plan, req.Key)
		if err == nil {
			return OrderUpdateResponse{Status: "success"}, err
		}

		return OrderUpdateResponse{
			Status: "error",
			Error:  err,
		}, err
	}
}

func MakeOrderDeleteEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OrderDeleteRequest)
		err := s.OrderDelete(ctx, req.OrderID)
		if err == nil {
			return OrderDeleteResponse{Status: "success"}, err
		}

		return OrderDeleteResponse{
			Status: "error",
			Error:  err,
		}, err
	}
}

func MakeOrderCancelEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OrderCancelRequest)
		err := s.OrderCancel(ctx, req.OrderID)
		if err == nil {
			return OrderCancelResponse{Status: "success"}, err
		}

		return OrderCancelResponse{
			Status: "error",
			Error:  err,
		}, err
	}
}
