package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ServerConfig struct {
	PortHTTPS 	int `yaml:"https_port"`
	PortHTTP  	int `yaml:"http_port"`
	PortHTTPTest  	int `yaml:"http_test_port"`
	PortHTTPStatic 	int `yaml:"http_static_port"`
}

type Clients struct {
	AutocompleteAddr 	string `yaml:"autocomplete_addr"`
	UserAddr 			string `yaml:"user_addr"`
	AlertAddr 			string `yaml:"alert_addr"`
	StationsAddr 		string `yaml:"stations_addr"`
	DaydegreeAddr 		string `yaml:"daydegree_addr"`
	WeatherbitAddr 		string `yaml:"weatherbit_addr"`
}

type Static struct {
	Folder  string `yaml:"folder"`
}

type Woocommerce struct {
	Key 			string 		`json:"key"`
	Secret 			string 		`json:"secret"`
	WHSecret 		string 		`json:"whsecret"`
}

type ApiConfig struct {
	Server   		ServerConfig	    `yaml:"server"`
	Clients  		Clients				`yaml:"clients"`
	AlertsEnable 	bool				`yaml:"alerts_enable"`
	Static			Static				`yaml:"static"`
	Domains 		[]string 			`yaml:"domains"`
	Woocommerce 	Woocommerce 		`yaml:"woocommerce"`
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

func (tc *ApiConfig)GetHTTPTestAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTPTest)
}

func (tc *ApiConfig)GetHTTPStaticAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTPStatic)
}
