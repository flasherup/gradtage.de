package utils

import (
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/daydegreesvc"
)

type JSONHDD struct {
	Date string `json:"date"`
	Value float64 `json:"value"`
}

type JSONHDDAvg struct {
	Date string `json:"date"`
	Value float64 `json:"value"`
	Average float64 `json:"average"`
}

type JSONDataCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type JSONData struct {
	Indicator string `json:"indicator"`
	Method string `json:"method"`
	BaseTemperature string `json:"base_temperature"`
	RoomTemperature string `json:"room_temperature"`
	Unit string `json:"unit"`
	Station string `json:"station"`
	Coordinates JSONDataCoordinates `json:"coordinates"`
	Description string `json:"description"`
	Source string `json:"source"`
	Data interface{} `json:"data"`
}

func GenerateJSON(temps []daydegreesvc.Degree, params daydegreesvc.Params, autocomplete autocompletesvc.Autocomplete) *JSONData {
	res := generateJSONHeader(params, autocomplete)
	data := make([]JSONHDD, len(temps))
	for i, v := range temps {
		data[i] = JSONHDD{
			Date:getFormattedDate(v.Date),
			Value:v.Temp,
		}
	}
	res.Data = data

	return res
}

func GenerateAvgJSON(temps []daydegreesvc.Degree, average []daydegreesvc.Degree, params daydegreesvc.Params, autocomplete autocompletesvc.Autocomplete) *JSONData {
	res := generateJSONHeader(params, autocomplete)
	data := make([]JSONHDDAvg, len(temps))
	avgLen := len(average)
	for i, v := range temps {
		doy := getAvgIndex(v.Date, params.Breakdown, params.WeekStart)
		avg := 0.00
		if avgLen >= doy && doy > 0{
			avg = average[doy-1].Temp
		}
		data[i] = JSONHDDAvg{
			Date: v.Date,
			Value: v.Temp,
			Average: avg,
		}
	}
	res.Data = data

	return res
}

func generateJSONHeader(params daydegreesvc.Params, autocomplete autocompletesvc.Autocomplete) *JSONData {
	res := JSONData{
		Indicator: getIndicator(params.Output),
		Method: getMethod(params.DayCalc),
		BaseTemperature: getTB(params.Tb, params.Metric),
		Unit: getUnits(params.Metric),
		Station:  getStation(autocomplete),
		Coordinates: JSONDataCoordinates{autocomplete.Latitude, autocomplete.Longitude},
		Description: getDescription(params),
		Source: "https://energy-data.io/",
	}

	if params.Output == common.DDType {
		res.RoomTemperature = getTR(params.Tr, params.Output, params.Metric)
	}

	return &res;
}