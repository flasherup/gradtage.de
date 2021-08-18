package impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/alertsvc"
	"github.com/flasherup/gradtage.de/apisvc"
	"github.com/flasherup/gradtage.de/apisvc/impl/security"
	"github.com/flasherup/gradtage.de/apisvc/impl/utils"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/dailysvc"
	"github.com/flasherup/gradtage.de/dailysvc/dlygrpc"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/noaascrapersvc"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/usersvc"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	ktprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"math"
	"strconv"
	"time"
)

type APISVC struct {
	logger  			log.Logger
	alert 				alertsvc.Client
	daily				dailysvc.Client
	hourly				hourlysvc.Client
	noaa 				noaascrapersvc.Client
	autocomplete 		autocompletesvc.Client
	user 				usersvc.Client
	stations			stationssvc.Client
	keyManager 			*security.KeyManager
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
		daily 			dailysvc.Client,
		hourly 			hourlysvc.Client,
		noaa 			noaascrapersvc.Client,
		autocomplete 	autocompletesvc.Client,
		user 			usersvc.Client,
		alert 			alertsvc.Client,
		stations    	stationssvc.Client,
		keyManager 		*security.KeyManager,
		woocommerce 	*utils.Woocommerce,
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
		stations:		stations,
		keyManager: 	keyManager,
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
	if params.Output == common.DDType {
		headerCSV = []string{ "ID", "Date", "DD", "DDAverage" }
	} else if  params.Output ==  common.HDDType {
		headerCSV = []string{ "ID","Date","HDD","HDDAverage" }
	} else if  params.Output ==  common.CDDType {
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
	order, _, err := as.validateUser(params.Key)
	if err != nil {
		level.Error(as.logger).Log("msg", "Get source data error", "err", err)
		return utils.CSVError(err), "error", err
	}
	level.Info(as.logger).Log("msg", "GetSourceData", "station", params.Station, "user", order.Email, "start", params.Start, "end", params.End, "type", params.Type)

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

func (as APISVC) Plan(ctx context.Context, params apisvc.ParamsPlan) (data [][]string, err error) {
	order, _, err := as.validateUser(params.Key)
	if err != nil {
		level.Error(as.logger).Log("msg", "User validation error", "err", err)
		return utils.CSVError(err), err
	}

	if !order.Admin {
		err := errors.New("not enough rights to create user")
		return utils.CSVError(err), err
	}

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

		if params.Output 		== common.HDDType {
			degree 	= calculateHDD(params.TB, v.Temperature)
			degreeA = calculateHDD(params.TB, aTemperature)
		} else if params.Output == common.DDType {
			degree 	= calculateDD(params.TB, params.TR, v.Temperature)
			degreeA = calculateDD(params.TB, params.TR, aTemperature)
		} else if params.Output == common.CDDType {
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


func (as APISVC)validateUser(key string) (usersvc.Order, usersvc.Plan, error) {
	_, err := as.keyManager.KeyGuard.APIKeyValid([]byte(key))
	if err == nil {
		/*user := "admin@gradtage.de"
		params, err := as.user.ValidateName(user)
		if err != nil {
			level.Error(as.logger).Log("msg", "User validation name is invalid", "err", err)
			return nil, err
		}
		return &params, nil*/
		return usersvc.Order{
			Email: "admin@gradtage.de",
			Admin: true,
		}, usersvc.Plan{}, nil
	}
	return as.user.ValidateKey(key)
}

func (as APISVC)validateRequest(params apisvc.Params) error {
	_, err := as.keyManager.KeyGuard.APIKeyValid([]byte(params.Key))
	if err == nil {
		return nil
	}

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



