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
	CheckwxKey		string	`yaml:"checkwxKey"`
}

type Clients struct {
	StationsAddr string `yaml:"stations_addr"`
	HourlyAddr string `yaml:"hourly_addr"`
}

type StationsConfig struct {
	Server   ServerConfig	`yaml:"server"`
	Sources  SourcesConfig	`yaml:"sources"`
	Clients  Clients		`yaml:"clients"`
}

func LoadConfig(path string) (config *StationsConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &StationsConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}

func (tc *StationsConfig)GetGRPCAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortGRPC)
}

func (tc *StationsConfig)GetHTTPAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTP)
}
