package middleware

import (
	"gen_gin_tpl/pkg/enums/code"
	httpLog "gen_gin_tpl/pkg/logger/http"
	"gen_gin_tpl/pkg/resp"
	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware panic 捕捉中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录 panic 错误，附带 stack trace
				httpLog.Panic(c).Interface("panic", err).Msg("捕捉到请求异常")

				// 可扩展：钉钉/飞书/邮件/Prometheus 警报等
				// sendDingTalkAlert(err)

				// 返回统一 JSON 响应
				resp.JsonSafe(c, code.InternalServerError, code.InternalServerError.Desc(), nil)

				// 强制中断后续
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}
