package middleware

import (
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/pkg/enums/code"
	httpLog "gen_gin_tpl/pkg/logger/http"
)

// RegisterNoRoute 注册404处理
func RegisterNoRoute(engine *core.Engine) {
	engine.NoRoute(func(c *core.Context) {
		httpLog.Error(c.Context).Msg("请求地址错误")
		// 此处可按需要修改
		c.JsonSafe(code.NotFound, code.NotFound.Desc(), nil)
		c.Abort()
	})
}
