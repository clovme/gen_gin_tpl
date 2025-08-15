package with

import (
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/constants"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// getUserInfo 获取用户信息
//
// 参数：
//   - ctx: Gin上下文对象
//
// 返回值：
//   - *models.User: 用户信息
func getUserInfo(ctx *gin.Context) *models.User {
	if value, exists := ctx.Get(constants.ContextUserInfo); exists {
		return value.(*models.User)
	}
	return nil
}

func NewWithContext[T any](ctx *gin.Context, newController func(c *Context[T])) {
	newController(&Context[T]{
		Context: ctx,
		DTOData: *new(T),
		Session: Session{
			session:     sessions.Default(ctx),
			isDebugging: gin.Mode() == gin.TestMode || gin.Mode() == gin.DebugMode,
		},
		IsLogin:  ctx.GetBool(constants.IsContextLogin),
		UserInfo: getUserInfo(ctx),
	})
}

func NewWithMiddlewareContext[T any](newMiddleware func(c *Context[T])) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		NewWithContext(ctx, newMiddleware)
	}
}
