package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
	"github.com/flasherup/gradtage.de/daydegreesvc/config"
	"github.com/flasherup/gradtage.de/daydegreesvc/ddgrpc"
	"github.com/flasherup/gradtage.de/daydegreesvc/impl"
	weatherbit "github.com/flasherup/gradtage.de/weatherbitsvc/impl"
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
			"svc", "daydegreesvc",
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


	var alertService alertsvc.Client
	if conf.AlertsEnable {
		alertService = alert.NewAlertSCVClient(conf.Clients.AlertAddr, logger)
	} else {
		alertService = common.NewSilentAlert()
	}

	weatherbitService := weatherbit.NewWeatherBitSVCClient(conf.Clients.WeatherBitAddr, logger)
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	alertService.SendAlert(impl.NewNotificationAlert("service started"))



	ctx := context.Background()
	dayDegreeService, err := impl.NewDayDegreeSVC(logger, weatherbitService, alertService, *conf)
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
		endpoints := daydegreesvc.MakeServerEndpoints(dayDegreeService)
		ddgrpc.RegisterDayDegreeSVCServer(gRPCServer, daydegreesvc.NewGRPCServer(ctx, endpoints))

		level.Info(logger).Log("transport", "GRPC", "addr", conf.GetGRPCAddress())
		errors <- gRPCServer.Serve(listener)
	}()

	metrics := daydegreesvc.NewMetricsTransport(dayDegreeService,logger)

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
