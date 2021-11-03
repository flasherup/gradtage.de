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

type WeatherbitConfig struct {
	UrlWeatherBit            string          `yaml:"url_weather_bit"`
	KeyWeatherBit            string          `yaml:"key_weather_bit"`
	NumberOfRequestPerSecond int             `yaml:"number_of_request_per_second"`
	NumberOfRequestPerDay    int             `yaml:"number_of_request_per_day"`
	NumberOfDays             int             `yaml:"number_of_days"`
	NumberOfDaysPerRequest   int             `yaml:"number_of_days_per_request"`
	ClearDay                 int             `yaml:"clear_day"`
	ForUpdate                ForUpdateConfig `yaml:"for_update"`
}

type ForUpdateConfig struct {
	NumberOfDaysUpdate int `yaml:"number_of_days_update"`
	Weekday            int `yaml:"weekday"`
}

type DatabaseConfig struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Clients struct {
	AlertAddr    string `yaml:"alert_addr"`
	StationsAddr string `yaml:"stations_addr"`
}

type WeatherBitUpdateConfig struct {
	Server       ServerConfig     `yaml:"server"`
	Database     DatabaseConfig   `yaml:"database"`
	Weatherbit   WeatherbitConfig `yaml:"weatherbit"`
	Clients      Clients          `yaml:"clients"`
	AlertsEnable bool             `yaml:"alerts_enable"`
}

func LoadConfig(path string) (config *WeatherBitUpdateConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &WeatherBitUpdateConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}

func (tc *WeatherBitUpdateConfig) GetGRPCAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortGRPC)
}

func (tc *WeatherBitUpdateConfig) GetHTTPAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTP)
}
