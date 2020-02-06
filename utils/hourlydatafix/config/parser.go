package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)


type DatabaseConfig struct {
	Name		string	`yaml:"name"`
	Host 		string 	`yaml:"host"`
	Port 		int 	`yaml:"port"`
	User 		string 	`yaml:"user"`
	Password 	string 	`yaml:"password"`
}

type HourlyFixConfig struct {
	Database 		DatabaseConfig	`yaml:"database"`
}

func LoadConfig(path string) (config *HourlyFixConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &HourlyFixConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}
