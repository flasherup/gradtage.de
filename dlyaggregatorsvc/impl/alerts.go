package impl

import (
	"github.com/flasherup/gradtage.de/alertsvc"
)


func NewTemperatureGapAlert(id string) alertsvc.Alert {
	return alertsvc.Alert{
		Name: "Plausibility",
		Desc: "Daily aggregator, not enough hourly data",
		Params: map[string]string{ "Station":id },
	}
}

func NewErrorAlert(err error) alertsvc.Alert {
	return alertsvc.Alert{
		Name: "Error",
		Desc: "Daily aggregator Service error",
		Params: map[string]string{ "Err":err.Error() },
	}
}

func NewNotificationAlert(notification string) alertsvc.Alert {
	return alertsvc.Alert{
		Name: "Notification",
		Desc: "Daily aggregator Service notification",
		Params: map[string]string{ "note":notification },
	}
}
