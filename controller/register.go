package controller

import (
	"im/dao"
	"im/define"
	"im/help"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Sex      int    `json:"sex"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

func Register(c *gin.Context) {
	// 参数校验
	req := new(RegisterRequest)
	err := c.ShouldBind(req)
	if err != nil {
		help.VczsLog("param analyse error", err)
		Response(c, define.PARAMETER_ANAIYSIS_FAILED, "", nil)
		return
	}
	if req.Account == "" || req.Password == "" || req.Name == "" || req.Email == "" || req.Code == "" {
		Response(c, define.PARAMETER_WRONG, "", nil)
		return
	}
	// 检验验证码是否正确
	code, err := dao.Redis.Get(c, req.Email).Result()
	if err != nil || code != req.Code {
		Response(c, define.EMAIL_CODE_WRONG, "", nil)
		return
	}
	// 检验用户名重复
	num, err := dao.GetUserCountByAccount(req.Account)
	if err != nil {
		help.VczsLog("get user count by account error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	if num > 0 {
		Response(c, define.ACCOUNT_EXIST, "", nil)
		return
	}
	// 构造用户并存储
	uid := help.GenerateUuid()
	user := &dao.User{
		Uid:      uid,
		Account:  req.Account,
		Password: help.GetMd5(req.Password),
		Name:     req.Name,
		Avatar:   req.Avatar,
		Sex:      req.Sex,
		Email:    req.Email,
		Status:   1,
		Ct:       time.Now().Unix(),
		Ut:       time.Now().Unix(),
	}
	err = dao.InsertUser(user)
	if err != nil {
		help.VczsLog("insert user error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	Response(c, define.OK, "", map[string]string{"uid": uid})
}
