package controller

import (
	"im/dao"
	"im/define"
	"im/help"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var socket = make(map[string]*websocket.Conn)

type MessageStruct struct {
	Rid     string `json:"rid"`
	Message string `json:"message"`
}

func Comm(c *gin.Context) {
	// 获取conn
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		help.VczsLog("websocket upgrade error", err)
		Response(c, -1, err.Error(), nil)
		return
	}
	defer conn.Close()
	// 获取用户信息
	uc, has := c.Get("user")
	if !has {
		return
	}
	user := uc.(*help.UserClaim)
	// 将用户添加到socket中
	socket[user.Uid] = conn
	// 用户消息处理
	for {
		// 解析消息
		ms := new(MessageStruct)
		err = conn.ReadJSON(ms)
		if err != nil {
			help.VczsLog("read message error", err)
			return
		}
		// 判断用户是否属于目标房间
		num, err := dao.GetUserRoomByUidRid(user.Uid, ms.Rid)
		if err != nil {
			help.VczsLog("get user room by uid and rid error", err)
			return
		}
		if num < 1 {
			conn.WriteJSON(map[string]interface{}{"code": -1, "message": define.Message(define.ACCESS_DENIED)})
			continue
		}
		// 保存消息
		message := &dao.Message{
			Mid:    help.GenerateUuid(),
			Uid:    user.Uid,
			Rid:    ms.Rid,
			Data:   ms.Message,
			Type:   1,
			Status: 1,
			Ct:     time.Now().Unix(),
			Ut:     time.Now().Unix(),
		}
		err = dao.InsertMessage(message)
		if err != nil {
			help.VczsLog("insert message error", err)
			return
		}
		// 获取该房间的所有用户
		urs, err := dao.FindUserRoomByRid(ms.Rid)
		if err != nil {
			help.VczsLog("get user room by rid error", err)
			return
		}
		// 过滤在线用户中非该房间用户
		for _, ur := range urs {
			if ur.Uid == user.Uid {
				continue
			}
			if co, ok := socket[ur.Uid]; ok {
				err = co.WriteMessage(websocket.TextMessage, []byte(ms.Message))
				if err != nil {
					help.VczsLog("write message error", err)
					return
				}
			}
		}
	}
}
