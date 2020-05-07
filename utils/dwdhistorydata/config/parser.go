package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)


type HrlAggregatorConfig struct {
	UrlDWD			string  `yaml:"url_dwd"`
	StationsAddr 	string `yaml:"stations_addr"`
	DailyAddr 		string `yaml:"daily_addr"`
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
