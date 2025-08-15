package core

import (
	"gen_gin_tpl/internal/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Context 自定义gin.Context
type Context struct {
	*gin.Context
	isContextEncrypted bool
	IsLogin            bool
	Session            Session
	UserInfo           *models.User
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
		Context: ctx,
		Session: Session{
			session:     sessions.Default(ctx),
			isDebugging: gin.Mode() == gin.TestMode || gin.Mode() == gin.DebugMode,
		},
	}
}
