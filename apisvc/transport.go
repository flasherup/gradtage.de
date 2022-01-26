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
const DayCalc = "day_calc"
const UserAction = "userAction"
const ServiceName = "serviceType"

func NewHTTPTSransport(s Service, logger log.Logger, staticFolder string) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	e := MakeServerEndpoints(s)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	r.Methods("POST").Path("/degreedays/{" + Method + "}/{" + DayCalc + "}/").Handler(kithttp.NewServer(
		e.GetHDDEndpoint,
		decodeGetHDDRequest,
		encodeGetHDDResponse,
		options...,
	))

	r.Methods("GET").Path("/degreedays/{" + Method + "}/{" + DayCalc + "}/zip/").Handler(kithttp.NewServer(
		e.GetZIPEndpoint,
		decodeGetZIPRequest,
		encodeGetZIPResponse,
		options...,
	))


	r.Methods("GET").Path("/degreedays/{" + Method + "}/{" + DayCalc + "}/").Handler(kithttp.NewServer(
		e.GetHDDSVEndpoint,
		decodeGetHDDCSVRequest,
		encodeGetHDDCSVResponse,
		options...,
	))

	r.Methods("GET").Path("/source/").Handler(kithttp.NewServer(
		e.GetSourceDataEndpoint,
		decodeGetSourceDataRequest,
		encodeGetSourceDataResponse,
		options...,
	))

	r.Methods("GET").Path("/search/").Handler(kithttp.NewServer(
		e.SearchEndpoint,
		decodeSearchRequest,
		encodeSearchResponse,
		options...,
	))

	r.Methods("GET").Path("/user/{" + UserAction + "}").Handler(kithttp.NewServer(
		e.UserEndpoint,
		decodeUserRequest,
		encodeUserResponse,
		options...,
	))

	r.Methods("POST").Path("/woocommerce").Handler(kithttp.NewServer(
		e.WoocommerceEndpoint,
		decodeWoocommerceRequest,
		encodeWoocommerceResponse,
		options...,
	))

	r.Methods("GET").Path("/service/{" + ServiceName + "}/").Handler(kithttp.NewServer(
		e.ServiceEndpoint,
		decodeServiceRequest,
		encodeServiceResponse,
		options...,
	))

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticFolder)))

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