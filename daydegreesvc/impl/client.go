package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
	"github.com/flasherup/gradtage.de/daydegreesvc/ddgrpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	googlerpc "google.golang.org/grpc"
)

type DayDegreeSVCClient struct{
	logger     log.Logger
	host string
}

func NewDayDegreeSVCClient(host string, logger log.Logger) *DayDegreeSVCClient {
	logger = log.With(logger,
		"client", "daydegree",
	)
	return &DayDegreeSVCClient{
		logger:logger,
		host: host,
	}
}

func (ddc * DayDegreeSVCClient) GetDegree(params daydegreesvc.Params) (temps []daydegreesvc.Degree, err error) {
	conn := ddc.openConn()
	defer conn.Close()

	p := daydegreesvc.ToGRPCParams(&params)

	client := ddgrpc.NewDayDegreeSVCClient(conn)
	grpc, err := client.GetDegree(context.Background(), &ddgrpc.GetDegreeRequest{Params: p})
	if err != nil {
		level.Error(ddc.logger).Log("msg", "Failed to get degree", "err", err)
	} else if grpc.Err != common.ErrorNilString {
		err = common.ErrorFromString(grpc.Err)
	} else {
		temps = *(daydegreesvc.ToDegree(&grpc.Degrees))
	}
	return temps, err
}


func (ddc * DayDegreeSVCClient) GetAverageDegree(params daydegreesvc.Params, years int) (temps []daydegreesvc.Degree, err error) {
		conn := ddc.openConn()
		defer conn.Close()

		p := daydegreesvc.ToGRPCParams(&params)
		client := ddgrpc.NewDayDegreeSVCClient(conn)
		grpc, err := client.GetAverageDegree(context.Background(), &ddgrpc.GetAverageDegreeRequest{ Params: p, Years:int32(years) })
		if err != nil {
			level.Error(ddc.logger).Log("msg", "Failed to GetAverage", "err", err)
		} else if grpc.Err != common.ErrorNilString {
			err = common.ErrorFromString(grpc.Err)
		} else {
			temps = *(daydegreesvc.ToDegree(&grpc.Degrees))
		}
		return temps, err
	}

func (ddc * DayDegreeSVCClient) openConn() *googlerpc.ClientConn {
	options := googlerpc.WithDefaultCallOptions(
		googlerpc.MaxCallRecvMsgSize(common.MaxMessageReceiveSize),
		googlerpc.MaxCallSendMsgSize(common.MaxMessageSendSize),
	)
	cc, err := googlerpc.Dial(ddc.host, googlerpc.WithInsecure(), options)
	if err != nil {
		level.Error(ddc.logger).Log("msg", "Failed to start gRPC connection", "err", err)
	}
	return cc
}
