package alertsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc/grpcalt"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type GRPCServer struct {
	sendAlert    	gt.Handler
	sendEmail    	gt.Handler
}

func (s *GRPCServer) SendAlert(ctx context.Context, req *grpcalt.SendAlertRequest) (*grpcalt.SendAlertResponse, error) {
	_, resp, err := s.sendAlert.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcalt.SendAlertResponse), err
}

func (s *GRPCServer) SendEmail(ctx context.Context, req *grpcalt.SendEmailRequest) (*grpcalt.SendEmailResponse, error) {
	_, resp, err := s.sendEmail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcalt.SendEmailResponse), err
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) grpcalt.AlertSVCServer {
	return &GRPCServer{
		sendAlert: gt.NewServer(
			endpoint.SendAlertEndpoint,
			DecodeSendAlertRequest,
			EncodeSendAlertResponse,
		),
		sendEmail: gt.NewServer(
			endpoint.SendEmailEndpoint,
			DecodeSendEmailRequest,
			EncodeSendEmailResponse,
		),
	}
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}