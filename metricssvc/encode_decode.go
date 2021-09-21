package metricssvc

import (
	"context"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/metricssvc/mtrgrpc"
)

func EncodeGetMetricsResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(GetMetricsResponse)
	return &mtrgrpc.GetMetricsResponse {
		Metrics: res.Metrics,
		Err: common.ErrorToString(res.Err),
	}, nil
}

func DecodeGetMetricsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*mtrgrpc.GetMetricsRequest)
	return GetMetricsRequest{req.Ids}, nil
}