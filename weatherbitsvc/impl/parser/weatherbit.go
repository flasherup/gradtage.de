package parser

import (
	"encoding/json"
	"fmt"
)



type Weather struct {
	Icon        string `json:"icon"`
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Data struct {
	Rh           float64 `json:"rh"`
	Pod          string  `json:"pod"`
	Lon          string `json:"lon"`
	Pres         float64 `json:"pres"`
	Timezone     string  `json:"timezone"`
	OnTime		 string  `json:"on_time"`
	ObTime       string  `json:"ob_time"`
	CountryCode  string  `json:"country_code"`
	Clouds       float64 `json:"clouds"`
	TS           int     `json:"ts"`
	SolarRad     float64 `json:"solar_rad"`
	StateCode    string  `json:"state_code"`
	CityName     string  `json:"city_name"`
	WindSPD      float64 `json:"wind_spd"`
	WindCdirFull string  `json:"wind_cdir_full"`
	WindCdir     string  `json:"wind_cdir"`
	SLP          float64 `json:"slp"`
	Vis          float64 `json:"vis"`
	HAngle       float64 `json:"h_angle"`
	Sunset       string  `json:"sunset"`
	DNI          float64 `json:"dni"`
	Dewpt        float64 `json:"dewpt"`
	Snow         float64  `json:"snow"`
	UV           float64 `json:"uv"`
	Precip       float64 `json:"precip"`
	WindDir      float64 `json:"wind_dir"`
	Sunrise      string  `json:"sunrise"`
	GHI          float64 `json:"ghi"`
	DHI          float64 `json:"dhi"`
	AQI          float64 `json:"aqi"`
	Lat          string  `json:"lat"`
	Weather      Weather `json:"weather"`
	DateTime     string  `json:"date_time"`
	Temp         float64 `json:"temp"`
	Station      string  `json:"station"`
	ElevAngle    float64 `json:"elev_angle"`
	AppTemp      float64 `json:"app_temp"`
}

type Minutely struct {
	TimestampUTC   string  `json:"timestamp_utc"`
	Snow           float64     `json:"snow"`
	TimestampLocal string  `json:"timestamp_local"`
	TS             int     `json:"ts"`
	Precip         float64 `json:"precip"`
}

type WeatherBitData struct {
	Count    int        `json:"count"`
	Data     []Data     `json:"data"`
	Minutely []Minutely `json:"minutely"`
}


func ParseWeatherBit(data *[]byte) (*WeatherBitData, error) {
	fmt.Println(string(*data))
	var wbd WeatherBitData
	err := json.Unmarshal(*data, &wbd)
	if err != nil {
		return nil, err
	}

	return &wbd,nil
}
