package main

import (
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/apisvc"
	"github.com/flasherup/gradtage.de/apisvc/config"
	"github.com/flasherup/gradtage.de/apisvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	daily "github.com/flasherup/gradtage.de/dailysvc/impl"

)

func main() {
	configFile := flag.String("config.file", "src/config.yml", "Config file name.")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "apisvc",
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

	dailyService := daily.NewDailySCVClient(conf.Clients.DailyAddr, logger)


	level.Info(logger).Log("msg", "service started", "config", configFile)
	defer level.Info(logger).Log("msg", "service ended")

	svc := impl.NewAPISVC(logger, dailyService)
	hs := apisvc.NewHTTPTSransport(svc,logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTPS", "addr", conf.GetHTTPSAddress())
		errs <- http.ListenAndServeTLS(conf.GetHTTPSAddress(), conf.Security.Cert, conf.Security.Key, hs)
	}()

	h := apisvc.NewHTTPTransport(svc,logger)

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", conf.GetHTTPAddress())
		server := &http.Server{
			Addr:    conf.GetHTTPAddress(),
			Handler: h,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)

}
