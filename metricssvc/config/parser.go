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
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Clients struct {
	AlertAddr    string `yaml:"alert_addr"`
	WeatherbitAddr string `yaml:"weatherbit_addr"`
}

type MetricsConfig struct {
	Server       ServerConfig     `yaml:"server"`
	Database     DatabaseConfig   `yaml:"database"`
	Clients      Clients          `yaml:"clients"`
	AlertsEnable bool             `yaml:"alerts_enable"`
}

func LoadConfig(path string) (config *MetricsConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &MetricsConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}

func (tc *MetricsConfig) GetGRPCAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortGRPC)
}

func (tc *MetricsConfig) GetHTTPAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTP)
}
