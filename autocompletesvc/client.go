package autocompletesvc

type Client interface {
	GetAutocomplete(text string) (map[string][]Source, error)
	AddSources(sources []Source) error
}