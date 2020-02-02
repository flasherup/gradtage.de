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

type SourcesConfig struct {
	CheckwxKey		string	`yaml:"checkwx_key"`
	UrlDWD			string  `yaml:"url_dwd"`
}

type Clients struct {
	AlertAddr 		string `yaml:"alert_addr"`
	StationsAddr 	string `yaml:"stations_addr"`
	HourlyAddr 		string `yaml:"hourly_addr"`
}

type HrlAggregatorConfig struct {
	Server   		ServerConfig	`yaml:"server"`
	Sources  		SourcesConfig	`yaml:"sources"`
	Clients  		Clients			`yaml:"clients"`
	AlertsEnable  	bool			`yaml:"alerts_enable"`
}

func LoadConfig(path string) (config *HrlAggregatorConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &HrlAggregatorConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}

func (tc *HrlAggregatorConfig)GetGRPCAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortGRPC)
}

func (tc *HrlAggregatorConfig)GetHTTPAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTP)
}
