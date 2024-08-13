package database

import (
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/autocompletesvc/acrpc"
)

type AutocompleteDB interface {
	GetAutocomplete(text string) (result map[string][]autocompletesvc.Autocomplete, err error)
	AddSources(sources []autocompletesvc.Autocomplete) (err error)
	GetAllStations() (map[string]*acrpc.Source, error)
	CreateAutocompleteTable() (err error)
	RemoveAutocompleteTable() (err error)
	CreateZipCodeTable() error
	RemoveZipCodeTable() error
	AddZipCodes(zipCodes []autocompletesvc.ZipCode) (err error)
	Dispose()
}
