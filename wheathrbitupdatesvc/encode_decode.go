package weatherbitsvc

import (
	"context"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/wheathrbitupdatesvc/wbugrpc"
)

func EncodeForceRestartResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(ForceRestartResponse)
	return &wbugrpc.ForceRestartResponse {
		Err: common.ErrorToString(res.Err),
	}, nil
}

func DecodeForceRestartRequest(_ context.Context, r interface{}) (interface{}, error) {
	return ForceRestartRequest{}, nil
}