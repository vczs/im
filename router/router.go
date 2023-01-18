package router

import (
	"fmt"
	"im/config"
	"im/controller"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	// 发送验证码
	r.POST("/login", controller.Login)

	r.Run(fmt.Sprintf(":%d", config.Config.Port))
}
