package stationssvc

import (
	"context"
)

type Service interface {
	GetAutocomplete(ctx context.Context, text string) (result map[string]string, err error)
}