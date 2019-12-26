package implementation

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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
	level.Info(ts.logger).Log("msg", "Text request", "text", text, "count", ts.count)

	return text, ts.count
}

