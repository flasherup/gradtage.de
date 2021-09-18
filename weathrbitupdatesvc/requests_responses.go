package weatherbitupdatesvc

type ForceRestartRequest struct {
}

type ForceRestartResponse struct {
	Err  	error 								`json:"err"`
}