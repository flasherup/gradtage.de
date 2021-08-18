package autocompletesvc

type Client interface {
	GetAutocomplete(text string) (map[string][]Autocomplete, error)
	AddSources(sources []Autocomplete) error
	ResetSources(sources []Autocomplete) error
}