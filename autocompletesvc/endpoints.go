package stationssvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	GetAutocompleteEndpoint  			endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetAutocompleteEndpoint:   			MakeGetAutocompleteEndpoint(s),
	}
}


func MakeGetAutocompleteEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAutocompleteRequest)
		result, id, err := s.GetAutocomplete(ctx, req.Text)
		return GetAutocompleteResponse{result, id, err}, err
	}
}