package router

import (
	"fmt"
	"im/controller"

	"github.com/gin-gonic/gin"
)

func Router(port int) {
	r := gin.Default()

	// 发送验证码
	r.POST("/login", controller.Login)

	r.Run(fmt.Sprintf(":%d", port))
}
