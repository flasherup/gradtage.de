package usersvc

type Client interface {
	CreateOrder(orderId string, email, plan, key string) (string, error)
	UpdateOrder(order Order) (string, error)
	DeleteOrder(orderId string) error
	AddPlan(plan Plan) error
	ValidateSelection(selection Selection) error
	ValidateKey(key string) (Order, Plan, error)
	ValidateOrder(orderId string) (Order, Plan, error)
}
