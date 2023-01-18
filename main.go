package main

import (
	"im/config"
	"im/model"
	"im/router"
)

func main() {
	config.Init()
	model.Init()
	router.Router()
}
