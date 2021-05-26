package usersvc

type Client interface {
	CreateUser(userName string, plan string, key string, email bool) (string, error)
	UpdateUser(user User, email bool) (string, error)
	DeleteUser(user User)  error
	AddPlan(plan Plan) error
	ValidateKey(key string) (Parameters, error)
	ValidateName(name string) (Parameters, error)
}