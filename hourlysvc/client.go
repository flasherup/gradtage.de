package hourlysvc

import "github.com/flasherup/gradtage.de/hourlysvc/grpc"

type Client interface {
	GetPeriod(ids []string, start string, end string) *grpc.GetPeriodResponse
	PushPeriod(temps []Temperature) *grpc.PushPeriodResponse
	GetUpdateDate(ids []string) *grpc.GetUpdateDateResponse
}