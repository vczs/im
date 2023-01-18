package config

import (
	"im/help"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var Config *config

type config struct {
	Port  int
	Redis struct {
		Host string
		Port int
	}
	Mongo struct {
		Host     string
		Port     int
		Name     string
		Password string
		Database string
	}
}

func Init() {
	config := &config{}
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		help.VczsLog("read ymal file failed", err)
		return
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		help.VczsLog("ymal data unmarshal failed", err)
		return
	}
	Config = config
}
