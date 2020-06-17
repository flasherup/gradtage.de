package alertsvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct {
	SendAlertEndpoint  	endpoint.Endpoint
	SendEmailEndpoint  	endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		SendAlertEndpoint:   	MakeSendAlertEndpoint(s),
		SendEmailEndpoint:   	MakeSendEmailEndpoint(s),
	}
}

func MakeSendAlertEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendAlertRequest)
		err := s.SendAlert(ctx, req.Alert)
		return SendAlertResponse{err}, err
	}
}

func MakeSendEmailEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendEmailRequest)
		err := s.SendEmail(ctx, req.Email)
		return SendEmailResponse{err}, err
	}
}