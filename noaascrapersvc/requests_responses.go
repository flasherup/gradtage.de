package noaascrapersvc

import "github.com/flasherup/gradtage.de/hourlysvc"

type GetPeriodRequest struct {
	ID    string `json:"id"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type GetPeriodResponse struct {
	Temps []hourlysvc.Temperature `json:"temps"`
	Err  error `json:"err"`
}

type GetUpdateDateRequest struct {
	IDs 	[]string `json:"ids"`
}

type GetUpdateDateResponse struct {
	Dates  map[string]string `json:"dates"`
	Err  error `json:"err"`
}