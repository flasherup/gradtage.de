package daydegreesvc

type GetDegreeRequest struct {
	Params Params `json:"params"`
}

type GetDegreeResponse struct {
	Degrees []Degree `json:"temps"`
	Err     error    `json:"err"`
}
