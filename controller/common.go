package controller

import (
	"im/define"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type BodyData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, code int, message string, data interface{}) {
	if data != nil {
		body := BodyData{Code: code, Message: message, Data: data}
		if message == "" {
			body.Message = define.Message(code)
		}
		c.JSON(http.StatusOK, body)
	} else {
		body := Body{Code: code, Message: message}
		if message == "" {
			body.Message = define.Message(code)
		}
		c.JSON(http.StatusOK, body)
	}
}
