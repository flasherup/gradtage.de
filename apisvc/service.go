package apisvc

import "context"

type Params struct {
	Key 	string  `json:"key"`
	Station string  `json:"station"`
	Start   string  `json:"start"`
	End     string  `json:"end"`
	HL 		float64 `json:"hl"`
	RT  	float64 `json:"rt"`
	Output	string  `json:"output"`
}

type ParamsSourceData struct {
	Key 	string  `json:"key"`
	Station string  `json:"station"`
	Start   string  `json:"start"`
	End     string  `json:"end"`
	Type	string	`json:"type"`
}

type Service interface {
	GetHDD(ctx context.Context, params Params) (data [][]string, err error)
	GetHDDCSV(cts context.Context, params Params) (data [][]string, fileName string, err error)
	GetSourceData(ctx context.Context, params ParamsSourceData) (data [][]string, fileName string, err error)
}