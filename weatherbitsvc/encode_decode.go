package weatherbitsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/hourlysvc"
	weathergrpc "github.com/flasherup/gradtage.de/weatherbitsvc/weatherbitgrpc"
)

func EncodeGetPeriodResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetPeriodResponse)
	encTemp := toGRPCTemps(res.Temps)
	return &weathergrpc.GetPeriodResponse {
		Temps: encTemp,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetPeriodRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*weathergrpc.GetPeriodRequest)
	return GetPeriodRequest{req.Ids, req.Start, req.End}, nil
}

func toGRPCTemps(src map[string][]hourlysvc.Temperature)  map[string]*weathergrpc.Temperatures {
	res := make(map[string]*weathergrpc.Temperatures)
	for k,v := range src {
		temps := make([]*weathergrpc.Temperature, len(v))
		for i,t := range v {
			temps[i] = &weathergrpc.Temperature{
				Date: 			t.Date,
				Temperature: 	t.Temperature,
			}
		}
		res[k] = &weathergrpc.Temperatures {
			Temps: temps,
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