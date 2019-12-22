package testsvc

import (
	"context"
	"errors"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type Service interface {
	Text(context.Context, string) (string, int)
}