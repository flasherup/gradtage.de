package daydegreesvc

import (
	"context"
	"github.com/flasherup/gradtage.de/daydegreesvc/ddgrpc"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type GRPCServer struct {
	getDegree      		gt.Handler
	getAverageDegree 	gt.Handler

}

func (s *GRPCServer) GetDegree(ctx context.Context, req *ddgrpc.GetDegreeRequest) (*ddgrpc.GetDegreeResponse, error) {
	_, resp, err := s.getDegree.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ddgrpc.GetDegreeResponse), err
}

func (s *GRPCServer) GetAverageDegree(ctx context.Context, req *ddgrpc.GetAverageDegreeRequest) (*ddgrpc.GetAverageDegreeResponse, error) {
	_, resp, err := s.getAverageDegree.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ddgrpc.GetAverageDegreeResponse), err
}



func NewGRPCServer(_ context.Context, endpoint Endpoints) *GRPCServer {
	server := GRPCServer{
		getDegree: gt.NewServer(
			endpoint.GetDegreeEndpoint,
			DecodeGetDegreeRequest,
			EncodeGetDegreeResponse,
		),
		getAverageDegree: gt.NewServer(
			endpoint.GetAverageDegreeEndpoint,
			DecodeGetAverageDegreeRequest,
			EncodeGetAverageDegreeResponse,
		),
	}
	return &server
}


func NewMetricsTransport(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}
