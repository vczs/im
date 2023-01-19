package controller

import (
	"im/dao"
	"im/define"
	"im/help"

	"github.com/gin-gonic/gin"
)

type SearchRequest struct {
	Account string `json:"account"`
}
type SearchResponse struct {
	Account  string `json:"account"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Sex      int    `json:"sex"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
	Register int64  `json:"register"`
	Friend   bool   `json:"friend"`
}

func Search(c *gin.Context) {
	// 参数校验
	req := new(SearchRequest)
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
	is, err := dao.IsFriend(u1.Uid, u2.Uid)
	if err != nil {
		help.VczsLog("is friend error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	// 返回结果
	resp := &SearchResponse{
		Account:  u2.Account,
		Name:     u2.Name,
		Avatar:   u2.Avatar,
		Sex:      u2.Sex,
		Email:    u2.Email,
		Status:   u2.Status,
		Register: u2.Ct,
		Friend:   is,
	}
	Response(c, define.OK, "", resp)
}
