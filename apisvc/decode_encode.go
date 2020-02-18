package apisvc

import (
	"context"
	"encoding/csv"
	"encoding/json"
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
	r.ParseForm()
	basehddStr := r.Form.Get("hl")
	basehdd, err := strconv.ParseFloat(basehddStr, 64);
	if  err != nil {
		basehdd = 0
	}

	baseddStr := r.Form.Get("rt")
	basedd, err := strconv.ParseFloat(baseddStr, 64);
	if  err != nil {
		basedd = 0
	}

	prm := Params{
		Key :		r.Form.Get("key"),
		Station : 	r.Form.Get("station"),
		Start : 	r.Form.Get("start"),
		End : 		r.Form.Get("end"),
		HL : 		basehdd,
		RT : 		basedd,
		Output : 	r.Form.Get("output"),
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

func encodeGetSourceDataResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
