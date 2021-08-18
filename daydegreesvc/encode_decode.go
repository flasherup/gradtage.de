package daydegreesvc

import (
	"context"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc/ddgrpc"
)

func DecodeGetDegreeRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*ddgrpc.GetDegreeRequest)
	params := ToParams(req.Params)
	return GetDegreeRequest{*params}, nil
}

func EncodeGetDegreeResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetDegreeResponse)
	degrees := toGRPCDegree(&res.Degrees)
	return &ddgrpc.GetDegreeResponse {
		Degrees: *degrees,
		Err: common.ErrorToString(res.Err),
	}, nil
}

func ToGRPCParams(params *Params) *ddgrpc.Params {
	return &ddgrpc.Params{
		Station:params.Station,
		Start:params.Start,
		End:params.End,
		Breakdown:params.Breakdown,
		Tb:params.Tb,
		Tr:params.Tr,
		Method:params.Method,
		DayCalc:params.DayCalc,
	}
}

func ToParams(params *ddgrpc.Params) *Params {
	return &Params{
		Station:params.Station,
		Start:params.Start,
		End:params.End,
		Breakdown:params.Breakdown,
		Tb:params.Tb,
		Tr:params.Tr,
		Method:params.Method,
		DayCalc:params.DayCalc,
	}
}

func toGRPCDegree(degree *[]Degree) *[]*ddgrpc.Degree {
	res := make([]*ddgrpc.Degree, len(*degree))
	for i,v := range *degree {
		res[i] =  &ddgrpc.Degree{
			Date: v.Date,
			Temp: v.Temp,
		}
	}
	return &res
}

func ToDegree(degree *[]*ddgrpc.Degree) *[]Degree{
	res := make([]Degree, len(*degree))
	for i,v := range *degree {
		res[i] =  Degree{
			Date: v.Date,
			Temp: v.Temp,
		}
	}
	return &res
}