package utils

import (
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
	"strconv"
	"time"
)

func WeeklyAverage(degrees []common.Temperature, params daydegreesvc.Params) ([]daydegreesvc.Degree, error){
	weeks := make(map[string][]float64)

	for _, v := range degrees {
		d, err := common.ParseTimeByBreakdown(v.Date, params.Breakdown)
		if err != nil{
			return []daydegreesvc.Degree{}, err
		}
		key := getDateKey(d, params.Breakdown, params.WeekStart)
		day, exist := weeks[key]
		if !exist {
			day = make([]float64, 0)
		}

		weeks[key] = append(day, v.Temp)
	}

	res := make([]common.Temperature, 0)

	initialDate, _ := time.Parse(common.TimeLayout, common.InitialDate)
	year := initialDate.Year()
	for initialDate.Year() == year {
		key := getDateKey(initialDate, params.Breakdown, params.WeekStart)
		day, exist := weeks[key]
		var temp = 0.0

		d := common.GetDateStringByBreakdown(initialDate, params.Breakdown)
		if exist {
			temp = common.GetAverageFloat64(day)
			temp = common.ToFixedFloat64(temp, 2)
		} else {
			temp = common.EmptyWeather
		}

		res = append(res, common.Temperature{
			Date: d,
			Temp: temp,
		})

		initialDate = addPeriod(initialDate, params.Breakdown)
	}

	return ToDegree(res), nil
}

func CommonAverage(degrees []common.Temperature, params daydegreesvc.Params) ([]daydegreesvc.Degree, error){
	days := make(map[string][]float64)

	for _, v := range degrees {
		d, err := common.ParseTimeByBreakdown(v.Date, params.Breakdown)
		if err != nil{
			return []daydegreesvc.Degree{}, err
		}
		key := getDateKey(d, params.Breakdown, params.WeekStart)
		day, exist := days[key]
		if !exist {
			day = make([]float64, 0)
		}

		days[key] = append(day, v.Temp)
	}

	res := make([]common.Temperature, 0)

	initialDate, _ := time.Parse(common.TimeLayout, common.InitialDate)
	year := initialDate.Year()
	for initialDate.Year() == year {
		key := getDateKey(initialDate, params.Breakdown, params.WeekStart)
		day, exist := days[key]
		var temp = 0.0

		d := common.GetDateStringByBreakdown(initialDate, params.Breakdown)
		if exist {
			temp = common.GetAverageFloat64(day)
			temp = common.ToFixedFloat64(temp, 2)
		} else {
			temp = common.EmptyWeather
		}

		res = append(res, common.Temperature{
			Date: d,
			Temp: temp,
		})

		initialDate = addPeriod(initialDate, params.Breakdown)
	}

	return ToDegree(res), nil
}

func addPeriod(src time.Time, breakdown string) time.Time {
	if breakdown == common.BreakdownWeekly {
		return src.AddDate(0, 0, 7)
	}

	if breakdown == common.BreakdownWeeklyISO {
		return src.AddDate(0, 0, 7)
	}

	if breakdown == common.BreakdownMonthly {
		return src.AddDate(0, 1, 0)
	}

	if breakdown == common.BreakdownYearly {
		return src.AddDate(1, 0, 0)
	}

	return src.AddDate(0, 0, 1)
}

func getDateKey(date time.Time, breakdown string, weekStart time.Weekday) string {
	if breakdown == common.BreakdownWeekly {
		w := common.WeekISO(date)
		return strconv.Itoa(w)
	}

	if breakdown == common.BreakdownWeeklyISO {
		w := common.Week(date, weekStart)
		return strconv.Itoa(w)
	}

	return fmt.Sprintf("%d-%d-%d", date.Month(), date.Day(), date.Hour())
}
