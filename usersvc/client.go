package usersvc

type Client interface {
	CreateOrder(orderId int, email, plan, key string) (string, error)
	UpdateOrder(order Order) (string, error)
	DeleteOrder(orderId int) error
	AddPlan(plan Plan) error
	ValidateSelection(selection Selection) error
	ValidateKey(key string) (Order, Plan, error)
	ValidateOrder(orderId int) (Order, Plan, error)
}