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

type TestConfig struct {
	Server   ServerConfig	`yaml:"server"`
}

func LoadConfig(path string) (config *TestConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &TestConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}

func (tc *TestConfig)GetHTTPAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTP)
}
