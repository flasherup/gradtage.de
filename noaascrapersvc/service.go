package noaascrapersvc

import "context"

type Service interface {
	ForceOverrideHourly(ctx context.Context, station string, start string, end string) error
}