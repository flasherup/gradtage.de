package hrlaggregatorsvc

import (
	"context"
)


type Status struct {
	Station 	string 	`json:"station"`
	Update 		string 	`json:"update"`
	Temperature	float32 `json:"temperature"`

}

type Service interface {
	GetStatus(ctx context.Context) (status []Status, err error)
}