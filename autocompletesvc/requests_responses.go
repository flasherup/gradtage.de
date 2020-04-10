package stationssvc

type GetAutocompleteRequest struct {
	Text string `json:"text"`
}

type GetAutocompleteResponse struct {
	Result 	map[string]string `json:"result"`
	Err     error  `json:"err"`
}
