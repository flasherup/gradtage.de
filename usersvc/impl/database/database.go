package database

import "github.com/flasherup/gradtage.de/usersvc"

type UserDB interface {
	SetUser(user usersvc.User) error
	GetUserDataByName(userName string)  (usersvc.Parameters, error)
	GetUserDataByKey(key string)  (usersvc.Parameters, error)
	SetPlan(plan usersvc.Plan) error
	GetPlan(name string) (usersvc.Plan, error)
	CreateUserTable() error
	CreatePlanTable() error
	RemoveUserTable() error
	RemovePlanTable() error
	Dispose()
}