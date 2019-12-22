package implementation

import (
	"context"
	"github.com/go-kit/kit/log"
)

type TestService struct {
	logger     log.Logger
	count      int
}

func NewService(logger log.Logger) *TestService {
	return &TestService{
		logger:logger,
	}
}

func (ts *TestService) Text(cxt context.Context, text string) (string,int) {
	ts.count++
	return text, ts.count
}

