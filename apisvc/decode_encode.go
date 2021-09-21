package apisvc

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"github.com/flasherup/gradtage.de/apisvc/impl/utils"
	"github.com/flasherup/gradtage.de/common"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func decodeGetHDDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := GetHDDRequest{}
	if e := json.NewDecoder(r.Body).Decode(&req.Params); e != nil {
		return nil, e
	}
	vars := mux.Vars(r)
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

func decodeGetHDDCSVRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	r.ParseForm()
	basehddStr := r.Form.Get("tb")
	basehdd, err := strconv.ParseFloat(basehddStr, 64);
	if  err != nil {
		basehdd = 0
	}

	baseddStr := r.Form.Get("tr")
	basedd, err := strconv.ParseFloat(baseddStr, 64);
	if  err != nil {
		basedd = 0
	}

	prm := Params{
		Key :      r.Form.Get("key"),
		Station :  r.Form.Get("station"),
		Start :    r.Form.Get("start"),
		End :      r.Form.Get("end"),
		Tb:        basehdd,
		Tr:        basedd,
		Output :   vars[Method],
		Breakdown: r.Form.Get("breakdown"),
		DayCalc:   vars[DayCalc],
	}

	req  := GetHDDCSVRequest{ prm }
	return req, nil
}

func encodeGetHDDCSVResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(GetHDDCSVResponse)
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=" + resp.FileName)
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
		Key :		r.Form.Get("key"),
		Station : 	r.Form.Get("station"),
		Start : 	r.Form.Get("start"),
		End : 		r.Form.Get("end"),
	}

	req  := GetSourceDataRequest{ prm }
	return req, nil
}

func encodeGetSourceDataResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(GetSourceDataResponse)
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=" + resp.FileName)
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
		Key: 		r.Form.Get("key"),
		Text :		r.Form.Get("text"),
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
		"key":r.Form.Get("key"),
		"plan":r.Form.Get("plan"),
		"name":r.Form.Get("email"),
	}

	if p["key"] == "" {
		return nil, errors.New("key is required")
	}

	params := ParamsUser{
		Key: 		p["key"],
		Action :	vars[UserAction],
		Params: 	p,
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

	req.Event = event;
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
	for k,v := range r.Form {
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

