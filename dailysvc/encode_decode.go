package dailysvc

import (
	"context"
	"github.com/flasherup/gradtage.de/dailysvc/grpc"
)


func EncodeGetPeriodResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetPeriodResponse)
	encTemp := toGRPCTemps(res.Temps)
	return &grpc.GetPeriodResponse {
		Temps: encTemp,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetPeriodRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpc.GetPeriodRequest)
	return GetPeriodRequest{req.Id, req.Start, req.End}, nil
}

func EncodePushPeriodResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(PushPeriodResponse)
	return &grpc.PushPeriodResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodePushPeriodRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpc.PushPeriodRequest)
	encTemp := toServiceTemps(req.Temps)
	return PushPeriodRequest{req.Id, encTemp}, nil
}

func EncodeGetUpdateDateResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetUpdateDateResponse)
	return &grpc.GetUpdateDateResponse {
		Dates: res.Dates,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetUpdateDateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpc.GetUpdateDateRequest)
	return GetUpdateDateRequest{req.Ids}, nil
}


func EncodeUpdateAvgForYearResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(UpdateAvgForYearResponse)
	return &grpc.UpdateAvgForYearResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeUpdateAvgForYearRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpc.UpdateAvgForYearRequest)
	return UpdateAvgForYearRequest{req.Id }, nil
}


func EncodeUpdateAvgForDOYResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(UpdateAvgForDOYResponse)
	return &grpc.UpdateAvgForDOYResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeUpdateAvgForDOYRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpc.UpdateAvgForDOYRequest)
	return UpdateAvgForDOYRequest{req.Id, int(req.Doy) }, nil
}

func EncodeGetAvgResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetAvgResponse)
	encTemp := toGRPCTemps(res.Temps)
	return &grpc.GetAvgResponse {
		Temps: encTemp,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetAvgRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpc.GetAvgRequest)
	return GetAvgRequest{req.Id}, nil
}



func toGRPCTemps(src []Temperature) []*grpc.Temperature {
	res := make([]*grpc.Temperature, len(src))
	for i,v := range src {
		res[i] = &grpc.Temperature {
			Date: 			v.Date,
			Temperature: 	float32(v.Temperature),
		}
	}
	return res
}

func toServiceTemps(src []*grpc.Temperature) []Temperature {
	res := make([]Temperature , len(src))
	for i,v := range src {
		res[i] =  Temperature{
			Date:			v.Date,
			Temperature:	float64(v.Temperature),
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


