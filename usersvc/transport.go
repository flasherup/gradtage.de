package usersvc

import (
	"context"
	"github.com/flasherup/gradtage.de/usersvc/grpcusr"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type GRPCServer struct {
	createUser    	gt.Handler
	createUserAuto  gt.Handler
	setPlan   		gt.Handler
	setStations		gt.Handler
	validateKey		gt.Handler
}

func (s *GRPCServer) CreateUser(ctx context.Context, req *grpcusr.CreateUserRequest) (*grpcusr.CreateUserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.CreateUserResponse), err
}

func (s *GRPCServer) CreateUserAuto(ctx context.Context, req *grpcusr.CreateUserAutoRequest) (*grpcusr.CreateUserAutoResponse, error) {
	_, resp, err := s.createUserAuto.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.CreateUserAutoResponse), err
}

func (s *GRPCServer) SetPlan(ctx context.Context, req *grpcusr.SetPlanRequest) (*grpcusr.SetPlanResponse, error) {
	_, resp, err := s.setPlan.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.SetPlanResponse), err
}

func (s *GRPCServer) SetStations(ctx context.Context, req *grpcusr.SetStationsRequest) (*grpcusr.SetStationsResponse, error) {
	_, resp, err := s.setStations.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.SetStationsResponse), err
}

func (s *GRPCServer) ValidateKey(ctx context.Context, req *grpcusr.ValidateKeyRequest) (*grpcusr.ValidateKeyResponse, error) {
	_, resp, err := s.validateKey.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.ValidateKeyResponse), err
}




func NewGRPCServer(_ context.Context, endpoint Endpoints) grpcusr.UserSVCServer {
	server := GRPCServer{
		createUser: gt.NewServer(
			endpoint.CreateUserEndpoint,
			DecodeCreateUserRequest,
			EncodeCreateUserResponse,
		),
		createUserAuto: gt.NewServer(
			endpoint.CreateUserAutoEndpoint,
			DecodeCreateUserAutoRequest,
			EncodeCreateUserAutoResponse,
		),
		setPlan: gt.NewServer(
			endpoint.SetPlanEndpoint,
			DecodeSetPlanRequest,
			EncodeSetPlanResponse,
		),
		setStations: gt.NewServer(
			endpoint.SetStationsEndpoint,
			DecodeSetStationsRequest,
			EncodeSetStationsResponse,
		),
		validateKey: gt.NewServer(
			endpoint.ValidateKeyEndpoint,
			DecodeValidateKeyRequest,
			EncodeValidateKeyResponse,
		),
	}
	return &server
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}