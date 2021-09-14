package weatherbitupdatesvc

import (
	"context"
)

type Service interface {
	ForceRestart(ctx context.Context) error
}
