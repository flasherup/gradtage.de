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