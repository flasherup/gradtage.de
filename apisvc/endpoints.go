package apisvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	GetHDDEndpoint  		endpoint.Endpoint
	GetHDDSVEndpoint  		endpoint.Endpoint
	GetCDDSVEndpoint  		endpoint.Endpoint
	GetSourceDataEndpoint	endpoint.Endpoint
	SearchEndpoint			endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetHDDEndpoint:   		MakeGetHDDEndpoint(s),
		GetHDDSVEndpoint: 		MakeGetHDDCSVEndpoint(s),
		GetCDDSVEndpoint: 		MakeGetCDDCSVEndpoint(s),
		GetSourceDataEndpoint:  MakeGetSourceDataEndpoint(s),
		SearchEndpoint:  		MakeSearchEndpoint(s),
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

func MakeGetCDDCSVEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCDDCSVRequest)
		data, filename, err := s.GetCDDCSV(ctx, req.Params)
		return GetCDDCSVResponse{ data, filename }, err
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