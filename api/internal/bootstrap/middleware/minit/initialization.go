package minit

import (
	"fmt"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/constants"
	httpLog "gen_gin_tpl/pkg/logger/http"
	"github.com/gin-gonic/gin"
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
func isAjax(c *gin.Context) bool {
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

func getContextUserInfo(c *core.Context, userID int64) {
	if user, err := query.Q.User.Where(query.User.ID.Eq(userID)).First(); err == nil {
		c.Set(constants.ContextUserInfo, user)
		c.Set(constants.IsContextLogin, true)
		return
	}
	c.Set(constants.IsContextLogin, false)
	c.Set(constants.HttpLogKey, "User 不存在，删除Session会话标识")
	c.Session.ClearSession() // 建议清理 session
}

// setContextUserInfo 设置上下文用户信息
func setContextUserInfo(c *core.Context) {
	userID, ok, isToken := c.Session.GetUserID(c.Context)

	if ok && !isToken { // 非Token请求
		getContextUserInfo(c, userID)
		return
	} else if ok && isToken { // Token请求，判断是否过期
		if token, err := query.Q.Token.Where(query.Q.Token.UserID.Eq(userID)).First(); err == nil {
			if tokenUpdate(token, c) {
				getContextUserInfo(c, userID)
				return
			}
		}
		c.Session.ClearSession()
	}
	c.Set(constants.IsContextLogin, false)
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
		// 1. 限制本地 IP 访问
		//host := strings.SplitN(c.Request.Host, ":", 2)[0]
		//if net.ParseIP(host) != nil {
		//	c.AbortWithStatus(http.StatusForbidden)
		//	return
		//}

		// 2. 基础上下文初始化
		c.Set(constants.IsContextLogin, false) // 默认未登录
		c.Set(constants.IsContextAjax, isAjax(c.Context))
		c.Header("Client-ID", c.Session.BrowserClientID())

		// 3. 设置用户信息（可能成功，也可能失败）
		setContextUserInfo(c)

		// 4. 公共资源：即便没登录，也直接放行（页面可以自己判断是否有用户信息）
		if strings.HasPrefix(c.Router.Path(c.Request.URL.Path).Group, "public") {
			c.Set(constants.HttpLogKey, "公共资源")
		}

		// 7. 放行
		c.Next()
	}
}
