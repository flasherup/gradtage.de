package impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/apisvc"
	"github.com/flasherup/gradtage.de/apisvc/impl/security"
	"github.com/flasherup/gradtage.de/apisvc/impl/stripe"
	"github.com/flasherup/gradtage.de/apisvc/impl/utils"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dailysvc/dlygrpc"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/noaascrapersvc"
	"github.com/flasherup/gradtage.de/usersvc"
	"github.com/flasherup/gradtage.de/usersvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"math"
	"time"
)

type APISVC struct {
	logger  		log.Logger
	alert 			alertsvc.Client
	daily			dailysvc.Client
	hourly			hourlysvc.Client
	noaa 			noaascrapersvc.Client
	autocomplete 	autocompletesvc.Client
	user 			usersvc.Client
	keyManager 		*security.KeyManager
	counter 		ktprom.Gauge
}

const (
	CDDType = "cdd"
	HDDType = "hdd"
	DDType  = "dd"
)

const (
	Hourly 	= "hourly"
	Daily 	= "daily"
	Noaa 	= "noaa"
)

func NewAPISVC(
		logger 			log.Logger,
		daily 			dailysvc.Client,
		hourly 			hourlysvc.Client,
		noaa 			noaascrapersvc.Client,
		autocomplete 	autocompletesvc.Client,
		user 			usersvc.Client,
		alert 			alertsvc.Client,
		keyManager 		*security.KeyManager,
	) *APISVC {
	options := prometheus.Opts{
		Name: "stations_count_total",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := APISVC{
		logger:  		logger,
		daily:	 		daily,
		hourly:			hourly,
		noaa: 			noaa,
		autocomplete: 	autocomplete,
		user: 			user,
		alert:   		alert,
		keyManager: 	keyManager,
		counter: 		*guage,
	}
	return &st
}

func (as APISVC) GetHDD(ctx context.Context, params apisvc.Params) (data [][]string, err error) {
	userName, err := as.isRequestValid(params)
	if err != nil {
		level.Error(as.logger).Log("msg", "User validation error", "err", err)
		return utils.CSVError(err), err
	}
	level.Info(as.logger).Log("msg", "GetHDD", "station", params.Station, "user", userName)
	temps, err := as.daily.GetPeriod(params.Station, params.Start, params.End)
	if err != nil {
		level.Error(as.logger).Log("msg", "GetHDD error", "err", err)
		as.sendAlert(NewErrorAlert(err))
	}

	avg, err := as.daily.GetAvg(params.Station)
	if err != nil {
		level.Error(as.logger).Log("msg", "GetHDD error", "err", err)
		as.sendAlert(NewErrorAlert(err))
	}

	headerCSV := []string{ "ID","Date","HDD","HDDAverage" }
	csv := as.generateCSV(headerCSV, temps.Temps, avg.Temps, params)
	return csv, err
}

func (as APISVC) GetHDDCSV(ctx context.Context, params apisvc.Params) (data [][]string, fileName string, err error) {
	userName, err := as.isRequestValid(params)
	if err != nil {
		level.Error(as.logger).Log("msg", "User validation error", "err", err)
		return utils.CSVError(err), "error", err
	}

	id, err := as.getStationID(params.Station)
	if err != nil {
		level.Error(as.logger).Log("msg", "GetHDD station id  not found", "err", err)
		return [][]string{}, "error", err
	}

	if id == "" {
		level.Error(as.logger).Log("msg", "GetHDD station id  not found", "station", params.Station)
		return [][]string{}, "error", errors.New("station id  not found, station:" + params.Station)
	}
	params.Station = id

	level.Info(as.logger).Log("msg", "GetHDD", "station", params.Station, "userId", userName)
	temps, err := as.daily.GetPeriod(params.Station, params.Start, params.End)
	if err != nil {
		level.Error(as.logger).Log("msg", "GetHDD error", "err", err)
		as.sendAlert(NewErrorAlert(err))
	}

	avg, err := as.daily.GetAvg(params.Station)
	if err != nil {
		level.Error(as.logger).Log("msg", "GetHDD error", "err", err)
		as.sendAlert(NewErrorAlert(err))
	}

	var headerCSV []string
	if params.Output == DDType {
		headerCSV = []string{ "ID", "Date", "DD", "DDAverage" }
	} else if  params.Output == HDDType {
		headerCSV = []string{ "ID","Date","HDD","HDDAverage" }
	} else if  params.Output == CDDType {
		headerCSV = []string{ "ID","Date","CDD","CDDAverage" }
	}

	csv := as.generateCSV(headerCSV, temps.Temps, avg.Temps, params)
	fileName = fmt.Sprintf("%s%s%s%g%g.csv",
		params.Station,
		params.Start,
		params.End,
		params.TB,
		params.TR)
	return csv,fileName,err
}

func (as APISVC) GetSourceData(ctx context.Context, params apisvc.ParamsSourceData) (data [][]string, fileName string, err error) {
	p, err := as.validateUser(params.Key)
	if err != nil {
		level.Error(as.logger).Log("msg", "User validation error", "err", err)
		return utils.CSVError(err), "error", err
	}
	level.Info(as.logger).Log("msg", "GetSourceData", "station", params.Station, "userId", p.User.Name, "start", params.Start, "end", params.End, "type", params.Type)

	var temps []hourlysvc.Temperature
	if params.Type == Daily {
		t, err := as.getDailyData(params.Station, params.Start, params.End)
		if err != nil {
			level.Error(as.logger).Log("msg", "GetSourceData error", "err", err)
			return [][]string{}, "error", err
		}

		temps = *t

	} else if params.Type == Hourly{
		t, err := as.getHourlyData(params.Station, params.Start, params.End)
		if err != nil {
			level.Error(as.logger).Log("msg", "GetSourceData error", "err", err)
			return [][]string{}, "error", err
		}

		temps = *t
	} else if params.Type == Noaa{
		t, err := as.getNoaaData(params.Station, params.Start, params.End)
		if err != nil {
			level.Error(as.logger).Log("msg", "GetSourceData error", "err", err)
			return [][]string{}, "error", err
		}

		temps = *t
	}

	headerCSV := []string{ "ID","Date","Temperature" }

	csv := as.generateSourceCSV(headerCSV, temps, params)
	fileName = fmt.Sprintf("source_%s_%s%s%s.csv",params.Type, params.Station, params.Start, params.End)
	return csv,fileName,err
}

func (as APISVC) Search(ctx context.Context, params apisvc.ParamsSearch) (data [][]string, err error) {
	p, err := as.validateUser(params.Key)
	if err != nil {
		level.Error(as.logger).Log("msg", "User validation error", "err", err)
		return utils.CSVError(err), err
	}

	level.Info(as.logger).Log("msg", "Search", "text", params.Text, "user", p.User.Name)
	sources, err := as.autocomplete.GetAutocomplete(params.Text)
	if err != nil {
		level.Error(as.logger).Log("msg", "Search error", "err", err)
		as.sendAlert(NewErrorAlert(err))
	}

	headerCSV := []string{ "FoundIn", "ID","ICAO","WMO","DWD", "Name" }
	csv := as.generateSearchCSV(headerCSV, sources)
	return csv, err
}

func (as APISVC) User(ctx context.Context, params apisvc.ParamsUser) (data [][]string, err error) {
	p, err := as.validateUser(params.Key)
	if err != nil {
		level.Error(as.logger).Log("msg", "User validation error", "err", err)
		return utils.CSVError(err), err
	}

	if !p.Plan.Admin {
		err := errors.New("not enough rights to create a user")
		return utils.CSVError(err), err
	}

	level.Info(as.logger).Log("msg", "User", "action", params.Action, "key", params.Key)


	switch params.Action {
		case CrateAction:
			return CreateUser(as.user, params, false)
		case AutoCrateAction:
			return CreateUser(as.user, params, true)
		case SetPLanAction:
			return SetUserPlan(as.user, params)
		case RenewAction:
			return RenewUser(as.user, params)
	}
	return [][]string{}, err
}

func (as APISVC) Plan(ctx context.Context, params apisvc.ParamsPlan) (data [][]string, err error) {
	p, err := as.validateUser(params.Key)
	if err != nil {
		level.Error(as.logger).Log("msg", "User validation error", "err", err)
		return utils.CSVError(err), err
	}

	if !p.Plan.Admin {
		err := errors.New("not enough rights to create user")
		return utils.CSVError(err), err
	}

	return [][]string{}, err
}

func (as APISVC) Stripe(ctx context.Context, event apisvc.StripeEvent) (json string, err error) {
	level.Info(as.logger).Log("msg", "Stripe event", "event", event.Type)
	if event.Type == stripe.InvoiceFinalize {
		invoiceFinalize, err := stripe.ParseInvoiceFinalize(event.Data.Object)
		if err != nil {
			level.Error(as.logger).Log("msg", "Invoice Finalize parse error", "err", err)
		} else {
			level.Info(as.logger).Log("msg", "Invoice Finalize parsed", "customer", invoiceFinalize.Customer, "email", invoiceFinalize.CustomerEmail)
			return ProcessUpdateStripeUser(as.user, invoiceFinalize.CustomerEmail, invoiceFinalize.Customer, usersvc.PlanStarter)
		}
	} else if event.Type == stripe.SubscriptionScheduleCanceled {
		subscriptionScheduleCanceled, err := stripe.ParseSubscriptionScheduleCanceled(event.Data.Object)
		if err != nil {
			level.Error(as.logger).Log("msg", "Subscription Schedule Canceled parse error", "err", err)
		} else {
			level.Info(as.logger).Log("msg", "Subscription Schedule Canceled parsed", "customer", subscriptionScheduleCanceled.Customer)
			return ProcessCancelStripeUser(as.user, subscriptionScheduleCanceled.Customer)
		}
	}
		json = "{\"status\":\"ok\"}"
	return json, err
}

func (as APISVC) getDailyData(id string, start string, end string) (*[]hourlysvc.Temperature, error){
	resp, err := as.daily.GetPeriod(id, start, end)
	if err != nil {
		level.Error(as.logger).Log("msg", "Get Daily Data error", "err", err)
		return nil, err
	}

	res := make([]hourlysvc.Temperature, len(resp.Temps))
	for i,v := range resp.Temps {
		res[i] = hourlysvc.Temperature{
			v.Date,
			v.Temperature,
		}
	}

	return &res, nil
}

func (as APISVC) getStationID(text string) (string, error) {
	sources, err := as.autocomplete.GetAutocomplete(text)
	if err != nil {
		as.sendAlert(NewErrorAlert(err))
		return "", err
	}

	for _,v := range sources {
		if len(v) == 0 {
			return "", errors.New("stations not found")
		}
		return v[0].ID, nil
	}

	return text,nil
}

func (as APISVC) getHourlyData(id string, start string, end string) (*[]hourlysvc.Temperature, error){
	resp, err := as.hourly.GetPeriod(id, start, end)
	if err != nil {
		level.Error(as.logger).Log("msg", "Get Hourly Data error", "err", err)
		return nil, err
	}

	res := make([]hourlysvc.Temperature, len(resp.Temps))
	for i,v := range resp.Temps {
		res[i] = hourlysvc.Temperature{
			v.Date,
			v.Temperature,
		}
	}

	return &res, nil
}

func (as APISVC) getNoaaData(id string, start string, end string) (*[]hourlysvc.Temperature, error){
	resp, err := as.noaa.GetPeriod(id, start, end)
	if err != nil {
		level.Error(as.logger).Log("msg", "Get NOAA Data error", "err", err)
		return nil, err
	}

	res := make([]hourlysvc.Temperature, len(resp.Temps))
	for i,v := range resp.Temps {
		res[i] = hourlysvc.Temperature{
			v.Date,
			v.Temperature,
		}
	}

	return &res, nil
}

func (as APISVC)generateCSV(names []string, temps []*dlygrpc.Temperature, tempsAvg map[int32]*dlygrpc.Temperature, params apisvc.Params) [][]string {
	res := [][]string{ names }
	var line []string
	var degree float64
	var degreeA float64
	for _, v := range temps {
		d, err := time.Parse(common.TimeLayout, v.Date)
		if err != nil {
			level.Error(as.logger).Log("msg", "Get " + params.Output + " generateCSV error", "err", err)
			as.sendAlert(NewErrorAlert(err))
		}
		doy := int32(getLeapSafeDOY(d))

		temp := tempsAvg[doy]
		if temp == nil {
			level.Warn(as.logger).Log("msg", "Get " + params.Output + " generateCSV, can't get Average temperature", "DOY", doy)
			continue
		}

		aTemperature := temp.Temperature

		if params.Output 		== HDDType {
			degree 	= calculateHDD(params.TB, v.Temperature)
			degreeA = calculateHDD(params.TB, aTemperature)
		} else if params.Output == DDType {
			degree 	= calculateDD(params.TB, params.TR, v.Temperature)
			degreeA = calculateDD(params.TB, params.TR, aTemperature)
		} else if params.Output == CDDType {
			degree 	= calculateCDD(params.TB, v.Temperature)
			degreeA = calculateCDD(params.TB, aTemperature)
		}

		line = []string{
			params.Station,
			v.Date,
			fmt.Sprintf("%.1f", degree),
			fmt.Sprintf("%.1f", degreeA),
		}

		res = append(res, line)
	}
	return res
}

func (as APISVC)generateSourceCSV(names []string, temps []hourlysvc.Temperature, params apisvc.ParamsSourceData) [][]string {
	res := [][]string{ names }
	var line []string
	for _, v := range temps {
		line = []string{
			params.Station,
			v.Date,
			fmt.Sprintf("%.1f", v.Temperature),
		}

		res = append(res, line)
	}
	return res
}

func (as APISVC)generateSearchCSV(names []string, sources map[string][]autocompletesvc.Source) [][]string {
	res := [][]string{ names }
	var line []string
	for k, v := range sources {
		for _,s := range v {
			line = []string{
				k,
				s.ID,
				s.Icao,
				s.Wmo,
				s.Dwd,
				s.Name,
			}
			res = append(res, line)
		}
	}
	return res
}

func ToFixed(x float64) float64 {
	unit := 10.0
	return math.Round(x*unit) / unit
}

func calculateHDD(baseHDD float64, value float64) float64 {
	if value >= baseHDD {
		return 0
	}
	return baseHDD - value
}

func calculateDD(baseHDD float64, baseDD float64, value float64) float64 {
	if value >= baseHDD || value >= baseDD{
		return 0
	}

	return baseDD - value
}


func calculateCDD(baseCDD float64, value float64) float64 {
	if value < baseCDD {
		return 0
	}
	return value-baseCDD
}


func (as APISVC)sendAlert(alert alertsvc.Alert) {
	err := as.alert.SendAlert(alert)
	if err != nil {
		level.Error(as.logger).Log("msg", "SendAlert Alert Error", "err", err)
	}
}

func (as APISVC)isRequestValid(params apisvc.Params) (string, error) {
	p, err := as.validateUser(params.Key)
	if err != nil {
		level.Error(as.logger).Log("msg", "User validation error", "err", err)
		return p.User.Name, err
	}

	err = as.validateRequest(params, p)
	if err != nil {
		level.Error(as.logger).Log("msg", "Request validation error", "err", err)
		return p.User.Name, err
	}

	return p.User.Name, nil
}

func (as APISVC)validateUser(key string) (*usersvc.Parameters, error) {
	_,err := as.keyManager.KeyGuard.APIKeyValid([]byte(key))
	if err == nil {
		user := "admin@gradtage.de"
		params, err := as.user.ValidateName(user)
		if err != nil {
			level.Error(as.logger).Log("msg", "User validation name is invalid", "err", err)
			return nil, err
		}
		return &params, nil
	}

	params, err := as.user.ValidateKey(key)
	if err != nil {
		level.Error(as.logger).Log("msg", "User validation key is invalid", "err", err)
		return nil, err
	}

	return &params, nil
}

func (as APISVC)validateRequest(request apisvc.Params, params *usersvc.Parameters) error {
	err := impl.ValidateStationId(request.Station, params)
	if err != nil {
		level.Error(as.logger).Log("msg", "Station ID is invalid", "err", err)
		err = impl.ValidateStationsCount(request.Station, params)
		if err != nil {
			level.Error(as.logger).Log("msg", "Station Count is full", "err", err)
			return err
		}

		params.User.Stations = append(params.User.Stations, request.Station)
		_, err = UpdateUser(as.user, params.User)
		if err != nil {
			level.Error(as.logger).Log("msg", "User stations update error", "err", err)
			return err
		}
		return nil
	}
	return nil
}




func getLeapSafeDOY(t time.Time) int {
	doy := t.YearDay()
	if isLeap(t.Year()) && doy >= 60 {
		return doy -1
	}

	return doy
}

func isLeap(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}

