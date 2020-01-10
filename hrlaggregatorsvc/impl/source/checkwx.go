package source

import (
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/impl/parser"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"net/http"
	"time"
)

type CheckWX struct {
	key string
	logger log.Logger
}

func NewCheckWX(key string, logger log.Logger) *CheckWX {
	return &CheckWX{
		key: 	key,
		logger:	logger,
	}
}

func (cwx CheckWX) FetchTemperature(ch chan *parser.StationData, ids []string) {
	for _,v := range ids {
		go cwx.fetchStation(v, ch)
	}
}


func (cwx CheckWX)fetchStation(id string, ch chan *parser.StationData) {
	url := "https://api.checkwx.com/metar/" + id + "/decoded"
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		level.Error(cwx.logger).Log("msg", "request error", "err", err)
		ch <- nil
		return
	}
	req.Header.Add("X-API-Key", cwx.key)
	resp, err := client.Do(req)
	if err != nil {
		level.Error(cwx.logger).Log("msg", "request error", "err", err)
		ch <- nil
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		level.Error(cwx.logger).Log("msg", "response read error", "err", err)
		ch <- nil
		return
	}

	temp, err := parser.ParseCheckWX(&contents)
	if err != nil {
		level.Error(cwx.logger).Log("msg", "response data parse error", "err", err, "response", string(contents))
		ch <- nil
		return
	}

	if len(*temp) < 1 {
		level.Error(cwx.logger).Log("msg", "response error", "err", "station:" + id + " not found", "response", string(contents))
		ch <- nil
		return
	}

	for _,v := range *temp{
		ch <- &v
	}
}