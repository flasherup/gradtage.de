package database

import "github.com/flasherup/gradtage.de/autocompletesvc"

type AutocompleteDB interface {
	GetAutocomplete(text string) (result map[string][]autocompletesvc.Autocomplete, err error)
	GetStationId(text string) (result map[string][]autocompletesvc.Autocomplete, err error)
	AddSources(sources []autocompletesvc.Autocomplete) (err error)
	CreateTable() (err error)
	RemoveTable() (err error)
	Dispose()
}