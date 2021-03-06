package parser

import (
	"encoding/json"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"time"
)
const checkWXTimeTemplate = "2006-01-02T15:04Z"

type Temperature struct {
	Celsius float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}

type Humidity struct {
	Percent float64 `json:"percent"`
}

type Elevation struct {
	Feet float64 `json:"feet"`
	Meters float64 `json:"meters"`
}

type Location struct {
	Coordinates []float64 `json:"coordinates"`
	Type string `json:"type"`
}

type Station struct {
	Name string `json:"name"`
}

type Clouds struct {
	Code string `json:"code"`
	Text string `json:"text"`
	BaseFeetAgl int `json:"base_feet_agl"`
	BaseMetersAgl int `json:""base_meters_agl""`
}

type Conditions struct {
	Prefix string `json:"prefix"`
	Code string `json:"code"`
	Text string `json:"text"`
}


type StationDataCheckWX struct {
	Temperature Temperature `json:"temperature"`
	PewPoint Temperature `json:"dewpoint"`
	Humidity Humidity `json:"humidity"`
	Elevation Elevation `json:"elevation"`
	Location Location `json:"location"`
	ICAO string `json:"icao"`
	Observed string `json:"observed"`
	RawText string `json:"raw_text"`
	Station Station `json:"station"`
	FlightCategory string `json:"flight_category"`
	Clouds []Clouds `json:"clouds"`
	Conditions []Conditions `json:"Conditions"`
}

type Result struct {
	Result int                `json:"result"`
	Data []StationDataCheckWX `json:"data"`
}

func ParseCheckWX(data *[]byte) (*hourlysvc.Temperature, error) {
	var parsed Result
	err := json.Unmarshal(*data, &parsed)
	if err != nil {
		return nil, err
	}

	if len(parsed.Data) == 0 {
		return nil, nil
	}

	st := parsed.Data[0]

	date, err := time.Parse(checkWXTimeTemplate, st.Observed)
	if err != nil {
		return nil, err
	}

	pTime := date.Format(common.TimeLayout)

	temp := hourlysvc.Temperature{
		Date:pTime,
		Temperature:st.Temperature.Celsius,
	}

	return &temp,nil
}
