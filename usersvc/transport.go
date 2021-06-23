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
	createOrder       gt.Handler
	updateOrder       gt.Handler
	deleteOrder       gt.Handler
	addPlan           gt.Handler
	validateSelection gt.Handler
	validateKey       gt.Handler
	validateOrder     gt.Handler
}

func (s *GRPCServer) CreateOrder(ctx context.Context, req *grpcusr.CreateOrderRequest) (*grpcusr.CreateOrderResponse, error) {
	_, resp, err := s.createOrder.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.CreateOrderResponse), err
}

func (s *GRPCServer) UpdateOrder(ctx context.Context, req *grpcusr.UpdateOrderRequest) (*grpcusr.UpdateOrderResponse, error) {
	_, resp, err := s.updateOrder.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.UpdateOrderResponse), err
}

func (s *GRPCServer) DeleteOrder(ctx context.Context, req *grpcusr.DeleteOrderRequest) (*grpcusr.DeleteOrderResponse, error) {
	_, resp, err := s.deleteOrder.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.DeleteOrderResponse), err
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

func (s *GRPCServer) ValidateOrder(ctx context.Context, req *grpcusr.ValidateOrderRequest) (*grpcusr.ValidateOrderResponse, error) {
	_, resp, err := s.validateOrder.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*grpcusr.ValidateOrderResponse), err
}


func NewGRPCServer(_ context.Context, endpoint Endpoints) grpcusr.UserSVCServer {
	server := GRPCServer{
		createOrder: gt.NewServer(
			endpoint.CreateOrderEndpoint,
			DecodeCreateOrderRequest,
			EncodeCreateOrderResponse,
		),
		updateOrder: gt.NewServer(
			endpoint.UpdateOrderEndpoint,
			DecodeUpdateOrderRequest,
			EncodeUpdateOrderResponse,
		),
		deleteOrder: gt.NewServer(
			endpoint.DeleteOrderEndpoint,
			DecodeDeleteOrderRequest,
			EncodeDeleteOrderResponse,
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
		validateOrder: gt.NewServer(
			endpoint.ValidateOrderEndpoint,
			DecodeValidateOrderRequest,
			EncodeValidateOrderResponse,
		),
	}
	return &server
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}