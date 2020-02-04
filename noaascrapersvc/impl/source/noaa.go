package source

import (
	"github.com/flasherup/gradtage.de/noaascrapersvc/impl/parser"
	"github.com/go-kit/kit/log"
)

type SourceNOAA struct {
	url string
	logger log.Logger
}


func NewNOAA(url string, logger log.Logger) *SourceNOAA {
	return &SourceNOAA{
		url: 	url,
		logger:	logger,
	}
}

func (sn SourceNOAA) FetchTemperature(ch chan *parser.ParsedData, ids map[string]string) {
	for k,v := range ids {
		go sn.fetchStation(k, v, ch)
	}
}

func (sn SourceNOAA)fetchStation(id string, srcId string, ch chan *parser.ParsedData) {


	temps, err := parser.ParseNOAA(&[]byte{})
	if err != nil {
		ch <- &parser.ParsedData{ Success:false, Error:err }
		return
	}

	res := parser.ParsedData{
		Success:true,
		StationID:id,
		Temps:*temps,
	}

	ch <- &res
}