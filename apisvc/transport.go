package apisvc

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

const Method = "method"
const UserAction = "userAction"
const PlanAction = "planAction"

func NewHTTPTSransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	e := MakeServerEndpoints(s)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	r.Methods("POST").Path("/temperature").Handler(kithttp.NewServer(
		e.GetHDDEndpoint,
		decodeGetHDDRequest,
		encodeGetHDDResponse,
		options...,
	))


	r.Methods("Get").Path("/temperature/csv/{" + Method + "}").Handler(kithttp.NewServer(
		e.GetHDDSVEndpoint,
		decodeGetHDDCSVRequest,
		encodeGetHDDCSVResponse,
		options...,
	))

	r.Methods("Get").Path("/source/").Handler(kithttp.NewServer(
		e.GetSourceDataEndpoint,
		decodeGetSourceDataRequest,
		encodeGetSourceDataResponse,
		options...,
	))

	r.Methods("Get").Path("/search/").Handler(kithttp.NewServer(
		e.SearchEndpoint,
		decodeSearchRequest,
		encodeSearchResponse,
		options...,
	))

	r.Methods("Post").Path("/user/{" + UserAction + "}").Handler(kithttp.NewServer(
		e.UserEndpoint,
		decodeUserRequest,
		encodeUserResponse,
		options...,
	))

	r.Methods("Post").Path("/plan/{" + PlanAction + "}").Handler(kithttp.NewServer(
		e.PlanEndpoint,
		decodePlanRequest,
		encodePlanResponse,
		options...,
	))
	return r
}

func NewHTTPTransport(s Service, logger log.Logger,) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}


func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}