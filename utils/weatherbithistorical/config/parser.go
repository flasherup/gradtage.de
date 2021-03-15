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

type WeatherBitConfig struct {
	Database 		DatabaseConfig	`yaml:"database"`
	Sources  		SourcesConfig	`yaml:"sources"`
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
