package dailysvc

import "github.com/flasherup/gradtage.de/dailysvc/grpc"

type Client interface {
	GetPeriod(id string, start string, end string) *grpc.GetPeriodResponse
	PushPeriod(id string, temps []Temperature) *grpc.PushPeriodResponse
	GetUpdateDate(ids []string) *grpc.GetUpdateDateResponse
	UpdateAvgForYear(id string) *grpc.UpdateAvgForYearResponse
	UpdateAvgForDOY(id string, doy int) *grpc.UpdateAvgForDOYResponse
	GetAvg(id string) *grpc.GetAvgResponse
}

