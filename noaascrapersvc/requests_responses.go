package noaascrapersvc

type ForceOverrideHourlyRequest struct {
	Station 	string 	`json:"station"`
	Start 		string	`json:"start"`
	End 		string	`json:"end"`
}

type ForceOverrideHourlyResponse struct {
	Err  	error 		`json:"err"`
}