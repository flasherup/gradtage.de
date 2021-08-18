package autocompletesvc

type GetAutocompleteRequest struct {
	Text string `json:"text"`
}

type GetAutocompleteResponse struct {
	Result 	map[string][]Autocomplete `json:"result"`
	Err     error                       `json:"err"`
}

type  AddSourcesRequest struct {
	Sources []Autocomplete `json:"sources"`
}

type AddSourcesResponse struct {
	Err     error  `json:"err"`
}

type ResetSourcesRequest struct {
	Sources []Autocomplete `json:"sources"`
}

type ResetSourcesResponse struct {
	Err     error  `json:"err"`
}