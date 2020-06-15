package usersvc

type Client interface {
	CreateUser(userName string, plan string, email bool) (string, error)
	UpdateUser(user User, email bool) (string, error)
	AddPlan(plan Plan) error
	ValidateKey(key string) (Parameters, error)
	ValidateName(name string) (Parameters, error)
	ValidateStripe(stripe string) (Parameters, error)
}