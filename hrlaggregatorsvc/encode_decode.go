package hrlaggregatorsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/hagrpc"
)


func EncodeGetStatusResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetStatusResponse)
	encStatus := toGRPCStatus(res.Status)
	return &hagrpc.GetStatusResponse {
		Status: encStatus,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetStatusRequest(_ context.Context, r interface{}) (interface{}, error) {
	return GetStatusRequest{}, nil
}


func toGRPCStatus(src []Status) []*hagrpc.Status {
	res := make([]*hagrpc.Status, len(src))
	for i,v := range src {
		res[i] = &hagrpc.Status {
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


