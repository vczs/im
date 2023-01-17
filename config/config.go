package config

import (
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

func InitConfig() (Config, error) {
	config := Config{}
	data, err := ioutil.ReadFile("etc/config.yaml")
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
