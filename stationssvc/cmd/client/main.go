package main

import (
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/stationssvc"
	"github.com/flasherup/gradtage.de/stationssvc/impl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "stationssvcc",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	client := impl.NewStationsSCVClient("localhost:8102",logger)

	level.Info(logger).Log("msg", "client started")
	defer level.Info(logger).Log("msg", "client ended")

	//client.AddStations(stations())
	//Just for test
	sts, err := client.GetStationsBySrcType([]string{ common.SrcTypeCheckWX })
	if err != nil {
		level.Error(logger).Log("msg", "GetStationsBySrcType error", "err", err)

	} else {
		for k,v := range sts.Sts {
			level.Info(logger).Log("msg", "sts", "ID", k,  " Name", v.Name,  "Timezone", v.Timezone)
		}
	}



}

func stations() []stationssvc.Station {
	return []stationssvc.Station {
		stationssvc.Station{ID:"KBOS", Timezone:"EST", Name:"Boston"},
		stationssvc.Station{ID:"EKCH", Timezone:"CET", Name:"Kopenhagen"},
		stationssvc.Station{ID:"EGLL", Timezone:"GMT", Name:"London"},
		stationssvc.Station{ID:"CYHU", Timezone:"EST", Name:"Montreal"},
		stationssvc.Station{ID:"KNYC", Timezone:"EST", Name:"New York"},
		stationssvc.Station{ID:"LFPG", Timezone:"CET", Name:"Paris"},
		stationssvc.Station{ID:"ESSB", Timezone:"CET", Name:"Stockholm"},
		stationssvc.Station{ID:"CXTO", Timezone:"EST", Name:"Toronto"},
		stationssvc.Station{ID:"KDCA", Timezone:"EST", Name:"Washington"},
		stationssvc.Station{ID:"EDDT", Timezone:"CET", Name:"Berlin (TXL)"},
		stationssvc.Station{ID:"EDDB", Timezone:"CET", Name:"Berlin (SXF)"},
		stationssvc.Station{ID:"EDDH", Timezone:"CET", Name:"Hamburg"},
		stationssvc.Station{ID:"EDDM", Timezone:"CET", Name:"München"},
		stationssvc.Station{ID:"EDDK", Timezone:"CET", Name:"Köln"},
		stationssvc.Station{ID:"EDDF", Timezone:"CET", Name:"Frankfurt am Main"},
		stationssvc.Station{ID:"EDDS", Timezone:"CET", Name:"Stuttgart"},
		stationssvc.Station{ID:"EDDL", Timezone:"CET", Name:"Düsseldorf"},
		stationssvc.Station{ID:"EDDP", Timezone:"CET", Name:"Leipzig"},
		stationssvc.Station{ID:"UUEE", Timezone:"MSK", Name:"Moskau"},
		stationssvc.Station{ID:"LTBA", Timezone:"EET", Name:"Instanbul"},
		stationssvc.Station{ID:"ULLI", Timezone:"MSK", Name:"St. Petersburg"},
		stationssvc.Station{ID:"LEMD", Timezone:"CET", Name:"Madrid"},
		stationssvc.Station{ID:"UKKK", Timezone:"CET", Name:"Kiev"},
		stationssvc.Station{ID:"LIRA", Timezone:"CET", Name:"Rom"},
		stationssvc.Station{ID:"LRBS", Timezone:"EET", Name:"Bucharest"},
		stationssvc.Station{ID:"UMMS", Timezone:"MSK", Name:"Minsk"},
		stationssvc.Station{ID:"LOWW", Timezone:"CET", Name:"Wien"},
		stationssvc.Station{ID:"EPWA", Timezone:"CET", Name:"Warschau"},
		stationssvc.Station{ID:"LHBP", Timezone:"CET", Name:"Budapest"},
		stationssvc.Station{ID:"LEBL", Timezone:"CET", Name:"Barcelona"},
		stationssvc.Station{ID:"LIME", Timezone:"CET", Name:"Milan/ Bergamo"},
		stationssvc.Station{ID:"LKPR", Timezone:"CET", Name:"Prag"},
		stationssvc.Station{ID:"EHAM", Timezone:"CET", Name:"Amsterdam"},
		stationssvc.Station{ID:"EBBR", Timezone:"CET", Name:"Brüssel"},
		stationssvc.Station{ID:"LFML", Timezone:"CET", Name:"Marseille"},
		stationssvc.Station{ID:"LPMT", Timezone:"WET", Name:"Lissabon"},
		stationssvc.Station{ID:"KCQT", Timezone:"PST", Name:"Los Angeles"},
		stationssvc.Station{ID:"KMDW", Timezone:"CST", Name:"Chicago"},
		stationssvc.Station{ID:"KPHL", Timezone:"EST", Name:"Philadelphia PA"},
		stationssvc.Station{ID:"KSAN", Timezone:"PST", Name:"San Diego CA"},
		stationssvc.Station{ID:"CYVR", Timezone:"PST", Name:"Vancouver"},
		stationssvc.Station{ID:"CYYC", Timezone:"MST", Name:"Calgary"},
		stationssvc.Station{ID:"KATT", Timezone:"CST", Name:"Austin TX"},
	}
}
