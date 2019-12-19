package monitor

type LogRequest struct {
	Log string `json:"log"`
}

type LogResponse struct {
	Err    error `json:"error,omitempty"`
}