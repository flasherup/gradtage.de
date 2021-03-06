package impl

import (
	"context"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/autocompletesvc/acrpc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	googlerpc "google.golang.org/grpc"
)

type AutocompleteSVCClient struct{
	logger     log.Logger
	host string
}

func NewAutocompleteSCVClient(host string, logger log.Logger) *AutocompleteSVCClient {
	return &AutocompleteSVCClient{
		logger:logger,
		host: host,
	}
}

func (acc AutocompleteSVCClient) GetAutocomplete(text string) (map[string][]autocompletesvc.Source , error) {
	conn := acc.openConn()
	defer conn.Close()

	client := acrpc.NewAutocompleteSVCClient(conn)
	resp,err := client.GetAutocomplete(context.Background(), &acrpc.GetAutocompleteRequest{Text:text})
	if err != nil {
		level.Error(acc.logger).Log("msg", "Failed to get stations", "err", err)
		return nil, err
	}
	res := autocompletesvc.DecodeSourcesMap(resp.Result)
	return res, common.ErrorFromString(resp.Err)
}


func (acc AutocompleteSVCClient) AddSources(source []autocompletesvc.Source) error {
	conn := acc.openConn()
	defer conn.Close()

	client := acrpc.NewAutocompleteSVCClient(conn)
	src := autocompletesvc.EncodeSources(source)
	resp,err := client.AddSources(context.Background(), &acrpc.AddSourcesRequest{Sources:src})
	if err != nil {
		level.Error(acc.logger).Log("msg", "Failed to add sources", "err", err)
		return err
	}
	return common.ErrorFromString(resp.Err)
}

func (acc AutocompleteSVCClient) ResetSources(source []autocompletesvc.Source) error {
	conn := acc.openConn()
	defer conn.Close()

	client := acrpc.NewAutocompleteSVCClient(conn)
	src := autocompletesvc.EncodeSources(source)
	resp,err := client.ResetSources(context.Background(), &acrpc.ResetSourcesRequest{Sources:src})
	if err != nil {
		level.Error(acc.logger).Log("msg", "Failed to reset sources", "err", err)
		return err
	}
	return common.ErrorFromString(resp.Err)
}

func (acc AutocompleteSVCClient) openConn() *googlerpc.ClientConn {
	cc, err := googlerpc.Dial(acc.host, googlerpc.WithInsecure())
	if err != nil {
		level.Error(acc.logger).Log("msg", "Failed to start gRPC connection", "err", err)
	}
	return cc
}
