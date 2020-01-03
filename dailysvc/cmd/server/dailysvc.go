package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dailysvc/config"
	"github.com/flasherup/gradtage.de/dailysvc/grpc"
	"github.com/flasherup/gradtage.de/dailysvc/impl"
	"github.com/flasherup/gradtage.de/dailysvc/impl/database"
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
			"svc", "hourlysvc",
			"ts", log.DefaultTimestampUTC,
			"caller", log.Caller(3),
		)
	}

	//Config
	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		level.Error(logger).Log("msg", "config loading error", "exit", err.Error())
		return
	}

	//Database
	db, err := database.NewPostgres(conf.Database)
	if err != nil {
		level.Error(logger).Log("msg", "database error", "exit", err.Error())
		return
	}


	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")



	ctx := context.Background()
	hourlyService, err := impl.NewDailySVC(logger, db)
	if err != nil {
		level.Error(logger).Log("msg", "service error", "exit", err.Error())
		return
	}
	errors := make(chan error)

	go func() {
		listener, err := net.Listen("tcp", conf.GetGRPCAddress())
		if err != nil {
			errors <- err
			return
		}

		gRPCServer := googlerpc.NewServer()
		grpc.RegisterDailySVCServer(gRPCServer, dailysvc.NewGRPCServer(ctx, dailysvc.Endpoints {
			GetPeriodEndpoint:    	dailysvc.MakeGetPeriodEndpoint(hourlyService),
			PushPeriodEndpoint: 	dailysvc.MakePushPeriodEndpoint(hourlyService),
			GetUpdateDateEndpoint: 	dailysvc.MakeGetUpdateDateEndpoint(hourlyService),
		}))

		level.Info(logger).Log("transport", "GRPC", "addr", conf.GetGRPCAddress())
		errors <- gRPCServer.Serve(listener)
	}()

	metrics := dailysvc.NewMetricsTransport(hourlyService,logger)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errors <- fmt.Errorf("%s", <-c)
	}()

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
