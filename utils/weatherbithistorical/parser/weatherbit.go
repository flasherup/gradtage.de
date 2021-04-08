package parser

import (
	"encoding/json"
)



type Weather struct {
	Icon        string `json:"icon"`
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Data struct {
	Rh           float64 `json:"rh"`
	Pod          string  `json:"pod"`
	Pres         float64 `json:"pres"`
	OnTime		 string  `json:"on_time"`
	ObTime       string  `json:"ob_time"`
	Clouds       float64 `json:"clouds"`
	TS           float64 `json:"ts"`
	SolarRad     float64 `json:"solar_rad"`
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

	Weather      Weather `json:"weather"`
	DateTime     string  `json:"date_time"`
	Temp         float64 `json:"temp"`
	ElevAngle    float64 `json:"elev_angle"`
	AppTemp      float64 `json:"app_temp"`
}

type WeatherBitData struct {
	Timezone    string  `json:"timezone"`
	CountryCode string  `json:"country_code"`
	StateCode   string  `json:"state_code"`
	CityName    string  `json:"city_name"`
	Lon         float64 `json:"lon"`
	Lat         float64 `json:"lat"`
	Station     string  `json:"station"`
	AQI         float64 `json:"aqi"`
	Count       int     `json:"count"`
	Data        []Data  `json:"data"`
}


func ParseWeatherBit(data *[]byte) (*WeatherBitData, error) {
	var wbd WeatherBitData
	err := json.Unmarshal(*data, &wbd)
	if err != nil {
		return nil, err
	}

	return &wbd,nil
}
