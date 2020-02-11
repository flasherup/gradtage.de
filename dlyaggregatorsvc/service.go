package dlyaggregatorsvc

import (
	"context"
)

type Service interface {
	ForceUpdate(ctx context.Context, ids []string, start string, end string) error
}