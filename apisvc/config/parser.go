package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ServerConfig struct {
	PortHTTPS int `yaml:"https_port"`
	PortHTTP  int `yaml:"http_port"`
}

type Clients struct {
	DailyAddr string `yaml:"daily_addr"`
}

type Security struct {
	Cert string	`yaml:"cert"`
	Key  string	`yaml:"key"`
}

type StationsConfig struct {
	Server   ServerConfig	`yaml:"server"`
	Clients  Clients		`yaml:"clients"`
	Security Security		`yaml:"security"`
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

func (tc *StationsConfig)GetHTTPAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTP)
}

func (tc *StationsConfig)GetHTTPSAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTPS)
}
