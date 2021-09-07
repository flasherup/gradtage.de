package weatherbitsvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ForceRestartEndpoint      endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		ForceRestartEndpoint:       MakeForceRestartEndpoint(s),
	}
}

func MakeForceRestartEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := s.ForceRestart(ctx)
		return ForceRestartResponse{err}, err
	}
}
