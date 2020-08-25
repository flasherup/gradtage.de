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

type EmailConfig struct {
	Host  		string   `yaml:"host"`
	Port  	 	string   `yaml:"port"`
	User  	 	string   `yaml:"user"`
	Pass  	 	string   `json:"pass"`
	From  	 	string   `json:"from"`
	Recipients  []string `json:"recipients"`
	EmailTemplates 	EmailTemplates  	`yaml:"email_templates"`
}

type EmailTemplates struct {
	UserPlanUpdate string `yaml:"user_plan_update"`
}


type AlertConfig struct {
	Server   		ServerConfig 		`yaml:"server"`
	EmailConfig 	EmailConfig  		`yaml:"email_config"`
}

func LoadConfig(path string) (config *AlertConfig, err error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config = &AlertConfig{}

	err = yaml.Unmarshal(c, config)
	if err != nil {
		return
	}

	return
}

func (tc *AlertConfig)GetGRPCAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortGRPC)
}

func (tc *AlertConfig)GetHTTPAddress() string {
	return fmt.Sprintf("%s:%d", "", tc.Server.PortHTTP)
}
