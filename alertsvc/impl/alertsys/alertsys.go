package alertsys

import "github.com/flasherup/gradtage.de/alertsvc"

type AlertSystem interface {
	Send(alert alertsvc.Alert) error
}