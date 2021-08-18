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

func EncodeResetSourcesResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(ResetSourcesResponse)
	return &acrpc.ResetSourcesResponse {
		Err:common.ErrorToString(res.Err),
	}, nil
}

func DecodeResetSourcesRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*acrpc.ResetSourcesRequest)
	encSources := DecodeSources(req.Sources)
	return ResetSourcesRequest{
		Sources:encSources,
	}, nil
}

func EncodeSources(sources []Autocomplete) []*acrpc.Source {
	res := make([]*acrpc.Source, len(sources))
	for i,v := range sources {
		res[i] = &acrpc.Source {
			ID:v.ID,
			SourceID:v.SourceID,
			Latitude:v.Latitude,
			Longitude:v.Longitude,
			Source:v.Source,
			Reports:v.Reports,
			ISO2Country:v.ISO2Country,
			ISO3Country:v.ISO3Country,
			Prio:v.Prio,
			CityNameEnglish:v.CityNameEnglish,
			CityNameNative:v.CityNameNative,
			CountryNameEnglish:v.CountryNameEnglish,
			CountryNameNative:v.CountryNameNative,
			ICAO:v.ICAO,
			WMO:v.WMO,
			CWOP:v.CWOP,
			Maslib:v.Maslib,
			National_ID:v.National_ID,
			IATA:v.IATA,
			USAF_WBAN:v.USAF_WBAN,
			GHCN:v.GHCN,
			NWSLI:v.NWSLI,
			Elevation:v.Elevation,
		}
	}
	return res
}

func DecodeSources(sources []*acrpc.Source) []Autocomplete {
	res := make([]Autocomplete, len(sources))
	for i,v := range sources {
		res[i] = Autocomplete{
			ID:v.ID,
			SourceID:v.SourceID,
			Latitude:v.Latitude,
			Longitude:v.Longitude,
			Source:v.Source,
			Reports:v.Reports,
			ISO2Country:v.ISO2Country,
			ISO3Country:v.ISO3Country,
			Prio:v.Prio,
			CityNameEnglish:v.CityNameEnglish,
			CityNameNative:v.CityNameNative,
			CountryNameEnglish:v.CountryNameEnglish,
			CountryNameNative:v.CountryNameNative,
			ICAO:v.ICAO,
			WMO:v.WMO,
			CWOP:v.CWOP,
			Maslib:v.Maslib,
			National_ID:v.National_ID,
			IATA:v.IATA,
			USAF_WBAN:v.USAF_WBAN,
			GHCN:v.GHCN,
			NWSLI:v.NWSLI,
			Elevation:v.Elevation,
		}
	}
	return res
}

func EncodeSourcesMap(sources map[string][]Autocomplete) map[string]*acrpc.Sources {
	res := make(map[string]*acrpc.Sources)
	for k,s := range sources {
		src := make([]*acrpc.Source, len(s))
		for i,v := range s{
			src[i] = &acrpc.Source {
				ID:v.ID,
				SourceID:v.SourceID,
				Latitude:v.Latitude,
				Longitude:v.Longitude,
				Source:v.Source,
				Reports:v.Reports,
				ISO2Country:v.ISO2Country,
				ISO3Country:v.ISO3Country,
				Prio:v.Prio,
				CityNameEnglish:v.CityNameEnglish,
				CityNameNative:v.CityNameNative,
				CountryNameEnglish:v.CountryNameEnglish,
				CountryNameNative:v.CountryNameNative,
				ICAO:v.ICAO,
				WMO:v.WMO,
				CWOP:v.CWOP,
				Maslib:v.Maslib,
				National_ID:v.National_ID,
				IATA:v.IATA,
				USAF_WBAN:v.USAF_WBAN,
				GHCN:v.GHCN,
				NWSLI:v.NWSLI,
				Elevation:v.Elevation,
			}
		}
		res[k] = &acrpc.Sources{Sources:src}
	}
	return res
}

func DecodeSourcesMap(sources map[string]*acrpc.Sources) map[string][]Autocomplete {
	res := make(map[string][]Autocomplete)
	for k,s := range sources {
		src := make([]Autocomplete, len(s.Sources))
		for i,v := range s.Sources{
			src[i] = Autocomplete{
				ID:v.ID,
				SourceID:v.SourceID,
				Latitude:v.Latitude,
				Longitude:v.Longitude,
				Source:v.Source,
				Reports:v.Reports,
				ISO2Country:v.ISO2Country,
				ISO3Country:v.ISO3Country,
				Prio:v.Prio,
				CityNameEnglish:v.CityNameEnglish,
				CityNameNative:v.CityNameNative,
				CountryNameEnglish:v.CountryNameEnglish,
				CountryNameNative:v.CountryNameNative,
				ICAO:v.ICAO,
				WMO:v.WMO,
				CWOP:v.CWOP,
				Maslib:v.Maslib,
				National_ID:v.National_ID,
				IATA:v.IATA,
				USAF_WBAN:v.USAF_WBAN,
				GHCN:v.GHCN,
				NWSLI:v.NWSLI,
				Elevation:v.Elevation,
			}
		}
		res[k] = src
	}
	return res
}


