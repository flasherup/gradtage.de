package alertsvc

type Client interface {
	SendAlert(alert Alert) 	error
	SendEmail(email Email) 	error
}