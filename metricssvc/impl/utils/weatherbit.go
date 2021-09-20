package utils

import (
	"fmt"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/impl/parser"
	"io/ioutil"
	"net/http"
	"time"
)

func MakeWeatherbitRequest(url string) (*parser.WeatherBitData, string, error){
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Sprintf("request create error for url '%s' ", url), err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Sprintf("request error for url '%s' ", url), err
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Sprintf("response read error for url '%s' ", url), err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, fmt.Sprintf("response body close error for url '%s' ", url), err
	}

	data, err := parser.ParseWeatherBit(&contents)
	if err != nil {
		return nil, fmt.Sprintf("weather bit data parse error for url '%s' ", url), err
	}

	return data, fmt.Sprintf("data update success '%s' ", url), nil
}
