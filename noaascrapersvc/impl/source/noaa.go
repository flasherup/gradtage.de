package source

import (
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/noaascrapersvc/impl/parser"
	"github.com/go-kit/kit/log"
	"github.com/gocolly/colly"
	"sync"
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
	wg := sync.WaitGroup{}
	c := colly.NewCollector()

	isDataScraped := false
	// Find and visit all links
	c.OnHTML("table", func(e *colly.HTMLElement) {
		if e.Index == 3 {
			if isDataScraped {
				return
			}
			isDataScraped = true

			temps, err := parser.ParseNOAATable(e)
			if err != nil {
				ch <- &parser.ParsedData{ Success:false, Error:err }
			} else {
				res := parser.ParsedData{
					Success:true,
					StationID:id,
					Temps:*temps,
				}
				ch <- &res
			}
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		ch <- &parser.ParsedData{ Success:false, Error:err }
		wg.Done()
	})

	c.OnScraped(func(r *colly.Response) {
		if !isDataScraped {
			isDataScraped = true
			ch <- &parser.ParsedData{ Success:false, Error:errors.New("data was not found") }
		}
		wg.Done()
	})

	wg.Add(1)
	url := fmt.Sprintf("%s%s.html", sn.url, srcId)
	c.Visit(url)
	wg.Wait()
}