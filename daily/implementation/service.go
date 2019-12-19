package implementation

import (
	"context"
	"github.com/go-kit/kit/log"
)

type DailyService struct {
	logger     log.Logger
}

func NewService(logger log.Logger) *DailyService {
	return &DailyService{
		logger:logger,
	}
}

func (ds DailyService) Status(cxt context.Context) (bool, error) {
	return true,nil
}

