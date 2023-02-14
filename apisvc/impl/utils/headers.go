package utils

import (
	"fmt"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
	"strings"
	"time"
)

func getIndicator(output string) string {
	if output == common.HDDType {
		return "Heating Degree Days"
	}

	if output == common.CDDType {
		return "Cooling Degree Days"
	}

	if output == common.DDType {
		return "Room Heating Degree Days"
	}

	return ""
}

func getMethod(dayCalc string) string {
	if dayCalc == common.DayCalcInt {
		return "Integration Method"
	}

	if dayCalc == common.DayCalcMean {
		return "Daily mean temperature"
	}

	if dayCalc == common.DayCalcMima {
		return "Daily min./max. average temperature"
	}

	return ""
}

func getTB(tb float64, metric bool) string {
	if metric {
		return fmt.Sprintf("%gC",tb)
	}

	return fmt.Sprintf("%gF",tb)
}

func getTR(tr float64, output string, metric bool) string{
	if output != common.DDType {
		return "---"
	}

	if metric {
		return fmt.Sprintf("%gC",tr)
	}

	return fmt.Sprintf("%gF",tr)
}

func getUnits(metric bool) string{
	if metric {
		return "celsius"
	}

	return "fahrenheit"
}

func getStation(autocomplete autocompletesvc.Autocomplete) string {
	return fmt.Sprintf("%s, %s, %s", autocomplete.ID, autocomplete.CityNameNative, autocomplete.CountryNameNative)
}

func getDescription(params daydegreesvc.Params) string {
	method := getMethodDescription(params.DayCalc)
	unit := "F"
	if params.Metric {
		unit = "C"
	}
	if params.Output == common.HDDType {
		return fmt.Sprintf(
			"Heating Degree Days with a base temperature of %g%s based on %s",
			params.Tb,
			unit,
			method,
		)
	}

	if params.Output == common.CDDType {
		return fmt.Sprintf(
			"Cooling Degree Days with a base temperature of %g%s based on %s",
			params.Tb,
			unit,
			method,
		)
	}

	if params.Output == common.DDType {
		return fmt.Sprintf(
			"Room Heating Degree Days (Gradtagzahl) with a base temperature of %g%s and a room temperature of %g%s based on %s",
			params.Tb,
			unit,
			params.Tr,
			unit,
			method,
		)
	}

	return ""
}

func getMethodDescription(method string) string {
	if method == common.DayCalcInt {
		return "integration method"
	}

	if method == common.DayCalcMean {
		return "daily mean temperature"
	}

	if method == common.DayCalcMima {
		return "daily min./max. average temperature"
	}

	return ""
}

func getFormattedValue(percentageValue float64) string{
	value := fmt.Sprintf("%.2f", percentageValue)
	return value
}

func getFormattedDate(date string) string{
	return strings.Replace(date, "-", "/", -1)
}

func getAvgIndex(date string, breakdown string, WeekStart time.Weekday) int {
	timeLayout := common.TimeLayoutDay
	if breakdown == common.BreakdownWeekly {
		timeLayout = common.TimeLayoutDay
	} else if breakdown == common.BreakdownWeeklyISO {
		timeLayout = common.TimeLayoutDay
	} else if breakdown == common.BreakdownMonthly {
		timeLayout = common.TimeLayoutMonth
	} else if breakdown == common.BreakdownYearly {
		timeLayout = common.TimeLayoutYear
	}

	d, err := time.Parse(timeLayout, date)
	if err == nil {
		if breakdown == common.BreakdownDaily {
			return common.LeapYearDay(d)
		} else if breakdown == common.BreakdownWeekly {
			return common.Week(d, WeekStart)
		}  else if breakdown == common.BreakdownWeeklyISO {
			return common.WeekISO(d)
		} else if breakdown == common.BreakdownMonthly {
			return int(d.Month())
		} else if breakdown == common.BreakdownYearly {
			return 1
		}
	} else {
		fmt.Println("Time parse error", err)
	}
	return 1
}
