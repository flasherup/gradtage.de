package source

import (
	"github.com/flasherup/gradtage.de/hourlysvc"
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

func (cwx CheckWX) FetchTemperature(ch chan *parser.ParsedData,  ids []string) {
	for _,v := range ids {
		cwx.fetchStation(v, ch)
	}
}


func (cwx CheckWX)fetchStation(id string, ch chan *parser.ParsedData) {
	url := "https://api.checkwx.com/metar/" + id + "/decoded"
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		level.Error(cwx.logger).Log("msg", "request error", "err", err)
		ch <- &parser.ParsedData{ Success:false, StationID:id, Error:err }
		return
	}
	req.Header.Add("X-API-Key", cwx.key)
	resp, err := client.Do(req)
	if err != nil {
		level.Error(cwx.logger).Log("msg", "request error", "err", err)
		ch <- &parser.ParsedData{ Success:false, StationID:id, Error:err }
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		level.Error(cwx.logger).Log("msg", "response read error", "err", err)
		ch <- &parser.ParsedData{ Success:false, StationID:id, Error:err }
		return
	}

	temp, err := parser.ParseCheckWX(&contents)
	if err != nil {
		level.Error(cwx.logger).Log("msg", "response data parse error", "err", err, "response", string(contents))
		ch <- &parser.ParsedData{ Success:false, StationID:id, Error:err }
		return
	}

	if temp == nil {
		level.Error(cwx.logger).Log("msg", "response error", "err", "station:" + id + " not found", "response", string(contents))
		ch <- &parser.ParsedData{ Success:false, StationID:id, Error:err }
		return
	}

	res := parser.ParsedData{
		Success:true,
		StationID:id,
		Temps:[]hourlysvc.Temperature{ *temp },
	}

	ch <- &res
}