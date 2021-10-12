package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	stations "github.com/flasherup/gradtage.de/stationssvc/impl"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/config"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/impl"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/impl/database"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/wbugrpc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	alert "github.com/flasherup/gradtage.de/alertsvc/impl"
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
			"svc", "weatherbitsvc",
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

	var alertService alertsvc.Client
	if conf.AlertsEnable {
		alertService = alert.NewAlertSCVClient(conf.Clients.AlertAddr, logger)
	} else {
		alertService = common.NewSilentAlert()
	}

	stationsService := stations.NewStationsSCVClient(conf.Clients.StationsAddr, logger)
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	alertService.SendAlert(impl.NewNotificationAlert("service started"))



	ctx := context.Background()
	weatherBitUpdateService, err := impl.NewWeatherBitUpdateSVC(logger, stationsService , db, alertService, *conf, config.WeatherbitConfig{})
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

		gRPCServer := googlerpc.NewServer(
				googlerpc.MaxRecvMsgSize(common.MaxMessageReceiveSize ),
				googlerpc.MaxSendMsgSize(common.MaxMessageSendSize ),
			)
		endpoints := weatherbitupdatesvc.MakeServerEndpoints(weatherBitUpdateService)
		wbugrpc.RegisterWeatherBitUpdateSVCServer(gRPCServer, weatherbitupdatesvc.NewGRPCServer(ctx, endpoints))

		level.Info(logger).Log("transport", "GRPC", "addr", conf.GetGRPCAddress())
		errors <- gRPCServer.Serve(listener)
	}()

	metrics := weatherbitupdatesvc.NewMetricsTransport(weatherBitUpdateService,logger)

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
	alertService.SendAlert(impl.NewNotificationAlert("service stopped"))
}
