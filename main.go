package main

import (
	"im/config"
	"im/model"
	"im/router"
)

func main() {
	c, err := config.AnalyConfig()
	if err != nil {
		return
	}
	err = model.Init(c)
	if err != nil {
		return
	}
	router.Router(c.Port)
}
