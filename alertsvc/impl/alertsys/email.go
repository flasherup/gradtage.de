package alertsys

import (
	"bytes"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/alertsvc/config"
	"html/template"
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

func (ue EmailAlert) SendEmail(email alertsvc.Email) error {
	auth := smtp.PlainAuth(
		"",
		ue.config.User,
		ue.config.Pass,
		ue.config.Host,
	)

	parameters := struct {
		FromName        string
		From 			string
		To 				string
		Subject 		string
		Text			string
		Email 			string
		Key 			string
		Plan 			string
	}{
		"User Name",
		"mytest@mytest.mytest",
		"Do Not Replay",
		"User Plan Upgrade",
		"Congratulations your plan is updated",
		email.Email,
		email.Params["key"],
		email.Params["plan"],
	}

	buffer := new(bytes.Buffer)

	//tmpl := template.Must(template.New("emailTemplate").Parse(emailTemplates.UserPlanUpdateTemplate))
	tmpl, err := template.ParseFiles(ue.config.EmailTemplates.UserPlanUpdate)
	if err != nil {
		return err
	}
	tmpl.Execute(buffer, &parameters)

	return smtp.SendMail(
		ue.config.Host + ":" + ue.config.Port,
		auth,
		ue.config.From,
		[]string{email.Email},
		buffer.Bytes(),
	)
}

func paramsToString(params map[string]string) string {
	res := ""
	for k,v := range params {
		res += k + ": " + v + "\n"
	}
	return res
}
