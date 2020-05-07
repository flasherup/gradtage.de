package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type StationsConfig struct {
	DailyAddr 			string `yaml:"daily_addr"`
	StationsAddr string `yaml:"stations_addr"`
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
