package config

import (

	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Clients struct {
	SRCAddr 		string `yaml:"src_addr"`
	RCVRAddr 		string `yaml:"rcvr_addr"`
}

type WeatherBitConfig struct {
	Clients		Clients `yaml:"clients"`
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
