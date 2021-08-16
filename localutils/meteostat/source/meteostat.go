package source

import (
	"github.com/flasherup/gradtage.de/localutils/meteostat/parser"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"net/http"
	"time"
)

type Meteostat struct {
	key string
	logger log.Logger
}

func NewMeteostat(key string, logger log.Logger) *Meteostat {
	return &Meteostat{
		key: 	key,
		logger:	logger,
	}
}

func (mst Meteostat) FetchTemperature(ch chan *parser.ParsedData,  ids []string) {
	for _,v := range ids {
		go mst.fetchStation(v, ch)
	}
}


func (mst Meteostat)fetchStation(id string, ch chan *parser.ParsedData) {
	url := "https://api.meteostat.net/v1/history/hourly?station=" + id + "&start=2017-01-01&end=2017-01-02&key=" + mst.key
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		level.Error(mst.logger).Log("msg", "request error", "err", err)
		ch <- &parser.ParsedData{ Success:false, Error:err }
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		level.Error(mst.logger).Log("msg", "request error", "err", err)
		ch <- &parser.ParsedData{ Success:false, Error:err }
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		level.Error(mst.logger).Log("msg", "response read error", "err", err)
		ch <- &parser.ParsedData{ Success:false, Error:err }
		return
	}

	temp, err := parser.ParseMeteostat(&contents)
	if err != nil {
		level.Error(mst.logger).Log("msg", "response data parse error", "err", err, "response", string(contents))
		ch <- &parser.ParsedData{ Success:false, Error:err }
		return
	}

	if temp == nil {
		level.Error(mst.logger).Log("msg", "response error", "err", "station:" + id + " not found", "response", string(contents))
		ch <- &parser.ParsedData{ Success:false, Error:err }
		return
	}

	res := parser.ParsedData{
		Success:true,
		StationID:id,
		Temps: *temp,
	}

	ch <- &res
}