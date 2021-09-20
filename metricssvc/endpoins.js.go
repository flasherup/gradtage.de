package metricssvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetMetricsEndpoint      endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetMetricsEndpoint:       MakeGetMetricsEndpoint(s),
	}
}

func MakeGetMetricsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetMetricsRequest)
		metrics, err := s.GetMetrics(ctx, req.Ids)
		return GetMetricsResponse{metrics, err}, err
	}
}
