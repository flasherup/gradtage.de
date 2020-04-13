package autocompletesvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	GetAutocompleteEndpoint  		endpoint.Endpoint
	AddSourcesEndpoint  			endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetAutocompleteEndpoint:   		MakeGetAutocompleteEndpoint(s),
		AddSourcesEndpoint:   			MakeAddSourcesEndpoint(s),
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