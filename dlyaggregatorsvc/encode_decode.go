package dlyaggregatorsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/dagrpc"
)


func EncodeGetStatusResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetStatusResponse)
	encStatus := toGRPCStatus(res.Status)
	return &dagrpc.GetStatusResponse {
		Status: encStatus,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetStatusRequest(_ context.Context, r interface{}) (interface{}, error) {
	return GetStatusRequest{}, nil
}


func toGRPCStatus(src []Status) []*dagrpc.Status {
	res := make([]*dagrpc.Status, len(src))
	for i,v := range src {
		res[i] = &dagrpc.Status {
			Station: 		v.Station,
			Update: 		v.Update,
			Temperature: 	v.Temperature,
		}
	}
	return res
}

func errorToString(err error) string{
	if err == nil {
		return "nil"
	}

	return err.Error()
}


