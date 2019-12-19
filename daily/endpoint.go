package daily

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func CreateStatusEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		status, err := s.Status(ctx)
		return StatusResponse{status, err}, nil
	}
}