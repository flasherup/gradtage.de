package autocompletesvc

type GetAutocompleteRequest struct {
	Text string `json:"text"`
}

type GetAutocompleteResponse struct {
	Result 	map[string][]Source `json:"result"`
	Err     error  `json:"err"`
}

type  AddSourcesRequest struct {
	Sources []Source `json:"sources"`
}

type AddSourcesResponse struct {
	Err     error  `json:"err"`
}
