package alertsvc

type Client interface {
	SendAlert(alert Alert) 	error
}