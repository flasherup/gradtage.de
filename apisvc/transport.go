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
const OrderID = "orderID"

func NewHandler(s Service, logger log.Logger, staticFolder string) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	r.Methods("POST").
		Path("/degreedays/{" + Method + "}/{" + DayCalc + "}/").
		Handler(kithttp.NewServer(
			MakeGetDataEndpoint(s),
			decodeGetHDDRequest,
			encodeGetHDDResponse,
			options...,
		))

	r.Methods("GET").
		Path("/degreedays/{" + Method + "}/{" + DayCalc + "}/").
		Handler(kithttp.NewServer(
			MakeGetDataEndpoint(s),
			decodeGetDataRequest,
			encodeGetDataResponse,
			options...,
		))

	r.Methods("GET").
		Path("/source/").
		Handler(kithttp.NewServer(
			MakeGetSourceDataEndpoint(s),
			decodeGetSourceDataRequest,
			encodeGetSourceDataResponse,
			options...,
		))

	r.Methods("GET").
		Path("/search/").
		Handler(kithttp.NewServer(
			MakeGetDataEndpoint(s),
			decodeSearchRequest,
			encodeSearchResponse,
			options...,
		))

	r.Methods("GET").
		Path("/user/{" + UserAction + "}").
		Handler(kithttp.NewServer(
			MakeUserEndpoint(s),
			decodeUserRequest,
			encodeUserResponse,
			options...,
		))

	r.Methods("POST").
		Path("/woocommerce").
		Handler(kithttp.NewServer(
			MakeWoocommerceEndpoint(s),
			decodeWoocommerceRequest,
			encodeWoocommerceResponse,
			options...,
		))

	r.Methods("GET").
		Path("/service/{" + ServiceName + "}/").
		Handler(kithttp.NewServer(
			MakeServiceEndpoint(s),
			decodeServiceRequest,
			encodeServiceResponse,
			options...,
		))

	r.Methods("POST").
		Path("/orders").
		Handler(kithttp.NewServer(
			MakeOrderCreateEndpoint(s),
			decodeOrderCreateRequest,
			encodeOrderCreateResponse,
			options...,
		))

	r.Methods("PUT").
		Path("/orders/{" + OrderID + "}").
		Handler(kithttp.NewServer(
			MakeOrderUpdateEndpoint(s),
			decodeOrderUpdateRequest,
			encodeOrderUpdateResponse,
			options...,
		))

	r.Methods("DELETE").
		Path("/orders/{" + OrderID + "}").
		Handler(kithttp.NewServer(
			MakeOrderDeleteEndpoint(s),
			decodeOrderDeleteRequest,
			encodeOrderDeleteResponse,
			options...,
		))

	r.Methods("POST").
		Path("/orders/{" + OrderID + "}/cancel").
		Handler(kithttp.NewServer(
			MakeOrderCancelEndpoint(s),
			decodeOrderCancelRequest,
			encodeOrderCancelResponse,
			options...,
		))

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticFolder)))

	return r
}

func NewHTTPTransport(s Service, logger log.Logger) http.Handler {
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
