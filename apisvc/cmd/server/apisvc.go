package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	alert "github.com/flasherup/gradtage.de/alertsvc/impl"
	"github.com/flasherup/gradtage.de/apisvc"
	"github.com/flasherup/gradtage.de/apisvc/config"
	"github.com/flasherup/gradtage.de/apisvc/impl"
	"github.com/flasherup/gradtage.de/apisvc/impl/utils"
	autocomplete "github.com/flasherup/gradtage.de/autocompletesvc/impl"
	"github.com/flasherup/gradtage.de/common"
	daydegree "github.com/flasherup/gradtage.de/daydegreesvc/impl"
	stations "github.com/flasherup/gradtage.de/stationssvc/impl"
	user "github.com/flasherup/gradtage.de/usersvc/impl"
	weatherbit "github.com/flasherup/gradtage.de/weatherbitsvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	var alertService alertsvc.Client
	if conf.AlertsEnable {
		alertService = alert.NewAlertSCVClient(conf.Clients.AlertAddr, logger)
	} else {
		alertService = common.NewSilentAlert()
	}

	daydegreeService := daydegree.NewDayDegreeSVCClient(conf.Clients.DaydegreeAddr, logger)
	weatherbitService := weatherbit.NewWeatherBitSVCClient(conf.Clients.WeatherbitAddr, logger)
	autocompleteService := autocomplete.NewAutocompleteSCVClient(conf.Clients.AutocompleteAddr, logger)
	userService := user.NewUsersSCVClient(conf.Clients.UserAddr, logger)
	stationsService := stations.NewStationsSCVClient(conf.Clients.StationsAddr, logger)

	woocommerce := utils.NewWoocommerce(conf.Woocommerce)

	level.Info(logger).Log("msg", "service started", "config", configFile)
	defer level.Info(logger).Log("msg", "service ended")

	alertService.SendAlert(impl.NewNotificationAlert("service started"))

	svc := impl.NewAPISVC(
		logger,
		daydegreeService,
		weatherbitService,
		autocompleteService,
		userService,
		alertService,
		stationsService,
		woocommerce)
	hs := apisvc.NewHTTPTSransport(svc,logger, conf.Static.Folder)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	mgr := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(conf.Domains...),
		Cache:      autocert.DirCache(conf.GetHTTPSAddress() + "cert/"), // to store certs
	}

	go func() {
		level.Info(logger).Log("transport", "Static", "addr", conf.GetHTTPStaticAddress())
		errs <- http.ListenAndServe(conf.GetHTTPStaticAddress(), mgr.HTTPHandler(nil))
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTPS", "addr", conf.GetHTTPSAddress())

		cfg := &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			},
			GetCertificate: mgr.GetCertificate,
		}

		server := &http.Server{
			Addr:    conf.GetHTTPSAddress(),
			Handler: hs,
			TLSConfig: cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}
		errs <- server.ListenAndServeTLS("", "")
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

	go func() {
		level.Info(logger).Log("transport", "HTTPTest", "addr", conf.GetHTTPTestAddress())
		server := &http.Server{
			Addr:    conf.GetHTTPTestAddress(),
			Handler: hs,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)
	alertService.SendAlert(impl.NewNotificationAlert("service stopped"))
}
