package usersvc

type CreateOrderRequest struct {
	OrderId		int 	`json:"order_id"`
	Email    	string 	`json:"email"`
	Plan		string 	`json:"plan"`
	Key			string 	`json:"key"`
}

type CreateOrderResponse struct {
	Key 	string `json:"key"`
	Err    	error `json:"err"`
}

type UpdateOrderRequest struct {
	Order	 Order 	`json:"order"`
}

type UpdateOrderResponse struct {
	Key 	string `json:"key"`
	Err    	error `json:"err"`
}

type DeleteOrderRequest struct {
	OrderId	    int 	`json:"order_id"`
}

type DeleteOrderResponse struct {
	Err    	error `json:"err"`
}

type AddPlanRequest struct {
	Plan		Plan 	`json:"plan"`
}

type AddPlanResponse struct {
	Err    error `json:"err"`
}

type ValidateSelectionRequest struct {
	Selection   Selection 	`json:"selection"`
}

type ValidateSelectionResponse struct {
	Err    		error 		`json:"err"`
}

type ValidateKeyRequest struct {
	Key    string 		`json:"key"`
}

type ValidateKeyResponse struct {
	Order		Order 		`json:"order"`
	Plan		Plan 		`json:"plan"`
	Err    		error 		`json:"err"`
}

type ValidateOrderRequest struct {
	OrderId    int 		`json:"order_id"`
}

type ValidateOrderResponse struct {
	Order		Order 		`json:"order"`
	Plan		Plan 		`json:"plan"`
	Err    		error 		`json:"err"`
}