package implementation

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type TestService struct {
	logger     log.Logger
	count      int
	prom 	   prometheus.Counter
}

func NewService(logger log.Logger) *TestService {
	promCounter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "test_request_count_total",
		Help: "The total number of requests",
	})
	return &TestService{
		logger: logger,
		count: 0,
		prom:  promCounter,
	}
}

func (ts *TestService) Text(cxt context.Context, text string) (string,int) {
	ts.prom.Inc()
	ts.count++
	return text, ts.count
}

