package usersvc

type CreateUserRequest struct {
	UserName    string 	`json:"user_name"`
	Plan		string 	`json:"plan"`
	Email 		bool 	`json:"email"`
}

type CreateUserResponse struct {
	Key 	string `json:"key"`
	Err    	error `json:"err"`
}

type UpdateUserRequest struct {
	User	    User 	`json:"user"`
	Email 		bool 	`json:"email"`
}

type UpdateUserResponse struct {
	Key 	string `json:"key"`
	Err    	error `json:"err"`
}

type AddPlanRequest struct {
	Plan		Plan 	`json:"plan"`
}

type AddPlanResponse struct {
	Err    error `json:"err"`
}

type ValidateKeyRequest struct {
	Key    string 		`json:"key"`
}

type ValidateKeyResponse struct {
	Parameters 	Parameters 	`json:"parameters"`
	Err    		error 		`json:"err"`
}

type ValidateNameRequest struct {
	Name    string 		`json:"name"`
}

type ValidateNameResponse struct {
	Parameters 	Parameters 	`json:"parameters"`
	Err    		error 		`json:"err"`
}