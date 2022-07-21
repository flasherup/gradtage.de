package impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/apisvc"
	"github.com/flasherup/gradtage.de/apisvc/impl/utils"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
	"github.com/flasherup/gradtage.de/metricssvc"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/usersvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"math"
	"strconv"
	"time"
)

type APISVC struct {
	logger       log.Logger
	alert        alertsvc.Client
	daydegree    daydegreesvc.Client
	weatherbit   weatherbitsvc.Client
	autocomplete autocompletesvc.Client
	user         usersvc.Client
	stations     stationssvc.Client
	metrics      metricssvc.Client
	woocommerce  *utils.Woocommerce
	counter      ktprom.Gauge
}

const (
	Hourly = "hourly"
	Daily  = "daily"
	Noaa   = "noaa"
)

func NewAPISVC(
	logger log.Logger,
	daydegree daydegreesvc.Client,
	weatherbit weatherbitsvc.Client,
	autocomplete autocompletesvc.Client,
	user usersvc.Client,
	alert alertsvc.Client,
	stations stationssvc.Client,
	metrics metricssvc.Client,
	woocommerce *utils.Woocommerce,
) *APISVC {
	options := prometheus.Opts{
		Name: "stations_count_total",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{"stations"})
	st := APISVC{
		logger:       logger,
		daydegree:    daydegree,
		weatherbit:   weatherbit,
		autocomplete: autocomplete,
		user:         user,
		alert:        alert,
		stations:     stations,
		metrics:      metrics,
		counter:      *guage,
		woocommerce:  woocommerce,
	}
	return &st
}

func (as APISVC) GetData(ctx context.Context, params []apisvc.Params) (data []*apisvc.DDResponse, err error) {
	for _, v := range params {
		dd, err := as.processDayDegree(v)
		name := utils.GetCSVName(v.Output, v.Station, v.Tb, v.Tr)
		if err != nil {
			name = "error-" + name
		}

		data = append(data, dd)
	}

	return
}

// Deprecated
func (as APISVC) GetHDD(ctx context.Context, params apisvc.Params) (data apisvc.CSVData, err error) {
	dd, err := as.processDayDegree(params)

	if err != nil {
		return
	}

	data = utils.GenerateCSV(dd.Temps, dd.Params, dd.Autocomplete)

	return
}

// Deprecated
func (as APISVC) GetHDDCSV(ctx context.Context, params apisvc.Params) (data apisvc.CSVData, fileName string, err error) {
	dd, err := as.processDayDegree(params)
	if err != nil {
		return data, "error", err
	}
	if params.Output == common.DDType {
		fileName = fmt.Sprintf("%s_DD_%gC_%gC.csv",
			params.Station,
			params.Tb,
			params.Tr)
	}
	if params.Output == common.HDDType {
		fileName = fmt.Sprintf("%s_HDD_%gC.csv",
			params.Station,
			params.Tb)
	}
	if params.Output == common.CDDType {
		fileName = fmt.Sprintf("%s_CDD_%gC.csv",
			params.Station,
			params.Tb)
	}

	data = utils.GenerateCSV(dd.Temps, dd.Params, dd.Autocomplete)

	return data, fileName, err
}

// Deprecated
func (as APISVC) GetZIP(ctx context.Context, params []apisvc.Params) (data []apisvc.CSVDataFile, fileName string, err error) {
	for _, v := range params {
		dd, err := as.processDayDegree(v)
		name := utils.GetCSVName(v.Output, v.Station, v.Tb, v.Tr)
		if err != nil {
			name = "error-" + name
		}

		d := utils.GenerateCSV(dd.Temps, dd.Params, dd.Autocomplete)

		data = append(data, apisvc.CSVDataFile{
			Name: name,
			Data: d,
		})
	}

	o := ""
	if len(params) > 0 {
		o = params[0].Output
	}

	fileName = utils.GetZIPName(o)
	return data, fileName, err
}

func (as APISVC) processDayDegree(params apisvc.Params) (data *apisvc.DDResponse, err error) {
	err = as.validateRequest(params)
	if err != nil {
		level.Error(as.logger).Log("msg", "User validation error", "err", err)
		return nil, err
	}

	autoComplete, err := as.getAutocomplete(params.Station)
	if err != nil {
		level.Error(as.logger).Log("msg", "GetHDD station id  not found", "err", err)
		return nil, err
	}

	if autoComplete.ID == "" {
		level.Error(as.logger).Log("msg", "GetHDD station id  not found", "station", params.Station)
		return nil, errors.New("station id  not found, station:" + params.Station)
	}
	params.Station = autoComplete.ID

	level.Info(as.logger).Log("msg", "GetHDD", "station", params.Station, "key", params.Key)
	ddParams := daydegreesvc.Params{
		Station:   params.Station,
		Start:     params.Start,
		End:       params.End,
		Breakdown: params.Breakdown,
		Tb:        params.Tb,
		Tr:        params.Tr,
		Output:    params.Output,
		DayCalc:   params.DayCalc,
		WeekStart: params.WeekStart,
	}

	if params.Breakdown == "" {
		ddParams.Breakdown = common.BreakdownDaily
	}

	if params.DayCalc == "" {
		ddParams.DayCalc = common.DayCalcMean
	}

	degree, err := as.daydegree.GetDegree(ddParams)
	if err != nil {
		level.Error(as.logger).Log("msg", "Get day degree data error", "err", err)
	}

	//Process Average
	if params.Avg > 0 {
		degreeAvg, err := as.daydegree.GetAverageDegree(ddParams, params.Avg)
		if err != nil {
			level.Error(as.logger).Log("msg", "Get day degree average data error", "err", err)
		}
		return &apisvc.DDResponse{
			Temps:        degree,
			Average:      degreeAvg,
			Params:       ddParams,
			Autocomplete: autoComplete,
		}, nil
	}

	/*//Process Average
	if params.Avg > 0 {
		degreeAvg, err := as.daydegree.GetAverageDegree(ddParams, params.Avg)
		if err != nil {
			level.Error(as.logger).Log("msg", "Get day degree average data error", "err", err)
		}
		return utils.GenerateAvgCSV(degree, degreeAvg, ddParams, autoComplete), nil
	}*/

	return &apisvc.DDResponse{
		Temps:        degree,
		Params:       ddParams,
		Autocomplete: autoComplete,
	}, nil

	//return utils.GenerateCSV(degree, ddParams, autoComplete), nil
}

func (as APISVC) GetSourceData(ctx context.Context, params apisvc.ParamsSourceData) (data apisvc.CSVData, fileName string, err error) {
	order, _, err := as.validateUser(params.Key)
	if err != nil {
		level.Error(as.logger).Log("msg", "Get source data error", "err", err)
		return utils.CSVError(err), "error", err
	}
	level.Info(as.logger).Log("msg", "GetSourceData", "station", params.Station, "user", order.Email, "start", params.Start, "end", params.End)

	temps, err := as.weatherbit.GetPeriod([]string{params.Station}, params.Start, params.End)
	if err != nil {
		level.Error(as.logger).Log("msg", "GetSourceData error", "err", err)
		return [][]string{}, "error", err
	}

	t := (*temps)[params.Station]

	headerCSV := []string{"ID", "Date", "Temperature"}

	csv := as.generateSourceCSV(headerCSV, t, params)
	fileName = fmt.Sprintf("source_%s_%s%s%s.csv", params.Station, params.Start, params.End)
	return csv, fileName, err
}

func (as APISVC) Search(ctx context.Context, params apisvc.ParamsSearch) (data apisvc.CSVData, err error) {
	order, _, err := as.validateUser(params.Key)
	if err != nil {
		level.Error(as.logger).Log("msg", "User validation error", "err", err)
		return utils.CSVError(err), err
	}

	level.Info(as.logger).Log("msg", "Search", "text", params.Text, "user", order.Email)
	sources, err := as.autocomplete.GetAutocomplete(params.Text)
	if err != nil {
		level.Error(as.logger).Log("msg", "Search error", "err", err)
		as.sendAlert(NewErrorAlert(err))
	}

	headerCSV := []string{"FoundIn", "ID", "ICAO", "WMO", "DWD", "Name"}
	csv := as.generateSearchCSV(headerCSV, sources)
	return csv, err
}

func (as APISVC) User(ctx context.Context, params apisvc.ParamsUser) (data apisvc.CSVData, err error) {
	order, _, err := as.validateUser(params.Key)
	if err != nil {
		level.Error(as.logger).Log("msg", "User validation error", "err", err)
		return utils.CSVError(err), err
	}

	if !order.Admin {
		err := errors.New("not enough rights to create a user")
		return utils.CSVError(err), err
	}

	level.Info(as.logger).Log("msg", "User", "action", params.Action, "key", params.Key)
	return [][]string{}, err
}

func (as APISVC) Woocommerce(ctx context.Context, event apisvc.WoocommerceEvent) (json string, err error) {
	level.Info(as.logger).Log("msg", "Woocommerce event", "Event", event.Type)
	fmt.Println("Header:", event.Header)
	fmt.Println("Body:", string(event.Body))

	if !utils.ValidateWoocommerceRequest(event.Signature, event.Body, as.woocommerce.WHSecret) {
		level.Error(as.logger).Log("msg", "Subscription update error", "err", "invalid signature")
		json = "{\"status\":\"error\"}"
		return json, err
	}

	if event.Type == common.WCUpdateEvent {
		orderId := event.UpdateEvent.ID
		productId := strconv.Itoa(event.UpdateEvent.LineItems[0].ProductID)
		email := *event.UpdateEvent.Billing.Email

		order, _, err := as.user.ValidateOrder(orderId)
		if err != nil {
			if !utils.IsSubscriptionRenewal(event.UpdateEvent.MetaData) {
				//Create new user
				key, err := as.woocommerce.GenerateAPIKey(orderId, email, productId)
				if err != nil {
					level.Error(as.logger).Log("msg", "Subscription update error", "orderId", orderId, "email", email, "productId", productId, "err", err)
				} else {
					err := CreateWoocommerceOrder(as.user, orderId, email, key, productId)
					if err != nil {
						level.Error(as.logger).Log("msg", "Subscription update error", "orderId", orderId, "email", email, "productId", productId, "err", err)
					} else {
						level.Info(as.logger).Log("msg", "Subscription update success", "orderId", orderId, "email", email, "productId", productId, "key", key)
					}
				}
			}
		} else {
			updateError := UpdateWoocommerceOrder(as.user, event.UpdateEvent.Status, email, productId, order)
			if updateError != nil {
				level.Error(as.logger).Log("msg", "Subscription update error", "email", email, "orderId", orderId, "productId", productId, "err", updateError)
			} else {
				level.Info(as.logger).Log("msg", "Subscription update success", "orderId", orderId, "email", email, "productId", productId)
			}
		}
	} else if event.Type == common.WCDeleteEvent {
		deleteError := as.user.DeleteOrder(event.DeleteEvent.ID)
		if deleteError != nil {
			level.Error(as.logger).Log("msg", "Delete order error", "orderId", event.DeleteEvent.ID, "err", err)
		}
	}

	json = "{\"status\":\"ok\"}"
	return json, err
}

func (as APISVC) Service(ctx context.Context, name string, params map[string]string) (response interface{}, err error) {
	level.Info(as.logger).Log("msg", "Service", "Name", name)

	resp := struct {
		Status   string      `json:"status"`
		Error    string      `json:"error"`
		Response interface{} `json:"response"`
	}{}

	_, _, err = as.validateUser(params["key"])
	if err != nil {
		level.Error(as.logger).Log("msg", "Run service error", "err", err)
		resp.Status = "error"
		resp.Error = err.Error()
		return resp, err
	}

	/*if !order.Admin {
		level.Error(as.logger).Log("msg", "Run command error", "err", "User validation error")
		resp.Status = "error"
		resp.Error = "Not enough rights to run commands"
		return resp, nil
	}*/

	r, err := RunService(&as, name, params)
	if err != nil {
		level.Error(as.logger).Log("msg", "Run service error", "err", err.Error())
		resp.Status = "error"
		resp.Error = err.Error()
		return resp, err
	} else {
		resp.Status = "ok"
		resp.Response = r
		return resp, nil
	}
	return "", err
}

func (as APISVC) getStationID(text string) (string, error) {
	sources, err := as.autocomplete.GetAutocomplete(text)
	if err != nil {
		as.sendAlert(NewErrorAlert(err))
		return "", err
	}

	for _, v := range sources {
		if len(v) == 0 {
			return "", errors.New("stations not found")
		}
		return v[0].ID, nil
	}

	return text, nil
}

func (as APISVC) getAutocomplete(text string) (autocompletesvc.Autocomplete, error) {
	sources, err := as.autocomplete.GetAutocomplete(text)
	if err != nil {
		as.sendAlert(NewErrorAlert(err))
		return autocompletesvc.Autocomplete{}, err
	}

	for _, v := range sources {
		if len(v) == 0 {
			return autocompletesvc.Autocomplete{}, errors.New("stations not found")
		}
		return v[0], nil
	}

	return autocompletesvc.Autocomplete{}, nil
}

func (as APISVC) generateSourceCSV(names []string, temps []common.Temperature, params apisvc.ParamsSourceData) [][]string {
	res := [][]string{names}
	var line []string
	for _, v := range temps {
		line = []string{
			params.Station,
			v.Date,
			fmt.Sprintf("%.1f", v.Temp),
		}

		res = append(res, line)
	}
	return res
}

func (as APISVC) generateSearchCSV(names []string, sources map[string][]autocompletesvc.Autocomplete) [][]string {
	res := [][]string{names}
	var line []string
	for k, v := range sources {
		for _, s := range v {
			line = []string{
				k,
				s.ID,
				s.SourceID,
				strconv.FormatFloat(s.Latitude, 'E', -1, 64),
				strconv.FormatFloat(s.Longitude, 'E', -1, 64),
				s.Source,
				s.Reports,
				s.ISO2Country,
				s.ISO3Country,
				s.Prio,
				s.CityNameEnglish,
				s.CityNameNative,
				s.CountryNameEnglish,
				s.CountryNameNative,
				s.ICAO,
				s.WMO,
				s.CWOP,
				s.Maslib,
				s.National_ID,
				s.IATA,
				s.USAF_WBAN,
				s.GHCN,
				s.NWSLI,
				strconv.FormatFloat(s.Elevation, 'E', -1, 64),
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
	if value >= baseHDD || value >= baseDD {
		return 0
	}

	return baseDD - value
}

func calculateCDD(baseCDD float64, value float64) float64 {
	if value < baseCDD {
		return 0
	}
	return value - baseCDD
}

func (as APISVC) sendAlert(alert alertsvc.Alert) {
	err := as.alert.SendAlert(alert)
	if err != nil {
		level.Error(as.logger).Log("msg", "SendAlert Alert Error", "err", err)
	}
}

func (as APISVC) validateUser(key string) (usersvc.Order, usersvc.Plan, error) {
	return as.user.ValidateKey(key)
}

func (as APISVC) validateRequest(params apisvc.Params) error {
	start, err := time.Parse(common.TimeLayoutWBH, params.Start)
	if err != nil {
		level.Error(as.logger).Log("msg", "Start time validation error", "err", err)
		as.sendAlert(NewErrorAlert(err))
		return err
	}

	end, err := time.Parse(common.TimeLayoutWBH, params.End)
	if err != nil {
		level.Error(as.logger).Log("msg", "End time validation error", "err", err)
		as.sendAlert(NewErrorAlert(err))
		return err
	}

	selection := usersvc.Selection{
		Key:       params.Key,
		StationID: params.Station,
		Method:    params.Output,
		Start:     start,
		End:       end,
	}

	return as.user.ValidateSelection(selection)
}

func getLeapSafeDOY(t time.Time) int {
	doy := t.YearDay()
	if isLeap(t.Year()) && doy >= 60 {
		return doy - 1
	}

	return doy
}

func isLeap(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}
