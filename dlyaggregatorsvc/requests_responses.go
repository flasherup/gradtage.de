package dlyaggregatorsvc

type ForceUpdateRequest struct {
	IDs 	[]string 	`json:"ids"`
	Start 	string 		`json:"start"`
	End   	string 		`json:"end"`
}

type ForceUpdateResponse struct {
	Err  	error 		`json:"err"`
}