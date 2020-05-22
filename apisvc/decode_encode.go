package apisvc

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
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

	urlParams := map[string][]string{}
	if e := json.NewDecoder(r.Body).Decode(&urlParams); e != nil {
		return nil, e
	}

	k, ok := urlParams["key"]
	if !ok {
		return nil, errors.New("key is required")
	}

	params := ParamsUser{
		Key: 		k[0],
		Action :	vars[UserAction],
		Params: 	urlParams,
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

func decodePlanRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	r.ParseForm()
	p := make(map[string][]string)
	for k,v := range r.Form {
		p[k] = v
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

