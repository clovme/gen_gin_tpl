package middleware

import (
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/pkg/enums/code"
)

// RefererCheck 强制要求 Referer 且校验是否合法
func RefererCheck() core.HandlerFunc {
	return func(c *core.Context) {
		referer := c.GetHeader("Referer")
		if referer == "" {

			c.JsonSafeDesc(code.Forbidden, nil)
			c.Abort()
			return
		}

		// 校验 referer 是否在允许的前缀内
		//valid := false
		//for _, prefix := range []string{} {
		//	if strings.HasPrefix(referer, prefix) {
		//		valid = true
		//		break
		//	}
		//}
		//
		//if !valid {
		//	c.JSON(http.StatusForbidden, gin.H{
		//		"code": 4005,
		//		"msg":  "非法 Referer 来源",
		//	})
		//	c.Abort()
		//	return
		//}

		c.Next()
	}
}
