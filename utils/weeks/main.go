package main

import (
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"log"
	"time"
)

func main() {
	//dayCycle()
	years()
}

func years() {
	startYear := 2010
	endYear := time.Now().Year()

	for year := startYear; year < endYear; year++ {
		start := getDateByYearMonthDay(year, 1, 1)
		week := common.Week(start, time.Monday)
		for week != 1 {
			start = start.Add(time.Hour * 24)
			week = common.Week(start, time.Monday)
		}

		end := getDateByYearMonthDay(year, 12, 31)
		week = common.Week(end, time.Monday)
		for week < 52 {
			end = end.Add(-time.Hour * 24)
			week = common.Week(end, time.Monday)
		}

		weeks := common.Week(end, time.Monday)

		fmt.Printf("Year: %d, weeks start: %d, weeks end: %d number of weeks: %d \n", year, start.Day(),  end.Day(), weeks)
	}
}

func dayCycle() {
	start, err := time.Parse(common.TimeLayout, "2010-01-01T00:00:00Z")
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()

	latestWeek := common.Week(start, time.Monday)
	for now.Sub(start) > 0 {
		start = start.Add(time.Hour * 24)
		week := common.Week(start, time.Monday)
		if week != latestWeek {
			if week == 1 {
				fmt.Printf("Year: %d, latestWeek: %d, Month: %d, Day: %d \n", start.Year(), latestWeek, start.Month(), start.Day())
			}
			latestWeek = week
		}
	}
}

func getDateByYearMonthDay(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 1,1,1,1,time.Now().Location())
}
