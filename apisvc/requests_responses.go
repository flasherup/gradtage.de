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

type UserRequest struct {
	Params ParamsUser `json:"params"`
}

type UserResponse struct {
	Data [][]string `json:"data"`
}

type PlanRequest struct {
	Params ParamsPlan `json:"params"`
}

type PlanResponse struct {
	Data [][]string `json:"data"`
}

type StripeRequest struct {
	Event StripeEvent `json:"event"`
}

type StripeResponse struct {
	Json string `json:"json"`
}