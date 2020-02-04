package noaascrapersvc

import (
	"context"
	"github.com/flasherup/gradtage.de/noaascrapersvc/noaascpc"
)


func EncodeForceOverrideHourlyResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(ForceOverrideHourlyResponse)
	return &noaascpc.ForceOverrideHourlyResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeForceOverrideHourlyRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*noaascpc.ForceOverrideHourlyRequest)
	return ForceOverrideHourlyRequest{
		req.Station,
		req.Start,
		req.End,
	}, nil
}


func errorToString(err error) string{
	if err == nil {
		return "nil"
	}

	return err.Error()
}


