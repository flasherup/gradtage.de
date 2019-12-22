package testsvc

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func NewHTTPTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	e := MakeServerEndpoints(s)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	r.Methods("POST").Path("/").Handler(kithttp.NewServer(
		e.LogEndpoint,
		decodeTestRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeTestRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req TestRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}


func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}