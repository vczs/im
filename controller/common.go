package controller

import (
	"im/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, code int, message string, data interface{}) {
	body := Body{Code: code, Message: message}
	if message == "" {
		body.Message = config.Message(code)
	}
	if data != nil {
		body.Data = data
	}
	c.JSON(http.StatusOK, body)
}
