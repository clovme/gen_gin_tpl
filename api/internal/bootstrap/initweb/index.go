package initweb

import (
	"context"
	"errors"
	"fmt"
	"gen_gin_tpl/internal/bootstrap/middleware"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/logger"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils/file"
	"gen_gin_tpl/pkg/utils/network"
	"gen_gin_tpl/public"
	"html/template"
	"io/fs"
	"net/http"
	"time"
)

var (
	server *http.Server
)

// StartInitializeWeb 启动初始化服务
func StartInitializeWeb() {
	exePath, err := file.GetFileAbsPath(".")

	if err != nil {
		fmt.Printf("获取程序所在路径失败: %v\n", err)
		return
	}
	engine := core.New()

	engine.Engine.Use(middleware.CorsMiddleware())
	engine.Engine.Use(middleware.FaviconMiddleware()) // /favicon.ico

	// 加载嵌入模板
	tmpl := template.Must(template.New("template").ParseFS(public.InitWebFS, "initweb/*.html"))
	engine.Engine.SetHTMLTemplate(tmpl)

	staticFS, _ := fs.Sub(public.InitWebFS, "initweb/assets")
	engine.Engine.StaticFS("/assets", http.FS(staticFS))

	engine.GET("/", viewHandler, "web", "index", "初始化首页")
	engine.GET("/logs", LogWebSocketHandler, "log", "wsLog", "WebSocket日志打印")
	engine.GET("/copyright", copyrightHandler, "api", "copyright", "版权信息")
	engine.GET("/initialize", formHandler, "api", "initialize", "获取初始化配置")
	engine.POST("/initialize", postHandler, "api", "postInitialize", "提交初始化配置")

	engine.NoRoute(func(c *core.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", network.GetLanIP(), network.GetPort(cfg.CWeb.Port+1)),
		Handler: engine.Engine,
	}

	fmt.Printf("初始化服务启动完成，访问 http://%s 进行初始化配置\n程序所在路径: %s\n", server.Addr, exePath)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("初始化服务异常退出: %v\n", err)
	}
}

// StopInitializeWeb 关闭初始化服务
func StopInitializeWeb() {
	time.Sleep(5 * time.Second) // 防空转
	for {
		if network.IsMainWebStart(cfg.CWeb.Host, cfg.CWeb.Port) {
			if server != nil {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				if err := server.Shutdown(ctx); err != nil {
					_ = server.Close()
				}
				cancel()
				log.Info().Msgf("初始化服务已关闭，访问 https://%s:%d 首页", network.GetLocalIP(cfg.CWeb.Host), cfg.CWeb.Port)
			}
			break
		}
		time.Sleep(time.Second)
	}
	time.Sleep(100 * time.Millisecond)
	logger.LogHub.SendMessage(fmt.Sprintf("000000https://%s:%d", cfg.CWeb.Host, cfg.CWeb.Port))
	logger.LogHub.CloseAllClient()
}
