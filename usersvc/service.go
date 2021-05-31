package usersvc

import (
	"context"
	"time"
)

const (
	PlanCanceled     = "canceled"
	PlanAdmin        = "admin"
	PlanTrial        = "trial"
	PlanLite         = "lite"
	PlanProfessional = "professional"
	PlanEnterprise   = "enterprise"
)

type Selection struct {
	Key 		string
	StationID 	string
	Method 		string
	Start 		time.Time
	End 		time.Time
}

type Parameters struct {
	User User
	Plan Plan
}

type Plan struct {
	Name 		string
	Stations 	int//Number of stations that user can get
	Limitation 	int //Number of requests per hour
	HDD 		bool
	DD 			bool
	CDD 		bool
	Start  		time.Time //Start time of data that user can get
	End 		time.Time //End time of data that user can get
	Period 		int //Period of days that key is valid
	Admin 		bool //Administrator rights
}

type User struct {
	Name 		string
	Key 		string
	RenewDate  	time.Time //Time the key was activated
	RequestDate time.Time //Latest request time
	Requests 	int //Number of request during hour
	Plan 		string //Plan name
	Stations 	[]string //The list of stations
}

type Service interface {
	CreateUser(ctx context.Context, userName string, plan string, key string, email bool) (string, error)
	UpdateUser(ctx context.Context, user User, email bool) (string, error)
	DeleteUser(ctx context.Context, user User) error
	AddPlan(ctx context.Context, plan Plan) error
	ValidateSelection(ctx context.Context, selection Selection) (bool, error)
	ValidateKey(ctx context.Context, key string) (Parameters, error)
	ValidateName(ctx context.Context, name string) (Parameters, error)
}