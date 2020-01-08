package alertsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc/altgrpc"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type GRPCServer struct {
	sendAlert    	gt.Handler
}

func (s *GRPCServer) SendAlert(ctx context.Context, req *altgrpc.SendAlertRequest) (*altgrpc.SendAlertResponse, error) {
	_, resp, err := s.sendAlert.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*altgrpc.SendAlertResponse), err
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) altgrpc.AlertSVCServer {
	return &GRPCServer{
		sendAlert: gt.NewServer(
			endpoint.SendAlertEndpoint,
			DecodeSendAlertRequest,
			EncodeSendAlertResponse,
		),
	}
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}