package core

import (
	"gen_gin_tpl/internal/libs"
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/constants"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Context 自定义gin.Context
type Context struct {
	*gin.Context
	IsContextEncrypted bool
	IsLogin            bool // true 登录，false 未登录
	IsAjax             bool // true 是 Ajax 请求，false 不是 Ajax 请求
	Session            Session
	Router             routesMap
	UserInfo           *models.User
	Config             *libs.Config
}

// NewContext 创建自定义gin.Context
//
// 参数:
//   - ctx: gin.Context对象
//
// 返回值:
//   - *Context: 自定义gin.Context对象
//
// 说明:
//   - 创建自定义gin.Context对象，用于自定义路由和中间件。
func NewContext(ctx *gin.Context) *Context {
	return &Context{
		Context:  ctx,
		IsLogin:  ctx.GetBool(constants.IsContextLogin),
		IsAjax:   ctx.GetBool(constants.IsContextAjax),
		UserInfo: getUserInfo(ctx),
		Router:   initRoutesMap(),
		Config:   libs.WebConfig,
		Session: Session{
			session: sessions.Default(ctx),
		},
	}
}
