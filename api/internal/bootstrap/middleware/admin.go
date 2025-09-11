package middleware

import (
	"gen_gin_tpl/internal/core"
)

func AdminMiddleware() core.HandlerFunc {
	return func(c *core.Context) {
		c.Next()
	}
}
