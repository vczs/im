package controller

import (
	"im/dao"
	"im/define"
	"im/help"
	"time"

	"github.com/gin-gonic/gin"
)

type AddContactsRequest struct {
	Account string `json:"account"`
}

func AddContacts(c *gin.Context) {
	// 参数校验
	req := new(AddContactsRequest)
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
		Response(c, define.NOT_ADD_YOURSELF, "", nil)
		return
	}
	// 判断是否为好友
	is, _ := dao.IsFriend(u1.Uid, u2.Uid)
	if is {
		Response(c, define.ALREADY_FRIEND, "", nil)
		return
	}
	// 构造数据
	room := &dao.Room{
		Rid:      help.GenerateUuid(),
		Uid:      u1.Uid,
		Describe: u1.Name + "、" + u2.Name,
		Status:   1,
		Ct:       time.Now().Unix(),
		Ut:       time.Now().Unix(),
	}
	urs := make([]interface{}, 0)
	ur1 := &dao.UserRoom{
		Uid:    u1.Uid,
		Rid:    room.Rid,
		Type:   define.UserRoomTypeAlone,
		Status: 1,
		Ct:     time.Now().Unix(),
		Ut:     time.Now().Unix(),
	}
	ur2 := &dao.UserRoom{
		Uid:    u2.Uid,
		Rid:    room.Rid,
		Type:   define.UserRoomTypeAlone,
		Status: 1,
		Ct:     time.Now().Unix(),
		Ut:     time.Now().Unix(),
	}
	urs = append(urs, ur1, ur2)
	// 数据持久化
	err = dao.InsertRoom(room)
	if err != nil {
		help.VczsLog("insert room error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	err = dao.InsertManyUserRoom(urs)
	if err != nil {
		help.VczsLog("insert many user_room error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	Response(c, define.OK, "添加好友成功!", nil)
}
