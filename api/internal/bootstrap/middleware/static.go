package middleware

import (
	"fmt"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/enums/code"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils/file"
	"gen_gin_tpl/public"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"os"
	"strings"
)

// loadStaticFS 加载静态资源
func loadStaticFS(group *core.Engine, name string, staticDir any) {
	if staticDir == nil {
		staticFS, err := fs.Sub(public.StaticFS, fmt.Sprintf("web/static/%s", name))
		if err != nil {
			panic(fmt.Errorf("failed to load static dir %s: %w", name, err))
		}
		group.Engine.StaticFS(fmt.Sprintf("/%s", name), http.FS(staticFS))
	} else {
		group.Engine.Static(fmt.Sprintf("/%s", name), staticDir.(string))
	}
}

func favicon(c *gin.Context) {
	c.Data(200, "image/x-icon", public.Favicon)
	c.Set(constants.HttpLogKey, "favicon")
	c.Abort()
}

// FaviconMiddleware 加载 favicon.ico
func FaviconMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/favicon.ico" {
			// 加载自定义 favicon.ico
			staticPath, err := file.GetFileAbsPath(cfg.C.Other.DataPath, "static", "favicon.ico")
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
			c.Set(constants.HttpLogKey, "favicon")
			c.Abort()
			return
		}
		c.Next()
	}
}

// ResourceDirInterception 静态路由拦截
func ResourceDirInterception(engine *core.Engine, staticDir string) {
	// 静态资源根目录名称
	rootDirMap := map[string]any{"static": staticDir}

	// 读取嵌入二进制的静态资源目录
	dirEntries, err := public.StaticFS.ReadDir("web/static")
	if err != nil {
		log.Error().Err(err).Msg("读取嵌入二进制的静态资源目录失败")
		return
	}

	for _, dir := range dirEntries {
		if dir.IsDir() {
			rootDirMap[dir.Name()] = nil
		}
	}

	// 静态资源路由拦截中间件处理，中间件部分
	engine.Use(func(c *core.Context) {
		if strings.HasSuffix(c.FullPath(), "*filepath") {
			c.Set(constants.HttpLogKey, "静态资源")
		}
		for name, _ := range rootDirMap {
			if strings.EqualFold(c.Request.URL.Path, fmt.Sprintf("/%s/", name)) {
				responseJsonOrHtml(c, code.RequestNotFound, http.StatusNotFound)
				return
			}
		}
	})

	// 加载静态资源
	for name, path := range rootDirMap {
		loadStaticFS(engine, name, path)
	}
}
