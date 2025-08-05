package middleware

import (
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/session"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// SessionMiddleware 设置 session
func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientId := session.Get(c, constants.ClientID)
		if clientId == nil {
			clientId = base64Captcha.RandomId()
			session.Set(c, constants.ClientID, clientId)
		} else {
			session.Set(c, constants.ClientID, clientId)
		}
		c.Header("Client-ID", clientId.(string))
		c.Next()
	}
}
