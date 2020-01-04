package impl

import (
	"context"
	"errors"
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dailysvc/dlygrpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	googlerpc "google.golang.org/grpc"
)

type DailySVCClient struct{
	logger     log.Logger
	host string
}

func NewDailySCVClient(host string, logger log.Logger) *DailySVCClient {
	return &DailySVCClient{
		logger:logger,
		host: host,
	}
}

func (scc DailySVCClient) GetPeriod(id string, start string, end string) (resp *dlygrpc.GetPeriodResponse, err error)   {
	conn := scc.openConn()
	defer conn.Close()

	client := dlygrpc.NewDailySVCClient(conn)
	resp, err = client.GetPeriod(context.Background(), &dlygrpc.GetPeriodRequest{ Id: id, Start:start, End:end })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get period", "err", err)
	}else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}
	return resp, err
}

func (scc DailySVCClient) PushPeriod(id string, temps []dailysvc.Temperature) (resp *dlygrpc.PushPeriodResponse, err error)  {
	conn := scc.openConn()
	defer conn.Close()

	client := dlygrpc.NewDailySVCClient(conn)
	tGRPC := toGRPCTemps(temps)
	resp, err = client.PushPeriod(context.Background(), &dlygrpc.PushPeriodRequest{Id: id, Temps:tGRPC})
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to push period", "err", err)
	}else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}
	return resp, err
}

func (scc DailySVCClient) GetUpdateDate(ids []string) (resp *dlygrpc.GetUpdateDateResponse, err error)  {
	conn := scc.openConn()
	defer conn.Close()

	client := dlygrpc.NewDailySVCClient(conn)
	resp, err = client.GetUpdateDate(context.Background(), &dlygrpc.GetUpdateDateRequest{ Ids: ids })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get update date", "err", err)
	}else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}
	return resp, err
}

func (scc DailySVCClient) UpdateAvgForYear(id string) (resp *dlygrpc.UpdateAvgForYearResponse, err error)  {
	conn := scc.openConn()
	defer conn.Close()

	client := dlygrpc.NewDailySVCClient(conn)
	resp, err = client.UpdateAvgForYear(context.Background(), &dlygrpc.UpdateAvgForYearRequest{ Id: id })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to update averages for year", "err", err)
	}else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}
	return resp, err
}

func (scc DailySVCClient) UpdateAvgForDOY(id string, doy int) (resp *dlygrpc.UpdateAvgForDOYResponse, err error)  {
	conn := scc.openConn()
	defer conn.Close()

	client := dlygrpc.NewDailySVCClient(conn)
	resp, err = client.UpdateAvgForDOY(context.Background(), &dlygrpc.UpdateAvgForDOYRequest{ Id: id, Doy:int32(doy) })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to update averages for DOY", "err", err)
	}else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}
	return resp, err
}

func (scc DailySVCClient) GetAvg(id string) (resp *dlygrpc.GetAvgResponse, err error) {
	conn := scc.openConn()
	defer conn.Close()

	client := dlygrpc.NewDailySVCClient(conn)
	resp, err = client.GetAvg(context.Background(), &dlygrpc.GetAvgRequest{ Id: id })
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to get average", "err", err)
	}else if resp.Err != "nil" {
		err = errors.New(resp.Err)
	}
	return resp, err
}

func (scc DailySVCClient) openConn() *googlerpc.ClientConn {
	cc, err := googlerpc.Dial(scc.host, googlerpc.WithInsecure())
	if err != nil {
		level.Error(scc.logger).Log("msg", "Failed to start gRPC connection", "err", err)
	}
	return cc
}

func toGRPCTemps(sts []dailysvc.Temperature) []*dlygrpc.Temperature {
	res := make([]*dlygrpc.Temperature, len(sts))
	for i,v := range sts {
		res[i] = &dlygrpc.Temperature{
			Date:			v.Date,
			Temperature:	float32(v.Temperature),
		}
	}

	return res
}
