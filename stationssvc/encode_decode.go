package stationssvc

import (
	"context"
	"github.com/flasherup/gradtage.de/stationssvc/stsgrpc"
)

func EncodeGetStationsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetStationsResponse)
	encStations := toGRPCMap(res.Stations)
	return &stsgrpc.GetStationsResponse {
		Sts: encStations,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetStationsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*stsgrpc.GetStationsRequest)
	return GetStationsRequest{req.Ids}, nil
}

func EncodeGetAllStationsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetAllStationsResponse)
	encStations := toGRPCMap(res.Stations)
	return &stsgrpc.GetAllStationsResponse {
		Sts: encStations,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGeAllStationsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return GetAllStationsRequest{}, nil
}

func EncodeGetStationsBySrcTypeResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetStationsBySrcTypeResponse)
	encStations := toGRPCMap(res.Stations)
	return &stsgrpc.GetStationsBySrcTypeResponse {
		Sts: encStations,
		Err: errorToString(res.Err),
	}, nil
}

func DecodeGetStationsBySrcTypeRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*stsgrpc.GetStationsBySrcTypeRequest)
	return GetStationsBySrcTypeRequest{req.Types}, nil
}

func EncodeAddStationsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(AddStationsResponse)
	return &stsgrpc.AddStationsResponse {
		Err: errorToString(res.Err),
	}, nil
}


func DecodeAddStationsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*stsgrpc.AddStationsRequest)
	encStations := make([]Station, len(req.Sts))
	for i,v := range req.Sts {
		encStations[i] = Station {
			ID:v.Id,
			Name:v.Name,
			Timezone:v.Timezone,
			SourceType:v.SourceType,
			SourceID:v.SourceId,
		}
	}
	return AddStationsRequest{encStations}, nil
}

func EncodeResetStationsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(ResetStationsResponse)
	return &stsgrpc.ResetStationsResponse {
		Err: errorToString(res.Err),
	}, nil
}


func DecodeResetStationsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*stsgrpc.ResetStationsRequest)
	encStations := make([]Station, len(req.Sts))
	for i,v := range req.Sts {
		encStations[i] = Station {
			ID:v.Id,
			Name:v.Name,
			Timezone:v.Timezone,
			SourceType:v.SourceType,
			SourceID:v.SourceId,
		}
	}
	return ResetStationsRequest{encStations}, nil
}

func toGRPCMap(src map[string]Station) map[string]*stsgrpc.Station {
	res := make(map[string]*stsgrpc.Station)
	for k,v := range src {
		res[k] = &stsgrpc.Station {
			Id:v.ID,
			Name:v.Name,
			Timezone:v.Timezone,
			SourceType:v.SourceType,
			SourceId:v.SourceID,
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


