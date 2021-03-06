package dailysvc

import (
	"context"
	"github.com/flasherup/gradtage.de/dailysvc/dlygrpc"
)


func EncodeGetPeriodResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetPeriodResponse)
	encTemp := EncodeTemperature(res.Temps)
	return &dlygrpc.GetPeriodResponse {
		Temps: encTemp,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetPeriodRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*dlygrpc.GetPeriodRequest)
	return GetPeriodRequest{req.Id, req.Start, req.End}, nil
}

func EncodePushPeriodResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(PushPeriodResponse)
	return &dlygrpc.PushPeriodResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodePushPeriodRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*dlygrpc.PushPeriodRequest)
	encTemp := DecodeTemperature(req.Temps)
	return PushPeriodRequest{req.Id, encTemp}, nil
}

func EncodeGetUpdateDateResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetUpdateDateResponse)
	return &dlygrpc.GetUpdateDateResponse {
		Dates: res.Dates,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetUpdateDateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*dlygrpc.GetUpdateDateRequest)
	return GetUpdateDateRequest{req.Ids}, nil
}


func EncodeUpdateAvgForYearResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(UpdateAvgForYearResponse)
	return &dlygrpc.UpdateAvgForYearResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeUpdateAvgForYearRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*dlygrpc.UpdateAvgForYearRequest)
	return UpdateAvgForYearRequest{req.Id }, nil
}


func EncodeUpdateAvgForDOYResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(UpdateAvgForDOYResponse)
	return &dlygrpc.UpdateAvgForDOYResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeUpdateAvgForDOYRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*dlygrpc.UpdateAvgForDOYRequest)
	return UpdateAvgForDOYRequest{req.Id, int(req.Doy) }, nil
}

func EncodeGetAvgResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetAvgResponse)
	encTemp := toGRPCMapTemps(res.Temps)
	return &dlygrpc.GetAvgResponse {
		Temps: encTemp,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetAvgRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*dlygrpc.GetAvgRequest)
	return GetAvgRequest{req.Id}, nil
}

func EncodeTemperature(src []Temperature) []*dlygrpc.Temperature {
	res := make([]*dlygrpc.Temperature, len(src))
	for i,v := range src {
		res[i] = &dlygrpc.Temperature {
			Date: 			v.Date,
			Temperature: 	v.Temperature,
		}
	}
	return res
}

func DecodeTemperature(src []*dlygrpc.Temperature) []Temperature {
	res := make([]Temperature , len(src))
	for i,v := range src {
		res[i] =  Temperature{
			Date:			v.Date,
			Temperature:	v.Temperature,
		}
	}
	return res
}

func toGRPCMapTemps(src map[int]Temperature) map[int32]*dlygrpc.Temperature {
	res := make(map[int32]*dlygrpc.Temperature)
	for k,v := range src {
		res[int32(k)] = &dlygrpc.Temperature{
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


