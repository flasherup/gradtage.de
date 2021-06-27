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
}

type Order struct {
	OrderId 	int
	Key			string
	Email		string
	Plan 		string
	Stations 	[]string
	RequestDate time.Time //Latest request time
	Requests 	int //Number of request during hour
	Admin		bool //True if user is administrator
}

type Service interface {
	CreateOrder(ctx context.Context, orderId int, email, plan, key string) (string, error)
	UpdateOrder(ctx context.Context, order Order) (string, error)
	DeleteOrder(ctx context.Context, orderId int) error
	AddPlan(ctx context.Context, plan Plan) error
	ValidateSelection(ctx context.Context, selection Selection) error
	ValidateKey(ctx context.Context, key string) (Order, Plan, error)
	ValidateOrder(ctx context.Context, orderId int) (Order, Plan, error)
}