package hourlysvc

import "github.com/flasherup/gradtage.de/hourlysvc/grpc"

type Client interface {
	GetPeriod(id string, start string, end string) *grpc.GetPeriodResponse
	PushPeriod(id string, temps []Temperature) *grpc.PushPeriodResponse
	GetUpdateDate(ids []string) *grpc.GetUpdateDateResponse
}