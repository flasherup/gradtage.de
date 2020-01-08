package alertsvc

import (
	"github.com/flasherup/gradtage.de/alertsvc/altgrpc"
)

type Client interface {
	SendAlert(alert Alert) 	(resp *altgrpc.SendAlertResponse, err error)
}