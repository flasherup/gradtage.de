package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/alertsvc/altgrpc"
	"github.com/flasherup/gradtage.de/alertsvc/config"
	"github.com/flasherup/gradtage.de/alertsvc/impl"
	"github.com/flasherup/gradtage.de/alertsvc/impl/alertsys"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	googlerpc "google.golang.org/grpc"
)

func main() {
	configFile := flag.String("config.file", "src/config.yml", "Config file name.")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "altssvc",
			"ts", log.DefaultTimestampUTC,
			"caller", log.Caller(3),
		)
	}

	//Config
	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		level.Error(logger).Log("msg", "config loading error", "exit", err.Error())
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	ctx := context.Background()
	alertSys := alertsys.NewEmailAlertSystem(conf.EmailConfig)
	alertService, err := impl.NewAlertSVC(logger, alertSys)
	if err != nil {
		level.Error(logger).Log("msg", "service error", "exit", err.Error())
		return
	}
	errors := make(chan error)

	//GRPC Server
	go func() {
		level.Info(logger).Log("transport", "GRPC", "addr", conf.GetGRPCAddress())
		listener, err := net.Listen("tcp", conf.GetGRPCAddress())
		if err != nil {
			errors <- err
			return
		}
		gRPCServer := googlerpc.NewServer()
		endpoints := alertsvc.MakeServerEndpoints(alertService)
		altgrpc.RegisterAlertSVCServer(gRPCServer, alertsvc.NewGRPCServer(ctx, endpoints))
		errors <- gRPCServer.Serve(listener)
	}()

	metrics := alertsvc.NewMetricsTransport(alertService,logger)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errors <- fmt.Errorf("%s", <-c)
	}()

	//HTTP Server
	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", conf.GetHTTPAddress())
		server := &http.Server{
			Addr:    conf.GetHTTPAddress(),
			Handler: metrics,
		}
		errors <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errors)
}
