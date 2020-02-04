package config

import (
	"fmt"
	"github.com/ghodss/yaml"
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
	CheckwxKey string `yaml:"checkwx_key"`
	UrlNoaa    string `yaml:"url_noaa"`
}

type Clients struct {
	AlertAddr 		string `yaml:"alert_addr"`
	StationsAddr 	string `yaml:"stations_addr"`
}

type NOAAScraperConfig struct {
	Server   		ServerConfig	`yaml:"server"`
	Database 		DatabaseConfig	`yaml:"database"`
	Sources  		SourcesConfig	`yaml:"sources"`
	Clients  		Clients			`yaml:"clients"`
	AlertsEnable  	bool			`yaml:"alerts_enable"`
}

func LoadConfig(path string) (config *NOAAScraperConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &NOAAScraperConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}

func (tc *NOAAScraperConfig)GetGRPCAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortGRPC)
}

func (tc *NOAAScraperConfig)GetHTTPAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTP)
}
