package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ServerConfig struct {
	PortHTTPS 	int `yaml:"https_port"`
	PortHTTP  	int `yaml:"http_port"`
}

type Clients struct {
	DailyAddr 			string `yaml:"daily_addr"`
	HourlyAddr 			string `yaml:"hourly_addr"`
	HoaaAddr 			string `yaml:"noaa_addr"`
	AutocompleteAddr 	string `yaml:"autocomplete_addr"`
	UserAddr 			string `yaml:"user_addr"`
	AlertAddr 			string `yaml:"alert_addr"`
	StationsAddr 		string `yaml:"stations_addr"`
}

type Security struct {
	Cert string	`yaml:"cert"`
	Key  string	`yaml:"key"`
}

type Static struct {
	Folder  string `yaml:"folder"`
}

type ApiConfig struct {
	Server   		ServerConfig	    `yaml:"server"`
	Clients  		Clients				`yaml:"clients"`
	Security 		Security			`yaml:"security"`
	Users 	 		map[string]string 	`yaml:"users"`
	AlertsEnable 	bool				`yaml:"alerts_enable"`
	Static			Static				`yaml:"static"`
	Domains 		[]string 			`yaml:"domains"`
}

func LoadConfig(path string) (config *ApiConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &ApiConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}

func (tc *ApiConfig)GetHTTPAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTP)
}

func (tc *ApiConfig)GetHTTPSAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTPS)
}
