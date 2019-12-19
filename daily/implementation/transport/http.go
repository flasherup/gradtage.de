package transport

import (
	"context"
	"encoding/json"
	"github.com/flasherup/gradtage.de/daily"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHTTPTransport(s daily.Service, logger log.Logger,) http.Handler {
	var (
		r = mux.NewRouter()
	)

	r.Methods("GET").Path("/status").Handler(kithttp.NewServer(
		daily.CreateStatusEndpoint(s),
		decodeStatusRequest,
		encodeStatusResponse,
	))

	return r
}

func decodeStatusRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req daily.StatusRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}


func encodeStatusResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}