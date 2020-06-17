package alertsys

import (
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/alertsvc/config"
	"net/smtp"
)

type EmailAlert struct {
	config config.EmailConfig
}

func NewEmailAlertSystem(config config.EmailConfig) *EmailAlert {
	return &EmailAlert{
		config:config,
	}
}

func (ea EmailAlert) SendAlert(alert alertsvc.Alert) error {
	auth := smtp.PlainAuth(
		"",
		ea.config.User,
		ea.config.Pass,
		ea.config.Host,
	)

	body := fmt.Sprintf("Alert name: %s\n%s\n%s",
		alert.Name,
		alert.Desc,
		paramsToString(alert.Params),
	)

	msg := "From: " + ea.config.From + "\n" +
		"To: " + "Admin" + "\n" +
		"Subject: Alert notification\n\n" +
		body

	err := smtp.SendMail(
		ea.config.Host + ":" + ea.config.Port,
		auth,
		ea.config.From,
		ea.config.Recipients,
		[]byte(msg),
	)
	return err
}

func (ea EmailAlert) SendEmail(email alertsvc.Email) error {
	auth := smtp.PlainAuth(
		"",
		ea.config.User,
		ea.config.Pass,
		ea.config.Host,
	)

	body := fmt.Sprintf("Email name: %s\n%s\n%s",
		email.Name,
		email.Email,
		paramsToString(email.Params),
	)

	msg := "From: " + ea.config.From + "\n" +
		"To: " + "Admin" + "\n" +
		"Subject: Alert notification\n\n" +
		body

	err := smtp.SendMail(
		ea.config.Host + ":" + ea.config.Port,
		auth,
		ea.config.From,
		[]string{email.Email},
		[]byte(msg),
	)
	return err
}

func paramsToString(params map[string]string) string {
	res := ""
	for k,v := range params {
		res += k + ": " + v + "\n"
	}
	return res
}
