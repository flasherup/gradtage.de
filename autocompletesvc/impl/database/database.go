package database

type AutocompleteDB interface {
	GetAutocomplete(text string) (result map[string]string, err error)
	CreateTable() (err error)
	RemoveTable() (err error)
	Dispose()
}