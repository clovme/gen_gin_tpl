package middleware

import (
	"fmt"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/constants"
	httpLog "gen_gin_tpl/pkg/logger/http"
	"strings"
	"time"
)

// isAjax 判断是否是Ajax请求
//
// 参数:
//   - c: 上下文对象
//
// 返回值:
//   - bool: 是否是Ajax请求，true 是 Ajax 请求，false 不是 Ajax 请求
func isAjax(c *core.Context) bool {
	referer := c.GetHeader("Referer")
	if referer != "" {
		referer = strings.SplitN(referer, "/", 3)[2]
		referer = strings.Split(referer, "/")[0]
	}
	protocol := "http"
	if c.Request.TLS != nil {
		protocol = "https"
	}
	accept := c.GetHeader("Accept") == "*/*" || strings.Contains(c.GetHeader("Accept"), "json")
	hostReferer := c.Request.Host == referer && c.GetHeader("Referer") != fmt.Sprintf("%s://%s%s", protocol, c.Request.Host, c.Request.RequestURI)
	xmlHttpRequest := c.GetHeader("X-Requested-With") == "XMLHttpRequest"
	return hostReferer && xmlHttpRequest && accept
}

// setContextUserInfo 设置上下文用户信息
func setContextUserInfo(userID int64, ok bool, c *core.Context) {
	if !ok {
		return
	}
	if user, err := query.Q.User.Where(query.User.ID.Eq(userID)).First(); err == nil {
		c.Set(constants.ContextUserInfo, user)
		c.Set(constants.IsContextLogin, true)
		return
	}
	c.Set(constants.IsContextLogin, false)
	httpLog.Info(c.Context).Msg("User 不存在，删除Session会话标识")
}

// tokenUpdate 更新Token
func tokenUpdate(token *models.Token, c *core.Context) bool {
	if !token.Revoked {
		c.Set(constants.IsContextLogin, false)
		httpLog.Info(c.Context).Msg("Token 已失效，删除Session会话标识")
		return false
	}
	mapClaims, err := c.Session.ParseUserToken(token.Token)
	if err != nil {
		c.Set(constants.IsContextLogin, false)
		httpLog.Info(c.Context).Msg("Token 不存在，删除Session会话标识")
		return false
	}
	now := time.Now().Unix()
	iat := mapClaims["iat"].(int64)
	exp := mapClaims["exp"].(int64)

	if exp-now <= now-iat {
		_, _ = query.Q.Token.Where(query.Q.Token.ID.Eq(token.ID)).Update(query.Q.Token.Revoked, false)
		c.Set(constants.IsContextLogin, false)
		httpLog.Info(c.Context).Msg("Token 已过期，删除Session会话标识")
		return false
	}
	c.Set(constants.IsContextLogin, true)
	return true
}

// InitializationMiddleware 初始化中间件
func InitializationMiddleware() core.HandlerFunc {
	return func(c *core.Context) {
		c.Set(constants.IsContextLogin, false)
		c.Set(constants.IsContextAjax, isAjax(c))
		c.Header("Client-ID", c.Session.BrowserClientID())
		userID, ok, isToken := c.Session.GetUserID(c.Context)

		// 非Token请求，直接返回
		if ok && !isToken {
			setContextUserInfo(userID, ok, c)
		}
		// Token请求，判断是否过期
		if ok && isToken {
			if token, err := query.Q.Token.Where(query.Q.Token.UserID.Eq(userID)).First(); err == nil {
				if tokenUpdate(token, c) {
					setContextUserInfo(userID, ok, c)
				}
			}
		}
		c.Next()
	}
}
