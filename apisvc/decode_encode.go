package apisvc

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
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
	return req, nil
}

func encodeGetHDDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(GetHDDResponse)
	w.Header().Set("Content-Type", "text/csv")
	wr := csv.NewWriter(w)
	err := wr.WriteAll(resp.Data)
	wr.Flush()
	if err != nil {
		http.Error(w, "Error sending csv: "+err.Error(), http.StatusInternalServerError)
	}
	return err
}

func encodeGetHDDCSVResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(GetHDDCSVResponse)
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
		Key :     r.Form.Get("key"),
		Station : r.Form.Get("station"),
		Start :   r.Form.Get("start"),
		End :     r.Form.Get("end"),
		TB:       basehdd,
		TR:       basedd,
		Output :  vars[Method],
	}

	req  := GetHDDCSVRequest{ prm }
	return req, nil
}

func decodeGetSourceDataRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	r.ParseForm()

	prm := ParamsSourceData{
		Key :		r.Form.Get("key"),
		Station : 	r.Form.Get("station"),
		Start : 	r.Form.Get("start"),
		End : 		r.Form.Get("end"),
		Type : 		r.Form.Get("type"),
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
	fmt.Println("decodeUserRequest")
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

	fmt.Println(params)
	return UserRequest{params}, nil
}

func encodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("encodeUserResponse")
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

func decodePlanRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	r.ParseForm()
	p := make(map[string]string)
	for k,v := range r.Form {
		p[k] = v[0]
	}

	params := ParamsPlan{
		Key: 		r.Form.Get("key"),
		Action :	vars[UserAction],
		Params: 	p,
	}
	return PlanRequest{params}, nil
}

func encodePlanResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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

func decodeStripeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := StripeRequest{}

	body, _ := ioutil.ReadAll(r.Body)
	if e := json.Unmarshal(body, &req.Event); e != nil {
		return nil, e
	}
	return req, nil
}

func encodeStripeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(StripeResponse)
	bt := new(bytes.Buffer)
	err := json.NewEncoder(bt).Encode(resp)

	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write(bt.Bytes())
	return err
}

func decodeCommandRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	r.ParseForm()
	p := make(map[string]string)
	for k,v := range r.Form {
		p[k] = v[0]
	}
	fmt.Println("params", p)
	req := CommandRequest{}
	req.Params = p;
	if name, ok := p["name"]; ok {
		req.Name = name
	}
	return req, nil
}

func encodeCommandResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(CommandResponse)
	bt := new(bytes.Buffer)
	err := json.NewEncoder(bt).Encode(resp.Json)

	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write(bt.Bytes())
	return err
}

