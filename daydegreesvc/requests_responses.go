package daydegreesvc

type GetDegreeRequest struct {
	Params Params `json:"params"`
}

type GetDegreeResponse struct {
	Degrees []Degree `json:"temps"`
	Err     error    `json:"err"`
}

type GetAverageDegreeRequest struct {
	Params Params `json:"params"`
	Years int	  `json:"years"`
}

type GetAverageDegreeResponse struct {
	Degrees []Degree `json:"temps"`
	Err     error    `json:"err"`
}
