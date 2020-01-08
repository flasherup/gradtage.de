package alertsvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	SendAlertEndpoint  	endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		SendAlertEndpoint:   	MakeSendAlertEndpoint(s),
	}
}

func MakeSendAlertEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendAlertRequest)
		err := s.SendAlert(ctx, req.Alert)
		return SendAlertResponse{err}, err
	}
}