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
	Name		string	`yaml:"name"`
	Host 		string 	`yaml:"host"`
	Port 		int 	`yaml:"port"`
	User 		string 	`yaml:"user"`
	Password 	string 	`yaml:"password"`
}

type Clients struct {
	AlertAddr 		string `yaml:"alert_addr"`
}

type Plans struct {
	FreeDefault string `yaml:"free_default"`
}

type UsersConfig struct {
	Server   		ServerConfig	`yaml:"server"`
	Database 		DatabaseConfig	`yaml:"database"`
	Clients  		Clients			`yaml:"clients"`
	AlertsEnable  	bool			`yaml:"alerts_enable"`
	Plans 			Plans			`yaml:"plans"`
}

func LoadConfig(path string) (config *UsersConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &UsersConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}

func (tc *UsersConfig)GetGRPCAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortGRPC)
}

func (tc *UsersConfig)GetHTTPAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTP)
}
