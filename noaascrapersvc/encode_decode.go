package noaascrapersvc

import (
	"context"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/noaascrapersvc/noaascpc"
)


func EncodeGetPeriodResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetPeriodResponse)
	encTemp := toGRPCTemps(res.Temps)
	return &noaascpc.GetPeriodResponse {
		Temps: encTemp,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetPeriodRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*noaascpc.GetPeriodRequest)
	return GetPeriodRequest{req.Id, req.Start, req.End}, nil
}

func EncodeGetUpdateDateResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetUpdateDateResponse)
	return &noaascpc.GetUpdateDateResponse {
		Dates: res.Dates,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetUpdateDateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*noaascpc.GetUpdateDateRequest)
	return GetUpdateDateRequest{req.Ids}, nil
}


func toGRPCTemps(src []hourlysvc.Temperature) []*noaascpc.Temperature {
	res := make([]*noaascpc.Temperature, len(src))
	for i,v := range src {
		res[i] = &noaascpc.Temperature {
			Date: 			v.Date,
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


