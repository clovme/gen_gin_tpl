package middleware

import (
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/utils/file"
	"gen_gin_tpl/public"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
)

func favicon(c *gin.Context) {
	data, err := fs.ReadFile(public.Favicon, "favicon.ico")
	if err != nil {
		c.Status(404)
		return
	}
	c.Data(200, "image/x-icon", data)
	c.Abort()
}

// FaviconMiddleware 加载 favicon.ico
func FaviconMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/favicon.ico" {
			// 加载自定义 favicon.ico
			staticPath, err := file.GetFileAbsPath(cfg.COther.DataPath, "static", "favicon.ico")
			if err != nil || !file.IsFileExist(staticPath) {
				// 加载默认 favicon.ico
				favicon(c)
				return
			}

			// 加载自定义 favicon.ico
			data, err := os.ReadFile(staticPath)
			if err != nil {
				// 加载默认 favicon.ico
				favicon(c)
				return
			}
			c.Data(200, "image/x-icon", data)
			c.Abort()
			return
		}
		c.Next()
	}
}
