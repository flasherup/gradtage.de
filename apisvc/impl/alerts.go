package impl

import (
	"github.com/flasherup/gradtage.de/alertsvc"
)

func NewErrorAlert(err error) alertsvc.Alert {
	return alertsvc.Alert{
		Name: "Error",
		Desc: "API Service error",
		Params: map[string]string{ "Err":err.Error() },
	}
}

func NewNotificationAlert(notification string) alertsvc.Alert {
	return alertsvc.Alert{
		Name: "Notification",
		Desc: "API Service notification",
		Params: map[string]string{ "note":notification },
	}
}
