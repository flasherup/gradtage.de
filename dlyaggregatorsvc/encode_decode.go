package dlyaggregatorsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/dagrpc"
)


func EncodeForceUpdateResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(ForceUpdateResponse)
	return &dagrpc.ForceUpdateResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeForceUpdateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*dagrpc.ForceUpdateRequest)
	return ForceUpdateRequest{req.Ids, req.Start, req.End}, nil
}

func errorToString(err error) string{
	if err == nil {
		return "nil"
	}

	return err.Error()
}


