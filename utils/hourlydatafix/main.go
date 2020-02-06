package main

import (
	"flag"
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/flasherup/gradtage.de/utils/hourlydatafix/config"
	"github.com/flasherup/gradtage.de/utils/hourlydatafix/database"
	"time"
)

func main() {
	configFile := flag.String("config.file", "config.yml", "Config file name.")
	flag.Parse()

	conf, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println("msg", "config loading error", "exit", err.Error())
		return
	}

	//Database
	db, err := database.NewPostgres(conf.Database)
	if err != nil {
		fmt.Println("msg", "database error", "exit", err.Error())
		return
	}

	tables, err := db.GetListOfTables()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _,v := range tables {
		tmep, err := db.GetAll(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Parse temp for station", v)
		fTemp := fixTemps(tmep)
		//fmt.Println(fTemp)

		err = db.RemoveTable(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = db.CreateTable(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = db.PushPeriod(v, fTemp)
		if err != nil {
			fmt.Println(err)
			continue
		}
		time.Sleep(time.Second * 5)
	}
}

const checkWXTimeTemplate = "2006-01-02T15:04:05.000Z"

func fixTemps(src []hourlysvc.Temperature) []hourlysvc.Temperature {
	res := make([]hourlysvc.Temperature, 0, len(src))
	saved := map[string]int{}
	var d time.Time
	var err error
	for _,v := range src {
		d, err = time.Parse(checkWXTimeTemplate, v.Date)
		if err != nil {
			d, err = time.Parse(common.TimeLayout, v.Date)
			if err != nil {
				continue
			}
		}
		dp := time.Date(d.Year(), d.Month(), d.Day(), d.Hour(), 0, 0, 0, d.Location())

		t := v.Temperature
		temp := hourlysvc.Temperature {
			Date: 			dp.Format(common.TimeLayout),
			Temperature: 	t,
		}

		if index, ok := saved[temp.Date]; ok {
			fmt.Println("Copy has found", temp.Date)
			res[index] = temp
		} else {
			res = append(res, temp)
			saved[temp.Date] = len(res)-1
		}
	}

	return res
}