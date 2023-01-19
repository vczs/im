package controller

import (
	"im/dao"
	"im/define"
	"im/help"

	"github.com/gin-gonic/gin"
)

type DeleteContactsRequest struct {
	Account string `json:"account"`
}

func DeleteContacts(c *gin.Context) {
	// 参数校验
	req := new(DeleteContactsRequest)
	err := c.ShouldBind(req)
	if err != nil {
		help.VczsLog("param analyse error", err)
		Response(c, define.PARAMETER_ANAIYSIS_FAILED, "", nil)
		return
	}
	if req.Account == "" {
		Response(c, define.PARAMETER_WRONG, "", nil)
		return
	}
	// 获取用户信息
	uc, has := c.Get("user")
	if !has {
		Response(c, -1, "", nil)
		return
	}
	u1 := uc.(*help.UserClaim)
	// 查询目标用户
	u2, err := dao.GetUserByAccount(req.Account)
	if err != nil {
		Response(c, define.USER_NOT_EXIST, "", nil)
		return
	}
	if u1.Uid == u2.Uid {
		Response(c, define.NOT_DELETE_YOURSELF, "", nil)
		return
	}
	// 判断是否为好友
	is, rid := dao.IsFriend(u1.Uid, u2.Uid)
	if !is {
		Response(c, define.NOT_FRIEND, "", nil)
		return
	}
	// 数据持久化
	err = dao.DeleteRoomByRid(rid)
	if err != nil {
		help.VczsLog("delete room error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	err = dao.DeleteUserRoomByRid(rid)
	if err != nil {
		help.VczsLog("delete user_room error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	Response(c, define.OK, "删除好友成功!", nil)
}
