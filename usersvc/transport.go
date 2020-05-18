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
	createUser    		gt.Handler
	updateUser			gt.Handler
	addPlan   			gt.Handler
	validateSelection	gt.Handler
	validateKey			gt.Handler
	validateName		gt.Handler
}

func (s *GRPCServer) CreateUser(ctx context.Context, req *grpcusr.CreateUserRequest) (*grpcusr.CreateUserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.CreateUserResponse), err
}

func (s *GRPCServer) UpdateUser(ctx context.Context, req *grpcusr.UpdateUserRequest) (*grpcusr.UpdateUserResponse, error) {
	_, resp, err := s.updateUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.UpdateUserResponse), err
}

func (s *GRPCServer) AddPlan(ctx context.Context, req *grpcusr.AddPlanRequest) (*grpcusr.AddPlanResponse, error) {
	_, resp, err := s.addPlan.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.AddPlanResponse), err
}

func (s *GRPCServer) ValidateSelection(ctx context.Context, req *grpcusr.ValidateSelectionRequest) (*grpcusr.ValidateSelectionResponse, error) {
	_, resp, err := s.validateSelection.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.ValidateSelectionResponse), err
}

func (s *GRPCServer) ValidateKey(ctx context.Context, req *grpcusr.ValidateKeyRequest) (*grpcusr.ValidateKeyResponse, error) {
	_, resp, err := s.validateKey.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.ValidateKeyResponse), err
}

func (s *GRPCServer) ValidateName(ctx context.Context, req *grpcusr.ValidateNameRequest) (*grpcusr.ValidateNameResponse, error) {
	_, resp, err := s.validateName.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.ValidateNameResponse), err
}




func NewGRPCServer(_ context.Context, endpoint Endpoints) grpcusr.UserSVCServer {
	server := GRPCServer{
		createUser: gt.NewServer(
			endpoint.CreateUserEndpoint,
			DecodeCreateUserRequest,
			EncodeCreateUserResponse,
		),
		updateUser: gt.NewServer(
			endpoint.UpdateUserEndpoint,
			DecodeUpdateUserRequest,
			EncodeUpdateUserResponse,
		),
		addPlan: gt.NewServer(
			endpoint.AddPlanEndpoint,
			DecodeAddPlanRequest,
			EncodeAddPlanResponse,
		),
		validateSelection: gt.NewServer(
			endpoint.ValidateSelectionEndpoint,
			DecodeValidateSelectionRequest,
			EncodeValidateSelectionResponse,
		),
		validateKey: gt.NewServer(
			endpoint.ValidateKeyEndpoint,
			DecodeValidateKeyRequest,
			EncodeValidateKeyResponse,
		),
		validateName: gt.NewServer(
			endpoint.ValidateNameEndpoint,
			DecodeValidateNameRequest,
			EncodeValidateNameResponse,
		),
	}
	return &server
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}