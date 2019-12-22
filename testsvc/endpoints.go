package testsvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	LogEndpoint  endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		LogEndpoint:   MakeTestEndpoint(s),
	}
}

func MakeTestEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TestRequest)
		text,count := s.Text(ctx, req.Text)
		return TestResponse{text, count}, nil
	}
}