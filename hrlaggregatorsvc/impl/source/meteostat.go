package source

import (
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/impl/parser"
	"github.com/go-kit/kit/log"
	"io/ioutil"
	"net/http"
	"time"
)

const meteostatTimeTemplate = "2006-01-02"

type Meteostat struct {
	url string
	key string
	logger log.Logger
}

func NewMeteostat(url string, key string, logger log.Logger) *Meteostat {
	return &Meteostat{
		url:	url,
		key: 	key,
		logger:	logger,
	}
}

func (mst Meteostat) FetchTemperature(ch chan *parser.ParsedData, daysNumber int,  ids map[string]string) {
	start, end := common.GetDatesFromNow(daysNumber, meteostatTimeTemplate)
	for k,v := range ids {
		mst.fetchStation(k, v, ch, start, end)
	}
}


func (mst Meteostat)fetchStation(id string, srcId string, ch chan *parser.ParsedData, start string, end string) {
	url := mst.url + "hourly?station=" + srcId + "&start=" + start + "&end=" + end + "&key=" + mst.key
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//level.Error(mst.logger).Log("msg", "request error", "err", err)
		ch <- &parser.ParsedData{ Success:false, StationID:id, Error:err }
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		//level.Error(mst.logger).Log("msg", "request error", "err", err)
		ch <- &parser.ParsedData{ Success:false, StationID:id, Error:err }
		return
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//level.Error(mst.logger).Log("msg", "response read error", "err", err)
		ch <- &parser.ParsedData{ Success:false, StationID:id, Error:err }
		return
	}

	temp, err := parser.ParseMeteostat(&contents)
	if err != nil {
		//level.Error(mst.logger).Log("msg", "response data parse error", "err", err, "response", string(contents))
		ch <- &parser.ParsedData{ Success:false, StationID:id, Error:err }
		return
	}

	if temp == nil {
		err = fmt.Errorf("station not found, url:%s",url)
		//level.Error(mst.logger).Log("msg", "response error", "err", err, "response", string(contents))
		ch <- &parser.ParsedData{ Success:false, StationID:id, Error:err }
		return
	}

	res := parser.ParsedData{
		Success:true,
		StationID:id,
		Temps: *temp,
	}

	ch <- &res
}