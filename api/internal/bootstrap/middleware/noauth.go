package middleware

import (
	"gen_gin_tpl/pkg/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NoAuthMiddleware 无权限中间件(登录后不允许访问)
func NoAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if session.IsLogin(c) {
			c.Redirect(http.StatusFound, "/") // 302 跳转更常用
			c.Abort()                         // 中断后续中间件和 handler 执行！
			return
		}
		c.Next()
	}
}
