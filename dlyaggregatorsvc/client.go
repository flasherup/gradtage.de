package dlyaggregatorsvc

import "github.com/flasherup/gradtage.de/dlyaggregatorsvc/grpc"

type Client interface {
	GetStatus() *grpc.GetStatusResponse

}