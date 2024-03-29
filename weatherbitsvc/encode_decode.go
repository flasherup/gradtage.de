package weatherbitsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/common"
	weathergrpc "github.com/flasherup/gradtage.de/weatherbitsvc/weatherbitgrpc"
)

func EncodeGetPeriodResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetPeriodResponse)
	encTemp := toGRPCTemps(res.Temps)
	return &weathergrpc.GetPeriodResponse {
		Temps: encTemp,
		Err: common.ErrorToString(res.Err),
	}, nil
}

func DecodeGetWBPeriodRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*weathergrpc.GetWBPeriodRequest)
	return GetWBPeriodRequest{req.Id, req.Start, req.End}, nil
}

func EncodeGetWBPeriodResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetWBPeriodResponse)
	encTemp := ToGRPCWBData(res.Temps)
	return &weathergrpc.GetWBPeriodResponse {
		Temps: encTemp,
		Err: common.ErrorToString(res.Err),
	}, nil
}

func DecodePushWBPeriodRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*weathergrpc.PushWBPeriodRequest)
	return PushWBPeriodRequest{req.Id, *ToWBData(req.Data)}, nil
}

func EncodePushWBPeriodResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(PushWBPeriodResponse)
	return &weathergrpc.PushWBPeriodResponse {
		Err: common.ErrorToString(res.Err),
	}, nil
}

func DecodeGetPeriodRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*weathergrpc.GetPeriodRequest)
	return GetPeriodRequest{req.Ids, req.Start, req.End}, nil
}

func EncodeGetUpdateDateResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetUpdateDateResponse)
	return &weathergrpc.GetUpdateDateResponse {
		Dates: res.Dates,
		Err: common.ErrorToString(res.Err),
	}, nil
}

func DecodeGetUpdateDateRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*weathergrpc.GetUpdateDateRequest)
	return GetUpdateDateRequest{req.Ids}, nil
}

func EncodeGetStationsListResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetStationsListResponse)
	return &weathergrpc.GetStationsListResponse {
		List: res.List,
		Err: common.ErrorToString(res.Err),
	}, nil
}

func DecodeGetStationsListRequest(_ context.Context, r interface{}) (interface{}, error) {
	return GetStationsListRequest{}, nil
}

func EncodeGetAverageResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetAverageResponse)
	encTemp := ToGRPCAverageTemps(res.Temps)
	return &weathergrpc.GetAverageResponse {
		List: encTemp,
		Err: common.ErrorToString(res.Err),
	}, nil
}

func DecodeGetAverageRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*weathergrpc.GetAverageRequest)
	return GetAverageRequest{req.Id, int(req.Years), req.End}, nil
}



func ToGRPCAverageTemps(src []common.Temperature) []*weathergrpc.Temperature {
	temps := make([]*weathergrpc.Temperature, len(src))
	for i,v := range src {
		temps[i] = &weathergrpc.Temperature{
			Date: 	v.Date,
			Temp: 	v.Temp,
		}
	}
		return temps
}

func ToCommonAverageTemps(src []*weathergrpc.Temperature) []common.Temperature {
	temps := make([]common.Temperature, len(src))
	for i,v := range src {
		temps[i] = common.Temperature{
			Date: 	v.Date,
			Temp: 	v.Temp,
		}
	}
	return temps
}

func ToGRPCWBData(src []WBData)  []*weathergrpc.WBData {
	res := make([]*weathergrpc.WBData, len(src))
	for i,v := range src {
		res[i] = &weathergrpc.WBData {
			Date:v.Date,
			Rh:v.Rh,
			Pod:v.Pod,
			Pres:v.Pres,
			Timezone:v.Timezone,
			CountryCode:v.CountryCode,
			Clouds:v.Clouds,
			Vis:v.Vis,
			SolarRad:v.SolarRad,
			WindSpd:v.WindSpd,
			StateCode:v.StateCode,
			CityName:v.CityName,
			AppTemp:v.AppTemp,
			Uv:v.Uv,
			Lon:v.Lon,
			Slp:v.Slp,
			HAngle:v.HAngle,
			Dewpt:v.Dewpt,
			Snow:v.Snow,
			Aqi:v.Aqi,
			WindDir:v.WindDir,
			ElevAngle:v.ElevAngle,
			Ghi:v.Ghi,
			Lat:v.Lat,
			Precip:v.Precip,
			Sunset:v.Sunset,
			Temp:v.Temp,
			Station:v.Station,
			Dni:v.Dni,
			Sunrise:v.Sunrise,
		}
	}
	return res
}

func ToWBData(src []*weathergrpc.WBData) *[]WBData {
	res := make([]WBData, len(src))
	for i,v := range src {
		res[i] = WBData {
			Date:v.Date,
			Rh:v.Rh,
			Pod:v.Pod,
			Pres:v.Pres,
			Timezone:v.Timezone,
			CountryCode:v.CountryCode,
			Clouds:v.Clouds,
			Vis:v.Vis,
			SolarRad:v.SolarRad,
			WindSpd:v.WindSpd,
			StateCode:v.StateCode,
			CityName:v.CityName,
			AppTemp:v.AppTemp,
			Uv:v.Uv,
			Lon:v.Lon,
			Slp:v.Slp,
			HAngle:v.HAngle,
			Dewpt:v.Dewpt,
			Snow:v.Snow,
			Aqi:v.Aqi,
			WindDir:v.WindDir,
			ElevAngle:v.ElevAngle,
			Ghi:v.Ghi,
			Lat:v.Lat,
			Precip:v.Precip,
			Sunset:v.Sunset,
			Temp:v.Temp,
			Station:v.Station,
			Dni:v.Dni,
			Sunrise:v.Sunrise,
		}
	}
	return &res
}

func toGRPCTemps(src map[string][]common.Temperature)  map[string]*weathergrpc.Temperatures {
	res := make(map[string]*weathergrpc.Temperatures)
	for k,v := range src {
		temps := make([]*weathergrpc.Temperature, len(v))
		for i,t := range v {
			temps[i] = &weathergrpc.Temperature{
				Date: 			t.Date,
				Temp: 	t.Temp,
			}
		}
		res[k] = &weathergrpc.Temperatures {
			Temps: temps,
		}
	}
	return res
}

func ToCommonTemps(src map[string]*weathergrpc.Temperatures) map[string][]common.Temperature {
	res := make(map[string][]common.Temperature)
	for k,v := range src {
		temps := make([]common.Temperature, len(v.Temps))
		for i,t := range v.Temps {
			temps[i] = common.Temperature{
				Date: 	t.Date,
				Temp: 	t.Temp,
			}
		}
		res[k] = temps
	}
	return res
}
