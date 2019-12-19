package main

import (
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/monitor"
	"github.com/flasherup/gradtage.de/monitor/implementation"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8001", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewJSONLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "monitor",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var svc monitor.Service
	{
		svc = implementation.NewService(logger)
		// Add service middleware here
		// Logging middleware
		//svc = middleware.LoggingMiddleware(logger)(svc)
	}

	var h http.Handler
	{
		h = monitor.NewHTTPTransport(svc,logger)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    *httpAddr,
			Handler: h,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)

}
