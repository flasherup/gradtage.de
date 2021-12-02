package autocompletesvc

import (
	"context"
	"github.com/flasherup/gradtage.de/autocompletesvc/acrpc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	GetAutocompleteEndpoint  		endpoint.Endpoint
	AddSourcesEndpoint  			endpoint.Endpoint
	ResetSourcesEndpoint  			endpoint.Endpoint
	GetAllStationsEndpoint  		endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetAutocompleteEndpoint:   		MakeGetAutocompleteEndpoint(s),
		AddSourcesEndpoint:   			MakeAddSourcesEndpoint(s),
		ResetSourcesEndpoint:   		MakeResetSourcesEndpoint(s),
		GetAllStationsEndpoint:   		MakeGetAllStationsEndpoint(s),
	}
}

func MakeGetAutocompleteEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAutocompleteRequest)
		result, err := s.GetAutocomplete(ctx, req.Text)
		return GetAutocompleteResponse{result, err}, err
	}
}

func MakeAddSourcesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddSourcesRequest)
		err := s.AddSources(ctx, req.Sources)
		return AddSourcesResponse{err}, err
	}
}

func MakeResetSourcesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ResetSourcesRequest)
		err := s.ResetSources(ctx, req.Sources)
		return ResetSourcesResponse{err}, err
	}
}

func MakeGetAllStationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		stations, err := s.GetAllStations(ctx)
		return acrpc.GetAllStationsResponse{Stations:stations, Err:common.ErrorToString(err)}, err
	}
}