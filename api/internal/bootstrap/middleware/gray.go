package middleware

import (
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/pkg/enums/code"
	"strconv"
	"time"
)

// abs 取绝对值
func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// GrayTimeCheck 灰度时间戳检测中间件
func GrayTimeCheck() core.HandlerFunc {
	return func(c *core.Context) {
		// 获取 _t 参数
		tStr := c.Query("_t")
		if tStr == "" {
			c.JsonSafe(code.BadRequest, "缺少_t参数", nil)
			c.Abort()
			return
		}

		// 转换成 int64
		timestamp, err := strconv.ParseInt(tStr, 10, 64)
		if err != nil {
			c.JsonSafe(code.BadRequest, "非法_t参数", nil)
			c.Abort()
			return
		}

		// 当前时间
		now := time.Now().Unix()

		// 判断时间差是否超出 10 分钟（600秒）
		if abs(now-timestamp) > 600 {
			c.JsonSafe(code.Forbidden, "请求已过期或时间异常", nil)
			c.Abort()
			return
		}

		// 放行
		c.Next()
	}
}
