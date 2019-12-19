package logwriter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flasherup/gradtage.de/monitor"
	"io/ioutil"
	"net/http"
	"time"
)

type MonitorWriter struct {
	client *http.Client
	url string
}

func NewMonitorWriter(url string) *MonitorWriter {
	return &MonitorWriter{
		client: &http.Client{ Timeout: time.Second * 10 },
		url:url,
	}
}

func (mw *MonitorWriter)Write(p []byte) (n int, err error) {
	fmt.Println("log", p)
	lr := monitor.LogRequest{Log:string(p)}

	bt := new(bytes.Buffer)
	json.NewEncoder(bt).Encode(lr)

	req, err := http.NewRequest("POST", mw.url + "/log/", bt)
	if err != nil {
		fmt.Println("log request create error:", err)
		return
	}

	//req.Header.Add("X-API-Key", `d8fd92d0101dd44410f9984550`)
	resp, err := mw.client.Do(req)
	if err != nil {
		fmt.Println("log error:", err)
		return
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("log response error:",err)
		return
	}
	n = len(p)
	return
}