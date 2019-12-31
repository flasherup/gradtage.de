package hrlaggregatorsvc

import "github.com/flasherup/gradtage.de/hourlysvc/grpc"

type Client interface {
	GetStatus() *grpc.GetPeriodResponse

}