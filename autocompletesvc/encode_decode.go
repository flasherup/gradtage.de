package autocompletesvc

import (
	"context"
	"github.com/flasherup/gradtage.de/autocompletesvc/acrpc"
	"github.com/flasherup/gradtage.de/common"
)

func EncodeGetAutocompleteResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetAutocompleteResponse)
	sources := EncodeSourcesMap(res.Result)
	return &acrpc.GetAutocompleteResponse {
		Result: sources,
		Err: common.ErrorToString(res.Err),
	}, nil
}

func DecodeGetAutocompleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*acrpc.GetAutocompleteRequest)
	return GetAutocompleteRequest{req.Text}, nil
}

func EncodeAddSourcesResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(AddSourcesResponse)
	return &acrpc.AddSourcesResponse {
		Err:common.ErrorToString(res.Err),
	}, nil
}

func DecodeAddSourcesRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*acrpc.AddSourcesRequest)
	encSources := DecodeSources(req.Sources)
	return AddSourcesRequest{
		Sources:encSources,
	}, nil
}

func EncodeSources(sources []Source) []*acrpc.Source {
	res := make([]*acrpc.Source, len(sources))
	for i,v := range sources {
		res[i] = &acrpc.Source {
			ID: v.ID,
			Name: v.Name,
			Icao: v.Icao,
			Dwd: v.Dwd,
			Wmo: v.Wmo,
		}
	}
	return res
}

func DecodeSources(sources []*acrpc.Source) []Source {
	res := make([]Source, len(sources))
	for i,v := range sources {
		res[i] = Source {
			ID: v.ID,
			Name: v.Name,
			Icao: v.Icao,
			Dwd: v.Dwd,
			Wmo: v.Wmo,
		}
	}
	return res
}

func EncodeSourcesMap(sources map[string][]Source) map[string]*acrpc.Sources {
	res := make(map[string]*acrpc.Sources)
	for k,v := range sources {
		src := make([]*acrpc.Source, len(v))
		for i,s := range v{
			src[i] = &acrpc.Source {
				ID: s.ID,
				Name: s.Name,
				Icao: s.Icao,
				Dwd: s.Dwd,
				Wmo: s.Wmo,
			}
		}
		res[k] = &acrpc.Sources{Sources:src}
	}
	return res
}

func DecodeSourcesMap(sources map[string]*acrpc.Sources) map[string][]Source {
	res := make(map[string][]Source)
	for k,v := range sources {
		src := make([]Source, len(v.Sources))
		for i,s := range v.Sources{
			src[i] = Source {
				ID: s.ID,
				Name: s.Name,
				Icao: s.Icao,
				Dwd: s.Dwd,
				Wmo: s.Wmo,
			}
		}
		res[k] = src
	}
	return res
}


