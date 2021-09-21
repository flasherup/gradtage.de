package database

import (
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/autocompletesvc/acrpc"
)

type AutocompleteDB interface {
	GetAutocomplete(text string) (result map[string][]autocompletesvc.Autocomplete, err error)
	GetStationId(text string) (result map[string][]autocompletesvc.Autocomplete, err error)
	AddSources(sources []autocompletesvc.Autocomplete) (err error)
	GetAllStations() (map[string]*acrpc.Source,error)
	CreateTable() (err error)
	RemoveTable() (err error)
	Dispose()
}