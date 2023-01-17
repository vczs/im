package main

import (
	"im/config"
	"im/help"
	"im/model"
	"im/router"
)

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		help.VczsLog("init config error", err)
		return
	}
	model.Init(conf)
	router.Router(conf.Port)
}
