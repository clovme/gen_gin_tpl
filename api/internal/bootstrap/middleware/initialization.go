package middleware

import (
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/pkg/constants"
	httpLog "gen_gin_tpl/pkg/logger/http"
	"github.com/mojocn/base64Captcha"
	"strings"
)

func setContextUserInfo(userID int64, ok bool, c *core.Context) {
	if !ok {
		return
	}
	if user, err := query.Q.User.Where(query.User.ID.Eq(userID)).First(); err == nil {
		c.Set(constants.ContextUserInfo, user)
		c.Set(constants.IsContextLogin, true)
		return
	}
	c.Session.DelUserSession()
	c.Set(constants.IsContextLogin, false)
	httpLog.Info(c.Context).Msg("User 不存在，删除Session会话标识")
}

// InitializationMiddleware 初始化中间件
func InitializationMiddleware() core.HandlerFunc {
	return func(c *core.Context) {
		if strings.HasSuffix(c.FullPath(), "*filepath") || strings.EqualFold(c.Request.URL.Path, "/favicon.ico") {
			c.Next()
			return
		}
		c.Set(constants.IsContextLogin, false)
		c.Header("Authorization", base64Captcha.RandomId())
		c.Header("Client-ID", c.Session.BrowserClientID())
		if userID, ok, isToken := c.Session.GetUserID(c.Context); ok {
			if isToken {
				// 处理token信息
			} else {
				c.Next()
				return
			}
			setContextUserInfo(userID, ok, c)
		}
		c.Next()
	}
}
