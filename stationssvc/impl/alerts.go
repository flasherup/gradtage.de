package impl

import (
	"github.com/flasherup/gradtage.de/alertsvc"
)

func NewErrorAlert(err error) alertsvc.Alert {
	return alertsvc.Alert{
		Name: "Error",
		Desc: "Stations Service error",
		Params: map[string]string{ "Err":err.Error() },
	}
}

func NewNotificationAlert(notification string) alertsvc.Alert {
	return alertsvc.Alert{
		Name: "Notification",
		Desc: "Stations Service notification",
		Params: map[string]string{ "note":notification },
	}
}
