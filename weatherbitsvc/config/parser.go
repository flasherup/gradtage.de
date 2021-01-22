package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ServerConfig struct {
	PortHTTP int `yaml:"http_port"`
	PortGRPC int `yaml:"grpc_port"`
}

type DatabaseConfig struct {
	Name		string	`yaml:"name"`
	Host 		string 	`yaml:"host"`
	Port 		int 	`yaml:"port"`
	User 		string 	`yaml:"user"`
	Password 	string 	`yaml:"password"`
}

type SourcesConfig struct {
	UrlWeatherBit	string `yaml:"url_weather_bit"`
	KeyWeatherBit	string `yaml:"key_weather_bit"`
}

type Clients struct {
	AlertAddr 		string `yaml:"alert_addr"`
}

type WeatherBitConfig struct {
	Server   		ServerConfig	`yaml:"server"`
	Database 		DatabaseConfig	`yaml:"database"`
	Sources  		SourcesConfig	`yaml:"sources"`
	Clients  		Clients			`yaml:"clients"`
	AlertsEnable  	bool			`yaml:"alerts_enable"`
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

func (tc *WeatherBitConfig)GetGRPCAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortGRPC)
}

func (tc *WeatherBitConfig)GetHTTPAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTP)
}
