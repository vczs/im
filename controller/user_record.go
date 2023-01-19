package controller

import (
	"im/dao"
	"im/define"
	"im/help"

	"github.com/gin-gonic/gin"
)

type RecordRequest struct {
	Rid   string `json:"rid"`
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
}

func Record(c *gin.Context) {
	// 解析参数
	req := new(RecordRequest)
	err := c.ShouldBind(req)
	if err != nil {
		help.VczsLog("param analyse error", err)
		Response(c, define.PARAMETER_ANAIYSIS_FAILED, "", nil)
		return
	}
	// 参数校验
	rid, page, limit := req.Rid, req.Page, req.Limit
	if rid == "" {
		Response(c, 20002, "rid is empty!", nil)
		return
	}
	// 获取用户信息
	uc, has := c.Get("user")
	if !has {
		help.VczsLog("user not exit!", nil)
		return
	}
	user := uc.(*help.UserClaim)
	// 查询用户是否属于目标房间
	num, err := dao.GetUserRoomCountByUidRid(user.Uid, rid)
	if err != nil {
		help.VczsLog("get user room count by uid and rid error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	if num < 1 {
		Response(c, define.ACCESS_DENIED, "", nil)
		return
	}
	// 处理参数
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = define.DefaultLimit
	}
	skip := (page - 1) * limit
	// 查询message
	mes, err := dao.FindMessageByRid(rid, &skip, &limit)
	if err != nil {
		help.VczsLog("find message by rid error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	Response(c, define.OK, "", mes)
}
