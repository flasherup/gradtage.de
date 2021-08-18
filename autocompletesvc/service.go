package autocompletesvc

import (
	"context"
)

type Source struct {
	ID                 string  `json:"id"`
	SourceID           string  `json:"source_id"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Source             string  `json:"source"`
	Reports            string  `json:"reports"`
	ISO2Country        string  `json:"iso_2_country"`
	ISO3Country        string  `json:"iso_3_country"`
	Prio               string  `json:"prio"`
	CityNameEnglish    string  `json:"city_name_english"`
	CityNameNative     string  `json:"city_name_native"`
	CountryNameEnglish string  `json:"country_name_english"`
	CountryNameNative  string  `json:"country_name_native"`
	ICAO               string  `json:"icao"`
	WMO                string  `json:"wmo"`
	CWOP               string  `json:"cwop"`
	Maslib             string  `json:"maslib"`
	National_ID        string  `json:"national_id"`
	IATA               string  `json:"iata"`
	USAF_WBAN          string  `json:"usaf_wban"`
	GHCN               string  `json:"ghcn"`
	NWSLI              string  `json:"nwsli"`
	Elevation          float64 `json:"elevation"`
}

type Service interface {
	GetAutocomplete(ctx context.Context, text string) (result map[string][]Source, err error)
	AddSources(ctx context.Context, sources []Source) (err error)
	ResetSources(ctx context.Context, sources []Source) (err error)
}
