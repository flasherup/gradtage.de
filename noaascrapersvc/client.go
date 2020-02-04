package noaascrapersvc

import (
	"github.com/flasherup/gradtage.de/noaascrapersvc/noaascpc"
)

type Client interface {
	ForceOverrideHourly() (resp *noaascpc.NoaaScraperSVCClient, err error)
}