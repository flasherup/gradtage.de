package utils

import (
	"fmt"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
)

func CSVError(err error) [][]string {
	res := [][]string{
		{"error"},
		{err.Error()},
	}
	return res
}

func GenerateCSV(temps []daydegreesvc.Degree, params daydegreesvc.Params, autocomplete autocompletesvc.Autocomplete) [][]string {
	res := generateCSVHeader(params, autocomplete)
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

func GenerateAvgCSV(temps []daydegreesvc.Degree, average []daydegreesvc.Degree, params daydegreesvc.Params, autocomplete autocompletesvc.Autocomplete) [][]string {
	res := generateCSVHeader(params, autocomplete)
	if params.Output == common.DDType {
		res = append(res, []string{"Date", fmt.Sprintf("DD (%g,%g)",params.Tb, params.Tr), "Average"})
	} else if  params.Output ==  common.HDDType {
		res = append(res, []string{"Date",fmt.Sprintf("HDD (%g)",params.Tb), "Average"})
	} else if  params.Output ==  common.CDDType {
		res = append(res, []string{"Date",fmt.Sprintf("CDD (%g)",params.Tb), "Average"})
	}

	var line []string
	avgLen := len(average)
	for _, v := range temps {
		doy := getAvgIndex(v.Date, params.Breakdown, params.WeekStart)
		avg := "---"
		if avgLen >= doy && doy > 0{
			avg = getFormattedValue(average[doy-1].Temp)
		}

		line = []string{
			getFormattedDate(v.Date),
			getFormattedValue(v.Temp),
			avg,
		}
		res = append(res, line)
	}
	return res
}

func generateCSVHeader(params daydegreesvc.Params, autocomplete autocompletesvc.Autocomplete) [][]string {
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
	return res;
}
