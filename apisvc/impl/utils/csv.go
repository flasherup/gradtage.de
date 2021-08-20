package utils

import (
	"fmt"
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

func GenerateCSV(temps []daydegreesvc.Degree, params daydegreesvc.Params) [][]string {
	res := [][]string{}
	res = append(res, []string{"Indicator:", params.Output})
	res = append(res, []string{"Method:", params.DayCalc})
	res = append(res, []string{"Unit:", "Celsius"})
	res = append(res, []string{"TB:", fmt.Sprintf("%g°C",params.Tb)})
	res = append(res, []string{"TR:", fmt.Sprintf("%g°C",params.Tr)})
	res = append(res, []string{"Description:", getDescription(params)})
	res = append(res, []string{"",""})
	if params.Output == common.DDType {
		res = append(res, []string{"Date", fmt.Sprintf("DD %gC %gC",params.Tb, params.Tr)})
	} else if  params.Output ==  common.HDDType {
		res = append(res, []string{"Date",fmt.Sprintf("HDD %gC",params.Tb)})
	} else if  params.Output ==  common.CDDType {
		res = append(res, []string{"Date",fmt.Sprintf("CDD %gC",params.Tb)})
	}

	var line []string
	for _, v := range temps {
		line = []string{
			v.Date,
			getFormattedValue(v.Temp),
		}
		res = append(res, line)
	}
	return res
}

func getDescription(params daydegreesvc.Params ) string {
	method := getMethodDescription(params.DayCalc)
	if params.Output == common.HDDType {
		return fmt.Sprintf(
			"Heating Degree Days with a base temperature of %g°C based on %s",
			params.Tb,
			method,
			)
	}

	if params.Output == common.CDDType {
		return fmt.Sprintf(
			"Cooling Degree Days with a base temperature of %g°C based on %s",
			params.Tb,
			method,
		)
	}

	if params.Output == common.DDType {
		return fmt.Sprintf(
			"Room Heating Degree Days with a base temperature of %g°C and a room temperature of %g°C based on %s",
			params.Tb,
			params.Tr,
			method,
		)
	}

	return ""
}

func getMethodDescription(method string) string {
	if method == common.DayCalcInt {
		return "Integration Output"
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
	return strings.Replace(value, ".", ",", -1)
}
