package autocompletesvc

import (
	"context"
)

type Source struct {
	ID 		string
	Name 	string
	Icao 	string
	Dwd 	string
	Wmo 	string
}

type Service interface {
	GetAutocomplete(ctx context.Context, text string) (result map[string][]Source, err error)
	AddSources(ctx context.Context, sources []Source) (err error)
}