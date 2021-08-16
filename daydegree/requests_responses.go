package weatherbitsvc

import (
	"github.com/flasherup/gradtage.de/common"
)

type GetPeriodRequest struct {
	IDs   []string `json:"ids"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type GetPeriodResponse struct {
	Temps 	map[string][]common.Temperature 	`json:"temps"`
	Err  	error 								`json:"err"`
}

type GetWBPeriodRequest struct {
	Id string `json:"id"`
	Start string`json:"start"`
	End   string `json:"end"`
}

type GetWBPeriodResponse struct {
	Temps []WBData `json:"temps"`
	Err error
}

type PushWBPeriodRequest struct {
	Id string `json:"id"`
	Data []WBData `json:"data"`
}

type PushWBPeriodResponse struct {
	Err error
}

type GetUpdateDateRequest struct {
	IDs 	[]string `json:"ids"`
}

type GetUpdateDateResponse struct {
	Dates  map[string]string `json:"dates"`
	Err  error `json:"err"`
}

type GetStationsListRequest struct {
}

type GetStationsListResponse struct {
	List []string `json:"list"`
	Err  error `json:"err"`
}