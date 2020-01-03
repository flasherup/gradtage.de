package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/config"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/grpc"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/impl"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/impl/source"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	hourly "github.com/flasherup/gradtage.de/hourlysvc/impl"
	stations "github.com/flasherup/gradtage.de/stationssvc/impl"
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
			"svc", "hrlaggregatorssvc",
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


	hourlyService := hourly.NewHourlySCVClient(conf.Clients.HourlyAddr, logger)
	stationsService := stations.NewStationsSCVClient(conf.Clients.StationsAddr, logger)
	dataSrc := source.NewCheckWX(conf.Sources.CheckwxKey, logger)

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")



	ctx := context.Background()
	hourlyAggregatorService, err := impl.NewHrlAggregatorSVC(logger, stationsService ,hourlyService, dataSrc)
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
		grpc.RegisterHrlAggregatorSVCServer(gRPCServer, hrlaggregatorsvc.NewGRPCServer(ctx, hrlaggregatorsvc.Endpoints {
			GetStatusEndpoint:    hrlaggregatorsvc.MakeGetStatusEndpoint(hourlyAggregatorService),
		}))

		level.Info(logger).Log("transport", "GRPC", "addr", conf.GetGRPCAddress())
		errors <- gRPCServer.Serve(listener)
	}()

	metrics := hrlaggregatorsvc.NewMetricsTransport(hourlyAggregatorService,logger)

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
