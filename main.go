package main

import (
	"im/config"
	"im/dao"
	"im/router"
)

func main() {
	config.Init()
	dao.Init()
	router.Router()
}
