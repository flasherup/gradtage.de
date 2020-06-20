package alertsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/alertsvc/grpcalt"
)

func EncodeSendAlertResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(SendAlertResponse)
	return &grpcalt.SendAlertResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeSendAlertRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcalt.SendAlertRequest)
	return SendAlertRequest{ decodeAlert(req.Alert) }, nil
}

func EncodeSendEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(SendEmailResponse)
	return &grpcalt.SendEmailResponse {
		Err: errorToString(res.Err),
	}, nil
}

func DecodeSendEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*grpcalt.SendEmailRequest)
	return SendEmailRequest{ decodeEmail(req.Email) }, nil
}

func decodeAlert(alert *grpcalt.Alert) (res Alert) {
	res = Alert{
		Name:alert.Name,
		Desc:alert.Desc,
		Params:alert.Params,
	}

	return res
}

func decodeEmail(email *grpcalt.Email) (res Email) {
	params := make(map[string]string)
	for k,v := range email.Params {
		params[k] = v
	}
	res = Email{
		Name:email.Name,
		Email:email.Email,
		Params:params,
	}

	return res
}

func EncodeEmail(src Email) *grpcalt.Email {
	params := make(map[string]string)
	for k,v := range src.Params {
		params[k] = v
	}
	return &grpcalt.Email{
		Name:	src.Name,
		Email:	src.Email,
		Params:	params,
	}
}

func errorToString(err error) string{
	if err == nil {
		return "nil"
	}

	return err.Error()
}


