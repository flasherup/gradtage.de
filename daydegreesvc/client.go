package daydegreesvc

type Client interface {
	GetDegree(params Params) ([]Degree, error)
	GetAverageDegree(params Params, years int) ([]Degree, error)
}
