package main

import (
	"context"
	"fmt"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/stationssvc/grpc"
	"github.com/flasherup/gradtage.de/stationssvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net"
	"os"

	googlerpc "google.golang.org/grpc"
)

func main() {
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

	ctx := context.Background()
	stationsService := impl.NewStationsSVC(logger)
	errors := make(chan error)

	go func() {
		listener, err := net.Listen("tcp", ":9090")
		if err != nil {
			errors <- err
			return
		}

		gRPCServer := googlerpc.NewServer()
		grpc.RegisterStationSVCServer(gRPCServer, stationssvc.NewGRPCServer(ctx, stationssvc.Endpoints {
			GetStationsEndpoint:    stationssvc.MakeGetStationsEndpoint(stationsService),
			GetAllStationsEndpoint: stationssvc.MakeGetAllStationsEndpoint(stationsService),
			AddStationsEndpoint: stationssvc.MakeAddStationsEndpoint(stationsService),
		}))

		fmt.Println("gRPC listen on 9090")
		errors <- gRPCServer.Serve(listener)
	}()

	fmt.Println(<-errors)
}
