package testsvc

type TestRequest struct {
	Text string `json:"text"`
}

type TestResponse struct {
	Text  string `json:"text"`
	Count int `json:"count"`
}