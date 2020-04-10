package stationssvc

import (
	"context"
	"github.com/flasherup/gradtage.de/autocompletesvc/acrpc"
)

func EncodeGetAutocompleteResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetAutocompleteResponse)
	return &acrpc.GetAutocompleteResponse {
		Result: res.Result,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetAutocompleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*acrpc.GetAutocompleteRequest)
	return GetAutocompleteRequest{req.Text}, nil
}

func errorToString(err error) string{
	if err == nil {
		return "nil"
	}
	return err.Error()
}


