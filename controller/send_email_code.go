package controller

import (
	"im/dao"
	"im/define"
	"im/help"
	"time"

	"github.com/gin-gonic/gin"
)

type SendEnailCodeRequest struct {
	Email string `json:"email"`
}

func SendEnailCode(c *gin.Context) {
	// 解析参数
	req := new(SendEnailCodeRequest)
	err := c.ShouldBind(req)
	if err != nil {
		help.VczsLog("param analyse error", err)
		Response(c, define.PARAMETER_ANAIYSIS_FAILED, "", nil)
		return
	}
	// 参数校验
	if req.Email == "" {
		Response(c, define.EMAIL_EMPTY, "", nil)
		return
	}
	// 检查是否已注册
	num, err := dao.GetUserCountByEmail(req.Email)
	if err != nil {
		help.VczsLog("get user count by email error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	if num > 0 {
		Response(c, define.EMAIL_HAS_REGISTERED, "", nil)
		return
	}
	// 限频
	down, err := dao.Redis.TTL(c, req.Email).Result()
	if err != nil {
		help.VczsLog("redis result error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	if down.Seconds() > 0 || down.Seconds() == -1 {
		Response(c, define.REQUEST_OFTEN, "", nil)
		return
	}

	// 生成验证码
	code := help.GenerateEmailCode()
	// 发送验证码
	err = help.SendEmailCode(req.Email, code)
	if err != nil {
		help.VczsLog("send email code error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	// 保存验证码
	err = dao.Redis.Set(c, req.Email, code, time.Second*time.Duration(define.CodeExpire)).Err()
	if err != nil {
		help.VczsLog("redis set email code error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	Response(c, define.OK, "发送成功!", nil)
}
