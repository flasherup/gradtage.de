package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/config"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/dagrpc"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/impl"
	"github.com/flasherup/gradtage.de/dlyaggregatorsvc/impl/source"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	daily "github.com/flasherup/gradtage.de/dailysvc/impl"
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
			"svc", "dlyaggregatorssvc",
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
	dailyService := daily.NewDailySCVClient(conf.Clients.DailyAddr, logger)
	stationsService := stations.NewStationsSCVClient(conf.Clients.StationsAddr, logger)
	sourceHourly := source.NewHourly(logger, hourlyService, dailyService)

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")



	ctx := context.Background()
	dailyAggregatorService, err := impl.NewDlyAggregatorSVC(logger, stationsService, dailyService, sourceHourly)
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
		dagrpc.RegisterDlyAggregatorSVCServer(gRPCServer, dlyaggregatorsvc.NewGRPCServer(ctx, dlyaggregatorsvc.Endpoints {
			GetStatusEndpoint:    dlyaggregatorsvc.MakeGetStatusEndpoint(dailyAggregatorService),
		}))

		level.Info(logger).Log("transport", "GRPC", "addr", conf.GetGRPCAddress())
		errors <- gRPCServer.Serve(listener)
	}()

	metrics := dlyaggregatorsvc.NewMetricsTransport(dailyAggregatorService,logger)

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
