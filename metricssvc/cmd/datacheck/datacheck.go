package main

import (
	"fmt"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/cmd/datacheck/utils"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/config"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/impl/database"
	"github.com/flasherup/gradtage.de/weathrbitupdatesvc/impl/parser"
)

func main() {
	//Config
	conf, err := config.LoadConfig("src/config.yml")
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

	err = db.CreateTable("test")
	if err != nil {
		fmt.Println("Can't create db error", err)
		return
	}

	checkResponse(utils.StationRespose_not_clear, db)
}

func checkResponse(response string, db *database.Postgres) {
	b := []byte(response)
	wbd, err := parser.ParseWeatherBit(&b)
	if err != nil {
		fmt.Println("Parse data error", err)
		return
	}

	err = db.PushData("test",wbd)
	if err != nil{
		fmt.Println("Data push error", err)
	}
}
