package middleware

import (
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/pkg/enums/code"
	"net/http"
)

// NoAuthMiddleware 无权限中间件(登录后不允许访问)
func NoAuthMiddleware() core.HandlerFunc {
	return func(c *core.Context) {
		if !c.IsLogin { // 中断后续中间件和 handler 执行！
			c.Next()
			return
		}
		responseJsonOrHtml(c, code.RequestForbidden, http.StatusForbidden)
	}
}
