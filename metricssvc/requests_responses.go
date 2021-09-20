package metricssvc

import "github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"

type GetMetricsRequest struct {
	Ids []string
}

type GetMetricsResponse struct {
	Metrics map[string]*mtrgrpc.Metrics `json:"metrics"`
	Err  	error 	`json:"err"`
}