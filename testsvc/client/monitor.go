package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type MonitorClient struct {
	url string
	client *http.Client
}

func NewMonitorClient(url string) *MonitorClient {
	return &MonitorClient{
		url:url,
		client: &http.Client{ Timeout: time.Second * 10 },
	}
}

func (mc *MonitorClient) Log(s string) (err error) {
	//lr := testsvc.LogRequest{Log: s}

	bt := new(bytes.Buffer)
	//json.NewEncoder(bt).Encode(lr)

	req, err := http.NewRequest("POST", mc.url + "/log/", bt)
	if err != nil {
		fmt.Println("log request create error:", err)
		return
	}

	resp, err := mc.client.Do(req)
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

	return
}
