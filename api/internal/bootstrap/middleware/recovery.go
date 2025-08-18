package middleware

import (
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/pkg/enums/code"
	httpLog "gen_gin_tpl/pkg/logger/http"
	"net/http"
)

// RecoveryMiddleware panic 捕捉中间件
func RecoveryMiddleware() core.HandlerFunc {
	return func(c *core.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录 panic 错误，附带 stack trace
				httpLog.Panic(c.Context).Interface("panic", err).Msg("捕捉到请求异常")

				// 可扩展：钉钉/飞书/邮件/Prometheus 警报等
				// sendDingTalkAlert(err)

				// 返回统一处理
				responseJsonOrHtml(c, code.ServerInternalError, http.StatusInternalServerError)

				// 强制中断后续
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}
