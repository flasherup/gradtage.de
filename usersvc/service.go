package usersvc

import (
	"context"
	"time"
)

type Plan struct {
	Name 	string
	Start 	time.Time
	End 	time.Time
	Admin 	bool
}

type UserParameters struct {
	Name 		string
	Plan 		Plan
	Stations 	[]string
}

type Service interface {
	CreateUser(ctx context.Context, userName string, plan Plan)error
	CreateUserAuto(ctx context.Context, userName string, plan Plan)error
	SetPlan(ctx context.Context, userName string, plan Plan)error
	SetStations(ctx context.Context, userName string, stations []string)error
	ValidateKey(ctx context.Context, key string) (parameters UserParameters, err error)
}