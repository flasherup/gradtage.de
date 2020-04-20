package apisvc

type GetHDDRequest struct {
	Params Params `json:"params"`
}

type GetHDDResponse struct {
	Data [][]string `json:"data"`
}

type GetHDDCSVRequest struct {
	Params Params `json:"params"`
}

type GetHDDCSVResponse struct {
	Data [][]string `json:"data"`
	FileName string `json:"fileName"`
}

type GetCDDCSVRequest struct {
	Params Params `json:"params"`
}

type GetCDDCSVResponse struct {
	Data [][]string `json:"data"`
	FileName string `json:"fileName"`
}

type GetSourceDataRequest struct {
	Params ParamsSourceData `json:"params"`
}

type GetSourceDataResponse struct {
	Data [][]string `json:"data"`
	FileName string `json:"fileName"`
}

type SearchRequest struct {
	Params ParamsSearch `json:"params"`
}

type SearchResponse struct {
	Data [][]string `json:"data"`
}