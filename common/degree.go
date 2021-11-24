package common

import (
	"time"
)

type Temperature struct {
	Date 		string `json:"date"`
	Temp		float64 `json:"temp"`
}

type TempGroup struct {
	Temps []Temperature
	Date time.Time
}

func CalculateCDDegree(temps []Temperature, baseCDD float64, outputPeriod string, dayCalc string) (res *[]Temperature) {
	cb := func(temp float64) float64 {
		return  calculateCDD(baseCDD, temp)
	}
	return calculateDegree(temps, outputPeriod, dayCalc, cb)
}

func CalculateDDegree(temps []Temperature, baseHDD, baseDD float64, outputPeriod string, dayCalc string) (res *[]Temperature) {
	cb := func(temp float64) float64 {
		return  calculateDD(baseHDD, baseDD, temp)
	}
	return calculateDegree(temps, outputPeriod, dayCalc, cb)
}

func CalculateHDDDegree(temps []Temperature, baseHDD float64, outputPeriod string, dayCalc string) (res *[]Temperature) {
	cb := func(temp float64) float64 {
		return  calculateHDD(baseHDD, temp)
	}
	return calculateDegree(temps, outputPeriod, dayCalc, cb)
}

func calculateDegree(temps []Temperature, outputPeriod string, dayCalc string, calcFunc func(float64) float64) *[]Temperature {
	daily := groupByPeriod(&temps, BreakdownDaily)
	dailyTemps := make([]Temperature, len(*daily))
	for i,v := range *daily {
		temp := calculateDayDegree(&v.Temps, dayCalc, calcFunc)
		dStr := getDateString(v.Date, BreakdownDaily)
		dailyTemps[i] = Temperature{
			dStr,
			temp,
		}
	}

	if outputPeriod != BreakdownDaily {
		return sumPeriod(&dailyTemps, outputPeriod, TimeLayoutDay)
	}

	return &dailyTemps;
}

func sumPeriod(temps *[]Temperature, outputPeriod string, tLayout string) *[]Temperature {
	res := make([]Temperature, 0)
	var lastDate time.Time
	sum := 0.0
	latestIndex := len(*temps)-1
	for i,temp := range *temps {
		currentDate, err := time.Parse(tLayout, temp.Date)
		if err != nil {
			continue
		}

		if !isTheSamePeriod(lastDate, currentDate, outputPeriod) || i == latestIndex {
			if !lastDate.IsZero() {
				dStr := getDateString(lastDate, outputPeriod)
				sum = ToFixedFloat64(sum, 2)
				res = append(res, Temperature{Date: dStr, Temp: sum})
			}
			sum = temp.Temp
			lastDate = currentDate
		} else {
			sum += temp.Temp
		}
	}

	return &res
}


func groupByPeriod(temps *[]Temperature, outputPeriod string) *[]TempGroup {
	res := make([]TempGroup, 0)
	var lastDate time.Time
	var period = make([]Temperature, 0)
	latestIndex := len(*temps)-1
	for i,temp := range *temps {
		currentDate, err := time.Parse(TimeLayout, temp.Date)
		if err != nil {
			continue
		}
		if isTheSamePeriod(lastDate, currentDate, outputPeriod) && i != latestIndex {
			period = append(period, temp)
		} else {
			if len(period) > 0 {
				res = append(res, TempGroup{period, lastDate})
			}
			period = []Temperature{temp}
			lastDate = currentDate
		}
	}
	return &res
}

func getDateString(date time.Time, breakdown string) string {
	if breakdown == BreakdownDaily {
		return date.Format(TimeLayoutDay)
	}

	if breakdown == BreakdownMonthly {
		return date.Format(TimeLayoutMonth)
	}

	if breakdown == BreakdownYearly {
		return date.Format(TimeLayoutYear)
	}

	return date.Format(TimeLayout)
}

func isTheSamePeriod(last, current time.Time, breakdown string) bool {
	if last.IsZero() {
		return false
	}
	return getPeriodDateMarker(last, breakdown) == getPeriodDateMarker(current, breakdown)
}

func getPeriodDateMarker(date time.Time, breakdown string) int {
	if breakdown == BreakdownDaily {
		return date.YearDay()
	}

	if breakdown == BreakdownMonthly {
		return int(date.Month())
	}

	if breakdown == BreakdownYearly {
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
func calculateDayDegree(data *[]Temperature, dayCalcType string, calcFunc func(float64) float64) float64 {
	res := 0.0
	if dayCalcType == DayCalcInt {
		daily := make([]float64, len(*data))
		for i,v := range *data {
			daily[i] = calcFunc(v.Temp)
			res = GetAverageFloat64(daily)
		}
	} else if dayCalcType == DayCalcMean {
		daily := make([]float64, len(*data))
		for i,v := range *data {
			daily[i] = v.Temp
			a := GetAverageFloat64(daily)
			res = calcFunc(a)
		}
	} else if dayCalcType == DayCalcMima {
		daily := make([]float64, len(*data))
		for i,v := range *data {
			daily[i] = v.Temp
			min,max := getMinMaxFloat64(daily)
			a := GetAverageFloat64([]float64{min,max})
			res = calcFunc(a)
		}
	}
	return ToFixedFloat64(res, 2)
}