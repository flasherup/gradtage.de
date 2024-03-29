package apisvc

type GetDataRequest struct {
	Params []Params `json:"params"`
}

type GetDataResponse struct {
	Data   []*DDResponse `json:"data"`
	Format string        `json:"format"`
	Err    error         `json:"err"`
}

type GetHDDRequest struct {
	Params Params `json:"params"`
}

type GetHDDResponse struct {
	Data CSVData `json:"data"`
}

type GetSourceDataRequest struct {
	Params ParamsSourceData `json:"params"`
}

type GetSourceDataResponse struct {
	Data     CSVData `json:"data"`
	FileName string  `json:"fileName"`
}

type SearchRequest struct {
	Params ParamsSearch `json:"params"`
}

type SearchResponse struct {
	Data CSVData `json:"data"`
}

type UserRequest struct {
	Params ParamsUser `json:"params"`
}

type UserResponse struct {
	Data CSVData `json:"data"`
}

type StripeResponse struct {
	Json string `json:"json"`
}

type WoocommerceRequest struct {
	Event WoocommerceEvent `json:"event"`
}

type WoocommerceResponse struct {
	Json string `json:"json"`
}

type ServiceRequest struct {
	Name   string            `json:"name"`
	Params map[string]string `json:"params"`
}

type ServiceResponse struct {
	Json interface{} `json:"json"`
}
