package controller

import (
	"im/dao"
	"im/define"
	"im/help"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

func UserLogin(c *gin.Context) {
	req := new(LoginRequest)
	err := c.ShouldBind(req)
	if err != nil {
		help.VczsLog("param analyse error", err)
		Response(c, define.PARAMETER_ANAIYSIS_FAILED, "", nil)
	}
	if req.Account == "" || req.Password == "" {
		Response(c, define.ACCOUNT_OR_PASSWORD_EMPTY, "", nil)
		return
	}
	user, err := dao.GetUserByAccountPassword(req.Account, help.GetMd5(req.Password))
	if err != nil {
		help.VczsLog("get user error", err)
		Response(c, define.ACCOUNT_OR_PASSWORD_ERROR, "", nil)
		return
	}
	token, err := help.GenerateToken(user.Uid, user.Name, define.TokenExpire)
	if err != nil {
		help.VczsLog("generate token error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	refreshToken, err := help.GenerateToken(user.Uid, user.Name, define.RefreshTokenExpire)
	if err != nil {
		help.VczsLog("generate refreshToken error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	res := LoginResponse{Token: token, RefreshToken: refreshToken}
	Response(c, define.OK, "登陆成功!", res)
}
