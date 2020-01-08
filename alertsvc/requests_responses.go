package alertsvc

type SendAlertRequest struct {
	Alert Alert `json:"alert"`
}

type SendAlertResponse struct {
	Err      error              `json:"err"`
}
