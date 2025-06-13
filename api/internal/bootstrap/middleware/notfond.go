package middleware

import (
	"gen_gin_tpl/pkg/enums/em_http"
	"gen_gin_tpl/pkg/logger/http_log"
	"gen_gin_tpl/pkg/resp"
	"github.com/gin-gonic/gin"
)

// RegisterNoRoute 注册404处理
func RegisterNoRoute(engine *gin.Engine) {
	engine.NoRoute(func(c *gin.Context) {
		http_log.Error(c).Msg("请求地址错误")
		// 此处可按需要修改
		resp.JsonSafeCode(c, em_http.ErrNotFound, "请输入正确的请求地址", nil)
		c.Abort()
	})
}
