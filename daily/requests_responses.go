package daily

type StatusRequest struct {

}

type StatusResponse struct {
	Status bool `json:"status"`
	Err    error `json:"error,omitempty"`
}