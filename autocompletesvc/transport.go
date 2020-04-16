package autocompletesvc

import (
	"context"
	"github.com/flasherup/gradtage.de/autocompletesvc/acrpc"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type GRPCServer struct {
	getAutocomplete    	gt.Handler
	addSources    		gt.Handler
	resetSources    	gt.Handler
}

func (s *GRPCServer) GetAutocomplete(ctx context.Context, req *acrpc.GetAutocompleteRequest) (*acrpc.GetAutocompleteResponse, error) {
	_, resp, err := s.getAutocomplete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*acrpc.GetAutocompleteResponse), nil
}

func (s *GRPCServer) AddSources(ctx context.Context, req *acrpc.AddSourcesRequest) (*acrpc.AddSourcesResponse, error) {
	_, resp, err := s.addSources.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*acrpc.AddSourcesResponse), nil
}

func (s *GRPCServer) ResetSources(ctx context.Context, req *acrpc.ResetSourcesRequest) (*acrpc.ResetSourcesResponse, error) {
	_, resp, err := s.resetSources.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*acrpc.ResetSourcesResponse), nil
}

func NewGRPCServer(_ context.Context, endpoint Endpoints) acrpc.AutocompleteSVCServer {
	return &GRPCServer{
		getAutocomplete: gt.NewServer(
			endpoint.GetAutocompleteEndpoint,
			DecodeGetAutocompleteRequest,
			EncodeGetAutocompleteResponse,
		),
		addSources: gt.NewServer(
			endpoint.AddSourcesEndpoint,
			DecodeAddSourcesRequest,
			EncodeAddSourcesResponse,
		),
		resetSources: gt.NewServer(
			endpoint.ResetSourcesEndpoint,
			DecodeResetSourcesRequest,
			EncodeResetSourcesResponse,
		),
	}
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}