package dlyaggregatorsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/grpc"
)


func EncodeGetStatusResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetStatusResponse)
	encStatus := toGRPCStatus(res.Status)
	return &grpc.GetStatusResponse {
		Status: encStatus,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetStatusRequest(_ context.Context, r interface{}) (interface{}, error) {
	return GetStatusRequest{}, nil
}


func toGRPCStatus(src []Status) []*grpc.Status {
	res := make([]*grpc.Status, len(src))
	for i,v := range src {
		res[i] = &grpc.Status {
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


