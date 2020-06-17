package alertsvc

type SendAlertRequest struct {
	Alert Alert `json:"alert"`
}

type SendAlertResponse struct {
	Err	error	`json:"err"`
}

type SendEmailRequest struct {
	Email Email `json:"email"`
}

type SendEmailResponse struct {
	Err	error	`json:"err"`
}

