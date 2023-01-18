package config

import (
	"im/help"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
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

func AnalyConfig() (*Config, error) {
	config := &Config{}
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		help.VczsLog("read ymal file failed", err)
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		help.VczsLog("ymal data unmarshal failed", err)
		return nil, err
	}
	return config, nil
}
