package database

import "github.com/flasherup/gradtage.de/autocompletesvc"

type AutocompleteDB interface {
	GetAutocomplete(text string) (result map[string][]autocompletesvc.Source, err error)
	GetStationId(text string) (result map[string][]autocompletesvc.Source, err error)
	AddSources(sources []autocompletesvc.Source) (err error)
	CreateTable() (err error)
	RemoveTable() (err error)
	Dispose()
}