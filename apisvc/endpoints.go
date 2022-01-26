package apisvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	GetHDDEndpoint  		endpoint.Endpoint
	GetHDDSVEndpoint  		endpoint.Endpoint
	GetZIPEndpoint  		endpoint.Endpoint
	GetSourceDataEndpoint	endpoint.Endpoint
	SearchEndpoint			endpoint.Endpoint
	UserEndpoint			endpoint.Endpoint
	WoocommerceEndpoint		endpoint.Endpoint
	ServiceEndpoint			endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetHDDEndpoint:   		MakeGetHDDEndpoint(s),
		GetHDDSVEndpoint: 		MakeGetHDDCSVEndpoint(s),
		GetZIPEndpoint: 		MakeGetZIPEndpoint(s),
		GetSourceDataEndpoint:  MakeGetSourceDataEndpoint(s),
		SearchEndpoint:  		MakeSearchEndpoint(s),
		UserEndpoint:  			MakeUserEndpoint(s),
		WoocommerceEndpoint:  	MakeWoocommerceEndpoint(s),
		ServiceEndpoint:  		MakeServiceEndpoint(s),
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

func MakeGetZIPEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetZIPRequest)
		data, filename, err := s.GetZIP(ctx, req.Params)
		return GetZIPResponse{ data, filename }, err
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
		req := request.(UserRequest)
		data, err := s.User(ctx, req.Params)
		return UserResponse{ data}, err
	}
}

func MakeWoocommerceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WoocommerceRequest)
		data, err := s.Woocommerce(ctx, req.Event)
		return WoocommerceResponse{ data}, err
	}
}

func MakeServiceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ServiceRequest)
		data, err := s.Service(ctx, req.Name, req.Params)
		return ServiceResponse{ data}, err
	}
}