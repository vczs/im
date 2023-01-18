package middleware

import (
	"im/define"
	"im/help"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthCheck(c *gin.Context) {
	token := c.GetHeader("token")
	user, err := help.AnalyseToken(token)
	if err != nil {
		help.VczsLog("auth analyse token failed", err)
		c.Abort()
		c.JSON(http.StatusOK, gin.H{
			"code":    define.TOKEN_INVALID,
			"messgae": define.Message(define.TOKEN_INVALID),
		})
		return
	}
	c.Set("user", user)
	c.Next()
}
