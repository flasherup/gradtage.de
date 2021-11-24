package daydegreesvc

import (
	"context"
)

type Params struct {
	Station   string
	Start     string
	End       string
	Breakdown string
	Tb        float64
	Tr        float64
	Output    string
	DayCalc   string
}

type Degree struct {
	Date string
	Temp float64
}


type Service interface {
	GetDegree(ctx context.Context, params Params) ([]Degree, error)
	GetAverageDegree(ctx context.Context, params Params, years int) ([]Degree, error)
}
