package parser

import (
	"encoding/json"
	"fmt"
)

type WeatherBitData struct {
	Count int `json:"count"`
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
