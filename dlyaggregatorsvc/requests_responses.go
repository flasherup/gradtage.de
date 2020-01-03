package dlyaggregatorsvc

type GetStatusRequest struct {
}

type GetStatusResponse struct {
	Status 	[]Status 	`json:"status"`
	Err  	error 		`json:"err"`
}