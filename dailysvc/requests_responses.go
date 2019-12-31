package dailysvc

type GetPeriodRequest struct {
	ID    string `json:"id"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type GetPeriodResponse struct {
	Temps []Temperature `json:"temp"`
	Err  error `json:"err"`
}


type PushPeriodRequest struct {
	ID    string        `json:"id"`
	Temps []Temperature `json:"temp"`
}

type PushPeriodResponse struct {
	Err  error `json:"err"`
}

type GetUpdateDateRequest struct {
	IDs 	[]string `json:"ids"`
}

type GetUpdateDateResponse struct {
	Dates  map[string]string `json:"dates"`
	Err  error `json:"err"`
}

