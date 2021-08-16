package parser

import (
	"encoding/json"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"time"
)
//const checkWXTimeTemplate = "2006-01-02T15:04:05.000Z"
const meteostatTimeTemplate = "2006-01-02 15:04:05"

type Meta struct {
	Source string `json:"source"`
}

type StationMetiostatData struct {
	Time 			string 	`json:"time"`
	Temperature 	float64 `json:"temperature"`
	TemperatureMin 	float64 `json:"temperature_min"`
	TemperatureMax 	float64 `json:"temperature_max"`
	Precipitation 	float64 `json:"precipitation"`
	Snowfall 		string  `json:"snowfall"`
	SnowDepth 		float64 `json:"snowdepth"`
	WindDirection 	int     `json:"winddirection"`
	WindSpeed 		float64 `json:"windspeed"`
	PeakGust 		string  `json:"peakgust"`
	Sunshine 		string  `json:"sunshine"`
	Pressure 		float64 `json:"pressure"`
}

type Result struct {
	Meta Meta                 	`json:"meta"`
	Data []StationMetiostatData `json:"data"`
}

func ParseMeteostat(data *[]byte) (*[]hourlysvc.Temperature, error) {
	var parsed Result
	err := json.Unmarshal(*data, &parsed)
	if err != nil {
		return nil, err
	}

	if len(parsed.Data) == 0 {
		return nil, nil
	}

	res := make([]hourlysvc.Temperature, len(parsed.Data))
	for i,st := range parsed.Data {
		date, err := time.Parse(meteostatTimeTemplate, st.Time)
		if err != nil {
			return nil, err
		}

		pTime := date.Format(common.TimeLayout)

		res[i] = hourlysvc.Temperature{
			Date:pTime,
			Temperature:st.Temperature,
		}
	}



	return &res,nil
}
