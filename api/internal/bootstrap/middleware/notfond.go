package middleware

import (
	"gen_gin_tpl/pkg/enums/code"
	httpLog "gen_gin_tpl/pkg/logger/http"
	"gen_gin_tpl/pkg/resp"
	"github.com/gin-gonic/gin"
)

// RegisterNoRoute 注册404处理
func RegisterNoRoute(engine *gin.Engine) {
	engine.NoRoute(func(c *gin.Context) {
		httpLog.Error(c).Msg("请求地址错误")
		// 此处可按需要修改
		resp.JsonSafe(c, code.NotFound, code.NotFound.Desc(), nil)
		c.Abort()
	})
}
