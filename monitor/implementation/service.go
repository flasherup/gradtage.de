package implementation

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
)

type LogService struct {
	logger     log.Logger
}

func NewService(logger log.Logger) *LogService {
	return &LogService{
		logger:logger,
	}
}

func (ls LogService) Log(cxt context.Context, log string) error {
	fmt.Println(log)
	return nil
}

