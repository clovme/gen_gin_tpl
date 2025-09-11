package middleware

import (
	"fmt"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/pkg/enums/code"
	"net/http"
	"time"
)

func AuthMiddleware() core.HandlerFunc {
	return func(c *core.Context) {
		// 检查登录状态
		if !c.IsLogin {
			if c.IsAjax {
				c.JsonSafeDesc(code.RequestUnauthorized, nil)
			} else {
				c.Redirect(http.StatusFound, fmt.Sprintf("/login.html?_t=%d", time.Now().Unix()))
			}
			c.Abort()
			return
		}
		c.Next()
	}
}
