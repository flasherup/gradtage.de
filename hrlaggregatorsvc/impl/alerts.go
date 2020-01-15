package impl

import (
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/hourlysvc/hrlgrpc"
)

func NewTemperatureChangeAlert(prev hrlgrpc.Temperature, cur hourlysvc.Temperature , stsId string) alertsvc.Alert {
	p := fmt.Sprintf("date: %s, temp:%g:", prev.Date, prev.Temperature)
	c := fmt.Sprintf("date: %s, temp:%g:",cur.Date, cur.Temperature)
	return alertsvc.Alert{
		Name: "Plausibility",
		Desc: "Hourly temperature change more than 10°",
		Params: map[string]string{ "Previous":p,  "Current":c , "Station":stsId },
	}
}

func NewTemperatureGapAlert(prev hrlgrpc.Temperature, cur hourlysvc.Temperature, stsId string) alertsvc.Alert {
	p := fmt.Sprintf("date: %s, temp:%g:", prev.Date, prev.Temperature)
	c := fmt.Sprintf("date: %s, temp:%g:",cur.Date, cur.Temperature)
	return alertsvc.Alert{
		Name: "Plausibility",
		Desc: "Hourly temperature update gap more then 3 hours",
		Params: map[string]string{ "Previous":p,  "Current":c, "Station":stsId },
	}
}

func NewTemperatureSameDateAlert(prev hrlgrpc.Temperature, cur hourlysvc.Temperature) alertsvc.Alert {
	p := fmt.Sprintf("date: %s, temp:%g:", prev.Date, prev.Temperature)
	c := fmt.Sprintf("date: %s, temp:%g:",cur.Date, cur.Temperature)
	return alertsvc.Alert{
		Name: "Plausibility",
		Desc: "Hourly temperature update period less the hour",
		Params: map[string]string{ "Previous":p,  "Current":c },
	}
}

func NewErrorAlert(err error) alertsvc.Alert {
	return alertsvc.Alert{
		Name: "Error",
		Desc: "Hourly Aggregator Service error",
		Params: map[string]string{ "Err":err.Error() },
	}
}

func NewNotificationAlert(notification string) alertsvc.Alert {
	return alertsvc.Alert{
		Name: "Notification",
		Desc: "Hourly Aggregator Service notification",
		Params: map[string]string{ "note":notification },
	}
}
