package router

import (
	"fmt"
	"im/config"
	"im/controller"
	"im/middleware"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	{
		// 发送邮箱验证码
		r.POST("/send/email/code", controller.SendEnailCode)
		// 用户登录
		r.POST("/login", controller.UserLogin)
	}

	// 用户认证
	auth := r.Group("/user", middleware.AuthCheck)
	{
		// 用户信息
		auth.POST("/info", controller.UserInfo)
	}

	r.Run(fmt.Sprintf(":%d", config.Config.Port))
}
