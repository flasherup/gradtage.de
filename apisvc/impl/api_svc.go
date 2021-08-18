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
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/usersvc"
	"github.com/flasherup/gradtage.de/weatherbitsvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"math"
	"strconv"
	"strings"
	"time"
)

type APISVC struct {
	logger  			log.Logger
	alert 				alertsvc.Client
	daydegree			daydegreesvc.Client
	weatherbit			weatherbitsvc.Client
	autocomplete 		autocompletesvc.Client
	user 				usersvc.Client
	stations			stationssvc.Client
	woocommerce			*utils.Woocommerce
	counter 			ktprom.Gauge
}

const (
	Hourly 	= "hourly"
	Daily 	= "daily"
	Noaa 	= "noaa"
)

func NewAPISVC(
		logger 			log.Logger,
		daydegree		daydegreesvc.Client,
		weatherbit		weatherbitsvc.Client,
		autocomplete 	autocompletesvc.Client,
		user 			usersvc.Client,
		alert 			alertsvc.Client,
		stations    	stationssvc.Client,
		woocommerce 	*utils.Woocommerce,
	) *APISVC {
	options := prometheus.Opts{
		Name: "stations_count_total",
		Help: "The total number oh stations",
	}
	guage := ktprom.NewGaugeFrom(prometheus.GaugeOpts(options), []string{ "stations" })
	st := APISVC{
		logger:  		logger,
		daydegree:	 	daydegree,
		weatherbit: 		weatherbit,
		autocomplete: 	autocomplete,
		user: 			user,
		alert:   		alert,
		stations:		stations,
		counter: 		*guage,
		woocommerce: 	woocommerce,
	}
	return &st
}

func (as APISVC) GetHDD(ctx context.Context, params apisvc.Params) (data [][]string, err error) {
	err = as.validateRequest(params)
	if err != nil {
		level.Error(as.logger).Log("msg", "User validation error", "err", err)
		return utils.CSVError(err), err
	}
	level.Info(as.logger).Log("msg", "GetHDD", "station", params.Station, "key", params.Key)
	ddParams := daydegreesvc.Params{
		Station: params.Station,
		Start: params.Start,
		End: params.End,
		Breakdown: params.Breakdown,
		Tb: params.TB,
		Tr: params.TR,
		Method: params.Output,
		DayCalc: params.DayCalc,
	}
	degree, err := as.daydegree.GetDegree(ddParams)
	if err != nil {
		level.Error(as.logger).Log("msg", "Get day degree data error", "err", err)
	}

	csv := as.generateCSV(degree, params)
	return csv, err
}

func (as APISVC) GetHDDCSV(ctx context.Context, params apisvc.Params) (data [][]string, fileName string, err error) {
	err = as.validateRequest(params)
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

	level.Info(as.logger).Log("msg", "GetHDD", "station", params.Station, "key", params.Key)
	ddParams := daydegreesvc.Params{
		Station: params.Station,
		Start: params.Start,
		End: params.End,
		Breakdown: params.Breakdown,
		Tb: params.TB,
		Tr: params.TR,
		Method: params.Output,
		DayCalc: params.DayCalc,
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
	csv := as.generateCSV(degree, params)
	if params.Output == common.DDType {
		fileName = fmt.Sprintf("%s_DD_%g°C_%g°C.csv",
			params.Station,
			params.TB,
			params.TR)
	}
	if params.Output == common.HDDType {
		fileName = fmt.Sprintf("%s_HDD_%g°C.csv",
			params.Station,
			params.TB)
	}
	if params.Output == common.CDDType {
		fileName = fmt.Sprintf("%s_CDD_%g°C.csv",
			params.Station,
			params.TB)
	}
	return csv,fileName,err
}

func (as APISVC) GetSourceData(ctx context.Context, params apisvc.ParamsSourceData) (data [][]string, fileName string, err error) {
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

	headerCSV := []string{ "ID","Date","Temperature" }

	csv := as.generateSourceCSV(headerCSV, t, params)
	fileName = fmt.Sprintf("source_%s_%s%s%s.csv", params.Station, params.Start, params.End)
	return csv,fileName,err
}

func (as APISVC) Search(ctx context.Context, params apisvc.ParamsSearch) (data [][]string, err error) {
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

	headerCSV := []string{ "FoundIn", "ID","ICAO","WMO","DWD", "Name" }
	csv := as.generateSearchCSV(headerCSV, sources)
	return csv, err
}

func (as APISVC) User(ctx context.Context, params apisvc.ParamsUser) (data [][]string, err error) {
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


	//TODO restore user creation functionality
	/*switch params.Action {
		case CrateAction:
			return CreateUser(as.user, params, false)
		case AutoCrateAction:
			return CreateUser(as.user, params, true)
		case SetPLanAction:
			return SetUserPlan(as.user, params)
		case RenewAction:
			return RenewUser(as.user, params)
	}*/
	return [][]string{}, err
}

func (as APISVC) Woocommerce(ctx context.Context, event apisvc.WoocommerceEvent) (json string, err error) {
	level.Info(as.logger).Log("msg", "Woocommerce event", "Event", event.Type)
	//level.Info(as.logger).Log("msg", "Woocommerce header", "Header", fmt.Sprintf("%v", event.Header))
	//level.Info(as.logger).Log("msg", "Woocommerce header", "Body", fmt.Sprintf("%q", event.Body))
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
		} else {
			updateError := UpdateWoocommerceOrder(as.user, event.UpdateEvent.Status, email, productId, order)
			if updateError != nil {
				level.Error(as.logger).Log("msg", "Subscription update error", "email", email, "orderId", orderId, "productId", productId,  "err", updateError)
			} else {
				level.Info(as.logger).Log("msg", "Subscription update success", "orderId", orderId, "email", email, "productId", productId)
			}
		}
	} else if event.Type == common.WCDeleteEvent {
		deleteError := as.user.DeleteOrder( event.DeleteEvent.ID)
		if deleteError != nil {
			level.Error(as.logger).Log("msg", "Delete order error", "orderId", event.DeleteEvent.ID, "err", err)
		}
	}

	json = "{\"status\":\"ok\"}"
	return json, err
}

func (as APISVC) Command(ctx context.Context, name string, params map[string]string) (response interface{}, err error) {
	level.Info(as.logger).Log("msg", "Command", "Name", name)

	resp := struct {
		Status string `json:"status"`
		Error string `json:"error"`
		Response interface{} `json:"response"`
	}{}

	order, _, err := as.validateUser(params["key"])
	if err != nil {
		level.Error(as.logger).Log("msg", "Run command error", "err", err)
		resp.Status = "error"
		resp.Error = err.Error()
		return resp, err
	}

	if !order.Admin {
		level.Error(as.logger).Log("msg", "Run command error", "err", "User validation error")
		resp.Status = "error"
		resp.Error = "Not enough rights to run commands"
		return resp, nil
	}

	r, err := ParseCommand(as, name, params)
	if err != nil {
		level.Error(as.logger).Log("msg", "Run command error", "err", err.Error())
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

	for _,v := range sources {
		if len(v) == 0 {
			return "", errors.New("stations not found")
		}
		return v[0].ID, nil
	}

	return text,nil
}

func (as APISVC)generateCSV(temps []daydegreesvc.Degree, params apisvc.Params) [][]string {
	res := [][]string{}
	res = append(res, []string{"Source:", "https://energy-data.io"})
	res = append(res, []string{"Description:", "Celsius..."})
	res = append(res, []string{"Station:", params.Station})
	res = append(res, []string{"Period:", fmt.Sprintf("%s - %s",params.Start,params.End)})
	res = append(res, []string{"Method:", params.DayCalc})
	res = append(res, []string{"Breakdown:", params.Breakdown})
	res = append(res, []string{"",""})
	if params.Output == common.DDType {
		res = append(res, []string{"Date", fmt.Sprintf("DD %g°C %g°C",params.TB, params.TR)})
	} else if  params.Output ==  common.HDDType {
		res = append(res, []string{"Date",fmt.Sprintf("HDD %g°C",params.TB)})
	} else if  params.Output ==  common.CDDType {
		res = append(res, []string{"Date",fmt.Sprintf("CDD %g°C",params.TB)})
	}

	var line []string
	for _, v := range temps {
		line = []string{
			v.Date,
			getFormattedValue(v.Temp),
		}

		res = append(res, line)
	}
	return res
}

func (as APISVC)generateSourceCSV(names []string, temps []common.Temperature, params apisvc.ParamsSourceData) [][]string {
	res := [][]string{ names }
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

func (as APISVC)generateSearchCSV(names []string, sources map[string][]autocompletesvc.Autocomplete) [][]string {
	res := [][]string{ names }
	var line []string
	for k, v := range sources {
		for _,s := range v {
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


func (as APISVC)validateUser(key string) (usersvc.Order, usersvc.Plan, error) {
	return as.user.ValidateKey(key)
}

func (as APISVC)validateRequest(params apisvc.Params) error {
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
		return doy -1
	}

	return doy
}

func isLeap(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}


func getFormattedValue(percentageValue float64) string{
	value := fmt.Sprintf("%.2f", percentageValue)
	return strings.Replace(value, ".", ",", -1)
}



