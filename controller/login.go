package controller

import (
	"im/define"
	"im/help"
	"im/model"

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

func Login(c *gin.Context) {
	req := new(LoginRequest)
	err := c.ShouldBind(req)
	if err != nil {
		Response(c, define.PARAMETER_ANAIYSIS_FAILED, "", nil)
	}
	if req.Account == "" || req.Password == "" {
		Response(c, define.ACCOUNT_OR_PASSWORD_EMPTY, "", nil)
		return
	}
	user, err := model.GetUserBasicByAccountPassword(req.Account, help.GetMd5(req.Password))
	if err != nil {
		Response(c, define.ACCOUNT_OR_PASSWORD_ERROR, "", nil)
		return
	}
	token, err := help.GenerateToken(user.Uid, user.Email, define.TokenExpire)
	if err != nil {
		Response(c, -1, err.Error(), nil)
		return
	}
	refreshToken, err := help.GenerateToken(user.Uid, user.Email, define.RefreshTokenExpire)
	if err != nil {
		Response(c, -1, err.Error(), nil)
		return
	}
	res := LoginResponse{Token: token, RefreshToken: refreshToken}
	Response(c, define.OK, "登陆成功!", res)
}
