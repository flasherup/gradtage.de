package usersvc

type Client interface {
	CreateUser(userName string, plan Plan)error
	CreateUserAuto(userName string, plan Plan)error
	SetPlan(userName string, plan Plan)error
	SetStations(userName string, stations []string)error
	ValidateKey(key string) (parameters UserParameters, err error)
}