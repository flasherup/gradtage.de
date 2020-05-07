package usersvc

type CreateUserRequest struct {
	UserName    string 	`json:"user_name"`
	Plan		Plan 	`json:"plan"`
}

type CreateUserResponse struct {
	Err    error `json:"err"`
}

type CreateUserAutoRequest struct {
	UserName    string 	`json:"user_name"`
	Plan		Plan 	`json:"plan"`
}

type CreateUserAutoResponse struct {
	Err    error `json:"err"`
}

type SetPlanRequest struct {
	UserName    string 	`json:"user_name"`
	Plan		Plan 	`json:"plan"`
}

type SetPlanResponse struct {
	Err    error `json:"err"`
}

type SetStationsRequest struct {
	UserName    string 		`json:"user_name"`
	Station		[]string 	`json:"stations"`
}

type SetStationsResponse struct {
	Err    error `json:"err"`
}

type ValidateKeyRequest struct {
	Key   string	`json:"key"`
}

type ValidateKeyResponse struct {
	Parameters  UserParameters 	`json:"err"`
	Err    		error 			`json:"err"`
}