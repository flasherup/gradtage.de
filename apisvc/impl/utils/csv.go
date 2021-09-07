package utils

import (
	"fmt"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
	"strings"
)

func CSVError(err error) [][]string {
	res := [][]string{
		{"error"},
		{err.Error()},
	}
	return res
}

func GenerateCSV(temps []daydegreesvc.Degree, params daydegreesvc.Params, autocomplete autocompletesvc.Autocomplete) [][]string {
	res := [][]string{}
	res = append(res, []string{"Indicator", getIndicator(params.Output)})
	res = append(res, []string{"Method", getMethod(params.DayCalc)})
	res = append(res, []string{"Base Temperature", fmt.Sprintf("%gC",params.Tb)})
	if params.Output == common.DDType {
		res = append(res, []string{"Room Temperature", getTR(params.Tr, params.Output)})
	}
	res = append(res, []string{"Unit", "Celsius"})
	res = append(res, []string{"Station", getStation(autocomplete)})
	res = append(res, []string{"Coordinates", fmt.Sprintf("%g, %g",autocomplete.Latitude, autocomplete.Longitude)})
	res = append(res, []string{"Description", getDescription(params)})
	res = append(res, []string{"Source", "https://energy-data.io/"})
	res = append(res, []string{""})
	if params.Output == common.DDType {
		res = append(res, []string{"Date", fmt.Sprintf("DD (%g,%g)",params.Tb, params.Tr)})
	} else if  params.Output ==  common.HDDType {
		res = append(res, []string{"Date",fmt.Sprintf("HDD (%g)",params.Tb)})
	} else if  params.Output ==  common.CDDType {
		res = append(res, []string{"Date",fmt.Sprintf("CDD (%g)",params.Tb)})
	}

	var line []string
	for _, v := range temps {
		line = []string{
			getFormattedDate(v.Date),
			getFormattedValue(v.Temp),
		}
		res = append(res, line)
	}
	return res
}

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

func getTR(tr float64, output string) string{
	if output != common.DDType {
		return "---"
	}
	return fmt.Sprintf("%gC",tr)
}

func getStation(autocomplete autocompletesvc.Autocomplete) string {
	return fmt.Sprintf("%s, %s, %s", autocomplete.ID, autocomplete.CityNameNative, autocomplete.CountryNameNative)
}

func getDescription(params daydegreesvc.Params ) string {
	method := getMethodDescription(params.DayCalc)
	if params.Output == common.HDDType {
		return fmt.Sprintf(
			"Heating Degree Days with a base temperature of %gC based on %s",
			params.Tb,
			method,
			)
	}

	if params.Output == common.CDDType {
		return fmt.Sprintf(
			"Cooling Degree Days with a base temperature of %gC based on %s",
			params.Tb,
			method,
		)
	}

	if params.Output == common.DDType {
		return fmt.Sprintf(
			"Room Heating Degree Days (Gradtagzahl) with a base temperature of %gC and a room temperature of %gC based on %s",
			params.Tb,
			params.Tr,
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
	//return strings.Replace(value, ".", ",", -1)
	return value
}

func getFormattedDate(date string) string{
	return strings.Replace(date, "-", "/", -1)
}
