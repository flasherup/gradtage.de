package parser

import (
	"errors"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"github.com/gocolly/colly"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	dateColumn = 0
	timeColumn = 1
	tempColumn = 6
)

type WeatherData struct {
	Date string
	Time string
	Temp string
}

func ParseNOAATable(table *colly.HTMLElement) (*[]hourlysvc.Temperature, error) {
	temps := make([]hourlysvc.Temperature, 0)
	wd := WeatherData{}
	//Table Rows
	table.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
		//Row columns
		tr.ForEach("td", func(i int, td *colly.HTMLElement) {
			switch i {
			case dateColumn:
				wd.Date = td.Text
			case timeColumn:
				wd.Time = td.Text
			case tempColumn:
				wd.Temp = td.Text
			}
		})

		temp, err := WeatherDataToTemperature(wd)
		if err == nil {
			temps = append([]hourlysvc.Temperature{*temp}, temps...)
		}
	})

	return &temps,nil
}

func WeatherDataToTemperature(wd WeatherData) (*hourlysvc.Temperature, error) {
	day, err := strconv.Atoi(wd.Date)
	if err != nil {
		return nil, err
	}

	h,m,err := parseTime(wd.Time)
	if err != nil {
		return nil,err
	}

	temp, err := strconv.ParseFloat(wd.Temp, 10)
	if err != nil {
		return nil,err
	}

	temp = fahrenheitToCelsius(temp)

	sDate := time.Now()
	date := time.Date(sDate.Year(), sDate.Month(), day, h, m, 0, 0, sDate.Location())

	time := date.Format(common.TimeLayout)
	return &hourlysvc.Temperature{ time, temp }, nil
}

func parseTime(src string) (h,m int, err error) {
	if len(src) == 0 {
		return 0,0, errors.New("time string is empty")
	}

	sp := strings.Split(src, ":")
	if len(sp) < 2 {
		return 0,0, errors.New("time string is wrong: " + src)
	}

	h, err = strconv.Atoi(sp[0])
	if err != nil {
		return 0,0, err
	}

	m, err = strconv.Atoi(sp[1])
	if err != nil {
		return 0,0, err
	}

	return h,m,err
}

func fahrenheitToCelsius(f float64) float64 {
	c := (f - 32.0) * 5.0 / 9.0
	c =	math.Floor(c * 10) / 10
	return c
}