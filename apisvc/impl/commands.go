package impl

var CommandUpdateAverage = "update_average"


func ParseCommand(apiService APISVC, name string, params map[string]string) (interface{}, error){
	resp := struct {
		Command string `json:"command"`
		Error string `json:"error"`
	}{
		name,
		"not found",
	}
	return resp, nil
}