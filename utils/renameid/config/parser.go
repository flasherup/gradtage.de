package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)


type Rename struct {
	CurrentID	string	`yaml:"current_id"`
	NewID 		string 	`yaml:"new_id"`
}

type Services struct {
	HourlyUrl	string 	`yaml:"hourly_url"`
	DailyUrl	string 	`yaml:"daily_url"`
	StationsUrl string	`yaml:"stations_url"`
}

type RenameIDConfig struct {
	Renames 		[]Rename 	`yaml:"renames"`
	Services		Services	`yaml:"services"`
}

func LoadConfig(path string) (config *RenameIDConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &RenameIDConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}
