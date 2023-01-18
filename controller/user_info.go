package controller

import (
	"im/dao"
	"im/define"
	"im/help"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	uc, has := c.Get("user")
	if !has {
		Response(c, -1, "", nil)
		return
	}
	user := uc.(*help.UserClaim)
	userInfo, err := dao.GetUserByUid(user.Uid)
	if err != nil {
		help.VczsLog("get user by uid error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	Response(c, define.OK, "", userInfo)
}
