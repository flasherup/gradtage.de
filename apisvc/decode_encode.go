package apisvc

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/apisvc/impl/utils"
	"github.com/flasherup/gradtage.de/common"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func decodeGetDataRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	prms := getParams(r, true)
	req := GetDataRequest{prms}

	return req, nil
}

func encodeGetDataResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(GetDataResponse)

	if resp.Err != nil {
		return resp.Err
	}

	if len(resp.Data) == 1 {
		dd := resp.Data[0]

		if resp.Format == common.FormatJSON {
			return writeJSON(dd, w)
		}

		return writeCSV(dd, w)
	}

	return writeZIP(resp.Data, resp.Format, w)
}

func decodeGetHDDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := GetHDDRequest{}
	if e := json.NewDecoder(r.Body).Decode(&req.Params); e != nil {
		return nil, e
	}
	vars := mux.Vars(r)
	req.Params.End = utils.WordToTime(req.Params.End)
	req.Params.Output = vars[Method]
	req.Params.DayCalc = vars[DayCalc]
	return req, nil
}

func encodeGetHDDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(GetHDDResponse)
	w.Header().Set("Content-Type", "text/csv")
	wr := csv.NewWriter(w)
	wr.Comma = ';'
	err := wr.WriteAll(resp.Data)
	wr.Flush()
	if err != nil {
		http.Error(w, "Error sending csv: "+err.Error(), http.StatusInternalServerError)
	}
	return err
}

func decodeGetSourceDataRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	r.ParseForm()

	prm := ParamsSourceData{
		Key:     r.Form.Get("key"),
		Station: r.Form.Get("station"),
		Start:   r.Form.Get("start"),
		End:     r.Form.Get("end"),
	}

	req := GetSourceDataRequest{prm}
	return req, nil
}

func encodeGetSourceDataResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(GetSourceDataResponse)
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename="+resp.FileName)
	wr := csv.NewWriter(w)
	err := wr.WriteAll(resp.Data)
	wr.Flush()
	if err != nil {
		http.Error(w, "Error sending csv: "+err.Error(), http.StatusInternalServerError)
	}
	return err
}

func decodeSearchRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	r.ParseForm()
	params := ParamsSearch{
		Key:  r.Form.Get("key"),
		Text: r.Form.Get("text"),
	}
	return SearchRequest{params}, nil
}

func encodeSearchResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(SearchResponse)
	w.Header().Set("Content-Type", "text/csv")
	wr := csv.NewWriter(w)
	err := wr.WriteAll(resp.Data)
	wr.Flush()
	if err != nil {
		http.Error(w, "Error sending csv: "+err.Error(), http.StatusInternalServerError)
	}
	return err
}

func decodeUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	r.ParseForm()
	p := map[string]string{
		"key":  r.Form.Get("key"),
		"plan": r.Form.Get("plan"),
		"name": r.Form.Get("email"),
	}

	if p["key"] == "" {
		return nil, errors.New("key is required")
	}

	params := ParamsUser{
		Key:    p["key"],
		Action: vars[UserAction],
		Params: p,
	}
	return UserRequest{params}, nil
}

func encodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(UserResponse)
	w.Header().Set("Content-Type", "text/csv")
	wr := csv.NewWriter(w)
	err := wr.WriteAll(resp.Data)
	wr.Flush()
	if err != nil {
		http.Error(w, "Error sending csv: "+err.Error(), http.StatusInternalServerError)
	}
	return err
}

func decodeWoocommerceRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := WoocommerceRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return req, err
	}
	eventType := utils.GetWoocommerceEventType(r.Header)

	event := WoocommerceEvent{}
	event.Type = eventType
	event.Signature = utils.GetWoocommerceSignature(r.Header)
	event.Body = body
	event.Header = r.Header

	if eventType == common.WCUpdateEvent {
		if e := json.Unmarshal(body, &event.UpdateEvent); e != nil {
			return req, e
		}
	}

	if eventType == common.WCDeleteEvent {
		if e := json.Unmarshal(body, &event.DeleteEvent); e != nil {
			return req, e
		}
	}

	req.Event = event
	return req, nil
}

func encodeWoocommerceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(WoocommerceResponse)
	bt := new(bytes.Buffer)
	err := json.NewEncoder(bt).Encode(resp)

	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write(bt.Bytes())
	return err
}

func decodeServiceRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	r.ParseForm()
	p := make(map[string]string)
	for k, v := range r.Form {
		p[k] = v[0]
	}
	req := ServiceRequest{}
	req.Params = p
	req.Name = vars[ServiceName]
	return req, nil
}

func encodeServiceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(ServiceResponse)
	bt := new(bytes.Buffer)
	err := json.NewEncoder(bt).Encode(resp.Json)

	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write(bt.Bytes())
	return err
}

func getParams(r *http.Request, single bool) []Params {
	vars := mux.Vars(r)
	r.ParseForm()
	basehddStr := r.Form.Get("tb")
	basehdd, err := strconv.ParseFloat(basehddStr, 64)
	if err != nil {
		basehdd = 0
	}

	baseddStr := r.Form.Get("tr")
	basedd, err := strconv.ParseFloat(baseddStr, 64)
	if err != nil {
		basedd = 0
	}

	avg, err := strconv.Atoi(r.Form.Get("avg"))
	if err != nil {
		avg = 0
	}

	WeekStart := common.StrDayToWeekday(r.Form.Get("week_start"))

	key := r.Form.Get("key")
	start := r.Form.Get("start")
	end := utils.WordToTime(r.Form.Get("end"))
	output := vars[Method]
	breakdown := r.Form.Get("breakdown")
	dayCalc := vars[DayCalc]
	format := r.Form.Get("format")
	metric := parseMetric(r.Form.Get("metric"))

	stsSrc := r.Form.Get("station")
	sts := common.ParseStations(stsSrc)

	prms := make([]Params, len(sts))
	for i, v := range sts {
		prms[i] = Params{
			Key:       key,
			Station:   v,
			Start:     start,
			End:       utils.WordToTime(end),
			Tb:        basehdd,
			Tr:        basedd,
			Output:    output,
			Breakdown: breakdown,
			DayCalc:   dayCalc,
			Avg:       avg,
			WeekStart: WeekStart,
			Format:    format,
			Metric:    metric,
		}
	}

	return prms
}

func writeCSV(data *DDResponse, w http.ResponseWriter) error {
	filename := utils.GetCSVName(data.Params.Output, data.Params.Station, data.Params.Tb, data.Params.Tr)
	d := getCSVData(data)

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)
	wr := csv.NewWriter(w)
	wr.Comma = ';'
	err := wr.WriteAll(d)
	wr.Flush()
	if err != nil {
		http.Error(w, "Error sending csv: "+err.Error(), http.StatusInternalServerError)
	}
	return err
}

func writeJSON(data *DDResponse, w http.ResponseWriter) error {
	filename := utils.GetJSONName(data.Params.Output, data.Params.Station, data.Params.Tb, data.Params.Tr)
	d := getJSONData(data)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)
	json.NewEncoder(w).Encode(d)
	return nil
}

func writeZIP(data []*DDResponse, format string, w http.ResponseWriter) error {
	if len(data) == 0 {
		return errors.New("no data found")
	}

	filename := utils.GetZIPName(data[0].Params.Output)

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	zipW := zip.NewWriter(w)
	defer zipW.Close()
	var name string
	var err error
	for _, v := range data {
		if v == nil {
			continue
		}
		//Write JSON
		if format == common.FormatJSON {
			name = utils.GetJSONName(v.Params.Output, v.Params.Station, v.Params.Tb, v.Params.Tr)
			d := getJSONData(v)

			f, err := zipW.Create(name)
			if err != nil {
				return err
			}
			b, err := json.Marshal(d)

			_, err = f.Write(b)
			if err != nil {
				return err
			}

			continue
		}

		//Write CSV
		name = utils.GetCSVName(v.Params.Output, v.Params.Station, v.Params.Tb, v.Params.Tr)
		d := getCSVData(v)

		f, err := zipW.Create(name)
		if err != nil {
			return err
		}
		var buffer bytes.Buffer
		writer := csv.NewWriter(&buffer)
		writer.Comma = ';'
		err = writer.WriteAll(d)
		if err != nil {
			return err
		}
		writer.Flush()

		_, err = f.Write(buffer.Bytes())
		if err != nil {
			return err
		}

	}

	return err
}

func getJSONData(data *DDResponse) *utils.JSONData {
	if len(data.Average) > 0 {
		return utils.GenerateAvgJSON(data.Temps, data.Average, data.Params, data.Autocomplete)
	} else {
		return utils.GenerateJSON(data.Temps, data.Params, data.Autocomplete)
	}
}

func getCSVData(data *DDResponse) [][]string {
	if len(data.Average) > 0 {
		return utils.GenerateAvgCSV(data.Temps, data.Average, data.Params, data.Autocomplete)
	} else {
		return utils.GenerateCSV(data.Temps, data.Params, data.Autocomplete)
	}
}

func parseMetric(metric string) bool {
	if metric == "" || metric == "true"{
		return true
	}
	return false
}
