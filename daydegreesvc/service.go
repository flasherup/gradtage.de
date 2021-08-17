package daydegreesvc

import (
	"context"
)

type Params struct {
	Station string
	Start string
	End string
	Breakdown int
	Tb float64
	Tr float64
	Method string
	DayCalc int
}

type Degree struct {
	Date string
	Temp float64
}


type Service interface {
	GetDegree(ctx context.Context, params Params) ([]Degree, error)
}
