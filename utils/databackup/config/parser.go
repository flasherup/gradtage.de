package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Clients struct {
	SRCAddr 		string `yaml:"src_addr"`
}

type DatabaseConfig struct {
	Name		string	`yaml:"name"`
	Host 		string 	`yaml:"host"`
	Port 		int 	`yaml:"port"`
	User 		string 	`yaml:"user"`
	Password 	string 	`yaml:"password"`
}

type WeatherBitConfig struct {
	Clients		Clients 		`yaml:"clients"`
	Database 	DatabaseConfig	`yaml:"database"`
}

func LoadConfig(path string) (config *WeatherBitConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &WeatherBitConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}