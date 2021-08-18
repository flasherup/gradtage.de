package database

import "github.com/flasherup/gradtage.de/usersvc"


const KeyLength = 36

type UserDB interface {
	GetOrderById(id int) (usersvc.Order, error)
	GetOrdersByUser(user string) ([]usersvc.Order, error)
	GetOrderByKey(key string) (usersvc.Order, error)
	DeleteOrders(orderIds []int) error
	SetOrder(order usersvc.Order) error
	CreateOrdersTable() error
	RemoveOrdersTable() error
	SetPlan(plan usersvc.Plan) error
	GetPlans(plans []string) ([]usersvc.Plan, error)
	CreatePlansTable() error
	RemovePlansTable() error
	Dispose()
}