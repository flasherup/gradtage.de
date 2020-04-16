package hourlysvc

import (
	"context"
	"github.com/flasherup/gradtage.de/hourlysvc/hrlgrpc"
)


func EncodeGetPeriodResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetPeriodResponse)
	encTemp := EncodeTemperature(res.Temps)
	return &hrlgrpc.GetPeriodResponse {
		Temps: encTemp,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetPeriodRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*hrlgrpc.GetPeriodRequest)
	return GetPeriodRequest{req.Id, req.Start, req.End}, nil
}

func EncodePushPeriodResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(PushPeriodResponse)
	return &hrlgrpc.PushPeriodResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodePushPeriodRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*hrlgrpc.PushPeriodRequest)
	encTemp := DecodeTemperature(req.Temps)
	return PushPeriodRequest{req.Id, encTemp}, nil
}

func EncodeGetUpdateDateResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetUpdateDateResponse)
	return &hrlgrpc.GetUpdateDateResponse {
		Dates: res.Dates,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetUpdateDateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*hrlgrpc.GetUpdateDateRequest)
	return GetUpdateDateRequest{req.Ids}, nil
}



func EncodeGetLatestResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetLatestResponse)
	return &hrlgrpc.GetLatestResponse {
		Temps: toGRPCMapTemps(res.Temps),
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetLatestRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*hrlgrpc.GetLatestRequest)
	return GetLatestRequest{req.Ids}, nil
}



func EncodeTemperature(src []Temperature) []*hrlgrpc.Temperature {
	res := make([]*hrlgrpc.Temperature, len(src))
	for i,v := range src {
		res[i] = &hrlgrpc.Temperature {
			Date: 			v.Date,
			Temperature: 	v.Temperature,
		}
	}
	return res
}

func DecodeTemperature(src []*hrlgrpc.Temperature) []Temperature {
	res := make([]Temperature , len(src))
	for i,v := range src {
		res[i] =  Temperature{
			Date:			v.Date,
			Temperature:	v.Temperature,
		}
	}
	return res
}

func toGRPCMapTemps(src map[string]Temperature) map[string]*hrlgrpc.Temperature {
	res := make(map[string]*hrlgrpc.Temperature)
	for k,v := range src {
		res[k] =  &hrlgrpc.Temperature{
			Date:			v.Date,
			Temperature:	v.Temperature,
		}
	}
	return res
}

func toServiceMapTemps(src map[string]*hrlgrpc.Temperature) map[string]Temperature {
	res := make(map[string]Temperature)
	for k,v := range src {
		res[k] =  Temperature{
			Date:			v.Date,
			Temperature:	v.Temperature,
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


