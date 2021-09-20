package weatherbitupdatesvc

import "github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"

type GetMetricsRequest struct {
	Ids []string
}

type GetMetricsResponse struct {
	Metrics []*mtrgrpc.Metrics `json:"metrics"`
	Err  	error 	`json:"err"`
}