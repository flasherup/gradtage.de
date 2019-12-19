package monitor

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	LogEndpoint  endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		LogEndpoint:   MakeLogEndpoint(s),
	}
}

func MakeLogEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LogRequest)
		err := s.Log(ctx, req.Log)
		return LogResponse{err}, nil
	}
}