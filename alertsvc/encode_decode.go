package alertsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc/altgrpc"
)

func EncodeSendAlertResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(SendAlertResponse)
	return &altgrpc.SendAlertResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeSendAlertRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*altgrpc.SendAlertRequest)
	return SendAlertRequest{ decodeAlert(req.Alert) }, nil
}

func decodeAlert(alert *altgrpc.Alert) (res Alert) {
	res = Alert{
		Name:alert.Name,
		Desc:alert.Desc,
		Params:alert.Params,
	}

	return res
}

func errorToString(err error) string{
	if err == nil {
		return "nil"
	}

	return err.Error()
}


