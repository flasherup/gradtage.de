package alertsys

import "github.com/flasherup/gradtage.de/alertsvc"

type AlertSystem interface {
	SendAlert(alert alertsvc.Alert) error
	SendEmail(alert alertsvc.Email) error
}