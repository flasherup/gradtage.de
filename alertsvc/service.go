package alertsvc
import (
	"context"
)


type Alert struct {
	Name 		string 				`json:"name"`
	Desc 		string 				`json:"desc"`
	Params 	 	map[string]string 	`json:params`
}

type Email struct {
	Name 		string 				`json:"name"`
	Email 		string 				`json:"email"`
	Params 	 	map[string]string 	`json:params`
}

type Service interface {
	SendAlert(ctx context.Context, alert Alert) error
	SendEmail(ctx context.Context, email Email) error
}