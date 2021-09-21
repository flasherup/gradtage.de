package autocompletesvc

import (
	"github.com/flasherup/gradtage.de/autocompletesvc/acrpc"
)

type Client interface {
	GetAutocomplete(text string) (map[string][]Autocomplete, error)
	AddSources(sources []Autocomplete) error
	ResetSources(sources []Autocomplete) error
	GetAllStations() (map[string]*acrpc.Source, error)
}