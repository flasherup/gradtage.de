package config

import (

	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DatabaseConfig struct {
	Name		string	`yaml:"name"`
	Host 		string 	`yaml:"host"`
	User 		string 	`yaml:"user"`
	Password 	string 	`yaml:"password"`
	Port 		int 	`yaml:"port"`
}

type SourcesConfig struct {
	UrlWeatherBit	string `yaml:"url_weather_bit"`
	KeyWeatherBit	string `yaml:"key_weather_bit"`
}

type WeatherConfig struct{
	NumberOfRequestPerSecond int `yaml:"number_of_request_per_second"`
	NumberOfRequestPerDay	 int `yaml:"number_of_request_per_day"`
	NumberOfYears 			 int `yaml:"number_of_years"`
	NumberOfDaysPerRequest   int `yaml:"number_of_days_per_request"`
}

type WeatherBitConfig struct {
	Database 		DatabaseConfig	`yaml:"database"`
	Sources  		SourcesConfig	`yaml:"sources"`
	WeatherBit      WeatherConfig   `yaml:"weatherbit"`
}

func LoadConfig(path string) (config *WeatherBitConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &WeatherBitConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}
