package stationssvc

import (
	"context"
	"github.com/flasherup/gradtage.de/stationssvc/grpc"
)

func EncodeGetStationsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetStationsResponse)
	encStations := toGRPCMap(res.Stations)
	return &grpc.GetStationsResponse {
		Sts: encStations,
		Err: res.Err.Error(),
	}, nil
}

func DecodeGetStationsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpc.GetStationsRequest)
	return GetStationsRequest{req.Ids}, nil
}

func EncodeGetAllStationsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetAllStationsResponse)
	encStations := toGRPCMap(res.Stations)
	return &grpc.GetAllStationsResponse {
		Sts: encStations,
		Err: res.Err.Error(),
	}, nil
}

func DecodeGeAllStationsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return GetAllStationsRequest{}, nil
}

func EncodeAddStationsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(AddStationsResponse)
	return &grpc.AddStationsResponse {
		Err: res.Err.Error(),
	}, nil
}


func DecodeAddStationsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpc.AddStationsRequest)
	encStations := make([]Station, len(req.Sts))
	for i,v := range req.Sts {
		encStations[i] = Station {
			ID:v.Id,
			Name:v.Name,
			Timezone:v.Timezone,
		}
	}
	return AddStationsRequest{encStations}, nil
}

func toGRPCMap(src map[string]Station) map[string]*grpc.Station {
	res := make(map[string]*grpc.Station)
	for k,v := range src {
		res[k] = &grpc.Station {
			Id:v.ID,
			Name:v.Name,
			Timezone:v.Timezone,
		}
	}
	return res
}


