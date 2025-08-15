package middleware

import (
	"gen_gin_tpl/internal/core"
	"net/http"
)

// NoAuthMiddleware 无权限中间件(登录后不允许访问)
func NoAuthMiddleware() core.HandlerFunc {
	return func(c *core.Context) {
		if c.IsLogin {
			c.Redirect(http.StatusFound, "/") // 302 跳转更常用
			c.Abort()                         // 中断后续中间件和 handler 执行！
			return
		}
	}
}
