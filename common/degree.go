package common

import (
	"time"
)

type Temperature struct {
	Date 		string `json:"date"`
	Temp		float64 `json:"temp"`
}

func CalculateCDDegree(temps []Temperature, baseCDD float64, outputPeriod int) (res *[]Temperature) {
	cb := func(temp float64) float64 {
		return  calculateCDD(baseCDD, temp)
	}
	return calculateDegree(temps, outputPeriod, cb)
}

func CalculateDDegree(temps []Temperature, baseHDD, baseDD float64, outputPeriod int) (res *[]Temperature) {
	cb := func(temp float64) float64 {
		return  calculateDD(baseHDD, baseDD, temp)
	}
	return calculateDegree(temps, outputPeriod, cb)
}

func CalculateHDDDegree(temps []Temperature, baseHDD float64, outputPeriod int) (res *[]Temperature) {
	cb := func(temp float64) float64 {
		return  calculateHDD(baseHDD, temp)
	}
	return calculateDegree(temps, outputPeriod, cb)
}

func calculateDegree(temps []Temperature, outputPeriod int, calcFunc func(float64) float64) *[]Temperature {
	res := make([]Temperature, 0)
	var lastDate time.Time
	var p []float64
	for i,temp := range temps {
		currentDate, err := time.Parse(TimeLayout, temp.Date)
		if err != nil {
			continue
		}
		if isTheSemePeriod(lastDate, currentDate, outputPeriod) && i != len(temps)-1 {
			dDegree := calcFunc(temp.Temp)
			p = append(p, dDegree)
		} else {
			if len(p) > 0 {
				avg := getAverageFloat64(p)
				avg = ToFixedFloat64(avg, 2)
				dStr := getDateString(lastDate, outputPeriod)
				res = append(res, Temperature{Date: dStr, Temp: avg})
			}
			lastDate = currentDate
		}
	}
	return &res
}

func getDateString(date time.Time, period int) string {
	if period == PeriodDay {
		return date.Format(TimeLayoutDay)
	}

	if period == PeriodMonth {
		return date.Format(TimeLayoutMonth)
	}

	if period == PeriodYear {
		return date.Format(TimeLayoutYear)
	}

	return date.Format(TimeLayout)
}

func isTheSemePeriod(last, current time.Time, period int) bool {
	if last.IsZero() {
		return false
	}
	return getPeriodDateMarker(last, period) == getPeriodDateMarker(current, period)
}

func getPeriodDateMarker(date time.Time, period int) int {
	if period == PeriodDay {
		return date.YearDay()
	}

	if period == PeriodMonth {
		return int(date.Month())
	}

	if period == PeriodYear {
		return date.Year()
	}

	return -1
}

func calculateHDD(baseHDD float64, value float64) float64 {
	if value >= baseHDD {
		return 0
	}
	return baseHDD - value
}

func calculateDD(baseHDD float64, baseDD float64, value float64) float64 {
	if value >= baseHDD || value >= baseDD{
		return 0
	}

	return baseDD - value
}


func calculateCDD(baseCDD float64, value float64) float64 {
	if value < baseCDD {
		return 0
	}
	return value-baseCDD
}