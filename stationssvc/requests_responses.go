package stationssvc

type GetStationsRequest struct {
	IDs []string `json:"ids"`
}

type GetStationsResponse struct {
	Stations map[string]Station `json:"sts"`
	Err      error              `json:"err"`
}

type GetAllStationsRequest struct {
}

type GetAllStationsResponse struct {
	Stations map[string]Station `json:"sts"`
	Err      error              `json:"err"`
}

type GetStationsBySrcTypeRequest struct {
	Types []string
}

type GetStationsBySrcTypeResponse struct {
	Stations map[string]Station `json:"sts"`
	Err      error              `json:"err"`
}

type AddStationsRequest struct {
	Stations []Station `json:"sts"`
}

type AddStationsResponse struct {
	Err error `json:"err"`
}
