package main

import (
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/testsvc"
	"github.com/flasherup/gradtage.de/testsvc/config"
	"github.com/flasherup/gradtage.de/testsvc/implementation"
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
		configFile = flag.String("config.file", "cfg/config.yml", "Config file name.")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "testsvc",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	//Config
	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		level.Error(logger).Log("msg", "config loading error", "exit", err.Error())
		return
	}

	level.Info(logger).Log("msg", "service started", "config", configFile)
	defer level.Info(logger).Log("msg", "service ended")

	svc := implementation.NewService(logger)
	h := testsvc.NewHTTPTransport(svc,logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    conf.GetHTTPAddress(),
			Handler: h,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)

}
