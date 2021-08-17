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

type Clients struct {
	AlertAddr 		string `yaml:"alert_addr"`
	WeatherBitAddr 	string `yaml:"weather_bit_addr"`
}

type DayDegreeConfig struct {
	Server   		ServerConfig	`yaml:"server"`
	Clients  		Clients			`yaml:"clients"`
	AlertsEnable  	bool			`yaml:"alerts_enable"`
}

func LoadConfig(path string) (config *DayDegreeConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &DayDegreeConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}

func (tc *DayDegreeConfig)GetGRPCAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortGRPC)
}

func (tc *DayDegreeConfig)GetHTTPAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTP)
}
