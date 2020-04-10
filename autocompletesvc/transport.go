package stationssvc

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
}

func (s *GRPCServer) GetAutocomplete(ctx context.Context, req *acrpc.GetAutocompleteRequest) (*acrpc.GetAutocompleteResponse, error) {
	_, resp, err := s.getAutocomplete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*acrpc.GetAutocompleteResponse), nil
}

func NewGRPCServer(_ context.Context, endpoint Endpoints) acrpc.AutocompleteSVCServer {
	return &GRPCServer{
		getAutocomplete: gt.NewServer(
			endpoint.GetAutocompleteEndpoint,
			DecodeGetAutocompleteRequest,
			EncodeGetAutocompleteResponse,
		),
	}
}

func NewMetricsTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}