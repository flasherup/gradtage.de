package main

import (
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/dailysvc"
	daily "github.com/flasherup/gradtage.de/dailysvc/impl"
	stations "github.com/flasherup/gradtage.de/stationssvc/impl"
	"github.com/flasherup/gradtage.de/utils/dwdhistorydata/config"
	"github.com/flasherup/gradtage.de/utils/dwdhistorydata/parser"
	"github.com/flasherup/gradtage.de/utils/dwdhistorydata/source"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gocolly/colly"
	"os"
	"strings"
	"sync"
)

func main() {
	configFile := flag.String("config.file", "dwdhistory.yml", "Config file name.")
	flag.Parse()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "DWDHistoryData",
			"ts", log.DefaultTimestampUTC,
			"caller", log.Caller(3),
		)
	}

	//Config
	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		level.Error(logger).Log("msg", "config loading error", "err", err.Error())
		return
	}

	stations := stations.NewStationsSCVClient(conf.StationsAddr, logger)
	daily := daily.NewDailySCVClient(conf.DailyAddr, logger)

	fileNames, err := getFileNames(conf.UrlDWD, logger)
	if err != nil {
		level.Error(logger).Log("msg", "DWD stations list not get", "err", err.Error())
		return
	}

	sts, err := stations.GetStationsBySrcType([]string{ common.SrcTypeDWD })
	if err != nil {
		level.Error(logger).Log("msg", "Get DWD Stations error", "err", err)
		return
	}

	ids := make(map[string]string)
	for k,v := range sts.Sts {
		fileName,exist := fileNames[v.SourceId]
		if !exist {
			level.Warn(logger).Log("msg", "Station not found on DWD history data", "id", v.SourceId)
		}
		ids[k] = fileName
	}

	dwdData := source.NewDWD(conf.UrlDWD, logger)

	ch := make(chan *parser.ParsedData)
	go dwdData.FetchTemperature(ch, ids)

	for range ids {
		pd := <-ch
		if pd.Error != nil {
			level.Error(logger).Log("msg", "Get DWD data error", "stationID", pd.StationID, "err", pd.Error)
			continue
		}
		fmt.Println(pd.StationID, "Length", len(pd.Temps), "tempFirst", pd.Temps[0], "tempLast", pd.Temps[len(pd.Temps)-1])
		_, err := daily.PushPeriod(pd.StationID, pd.Temps)
		if err != nil {
			level.Error(logger).Log("msg", "PushPeriod Error", "err", err)
		} else {
			//updateAverage(pd.StationID, daily, logger)
		}
	}
}

func getFileNames(url string, logger log.Logger) (map[string]string, error) {
	res := make(map[string]string)
	var err error
	wg := sync.WaitGroup{}
	c := colly.NewCollector()
	c.OnHTML("a", func(e *colly.HTMLElement) {
		if strings.Index(e.Text,".zip") > -1 {
			bs := strings.Split(e.Text, "_")
			if len(bs) > 2 {
				res[bs[2]] = e.Text
			}
		}
	})

	c.OnError(func(_ *colly.Response, e error) {
		err = e
		wg.Done()
	})

	c.OnScraped(func(r *colly.Response) {
		wg.Done()
	})

	wg.Add(1)
	c.Visit(url)
	wg.Wait()
	return res, err
}

func updateAverage(id string, daily dailysvc.Client, logger log.Logger) {
	for i := 0; i< 356; {
		_, err := daily.UpdateAvgForDOY(id, i)
		if err != nil {
			level.Error(logger).Log("msg", "Average Update Error", "err", err)
		}
	}
}
