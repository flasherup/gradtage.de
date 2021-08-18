package daydegreesvc

type Client interface {
	GetDegree(params Params) ([]Degree,error)
}
