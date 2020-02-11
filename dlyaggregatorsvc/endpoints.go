package dlyaggregatorsvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ForceUpdateEndpoint  	endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		ForceUpdateEndpoint:   	MakeForceUpdateEndpoint(s),
	}
}

func MakeForceUpdateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ForceUpdateRequest)
		err := s.ForceUpdate(ctx, req.IDs, req.Start, req.End)
		return ForceUpdateResponse{ err }, err
	}
}