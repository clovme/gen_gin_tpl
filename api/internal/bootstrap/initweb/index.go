package initweb

import (
	"errors"
	"fmt"
	"gen_gin_tpl/internal/bootstrap/middleware"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/utils/file"
	"gen_gin_tpl/pkg/utils/network"
	"gen_gin_tpl/public"
	"html/template"
	"io/fs"
	"net/http"
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

	engine.GET("/", viewHandler, "init", "index", "初始化首页")
	engine.GET("/logs", LogWebSocketHandler, "init", "wsLog", "WebSocket日志打印")
	engine.GET("/copyright", copyrightHandler, "init", "copyright", "版权信息")
	engine.GET("/initialize", formHandler, "init", "initialize", "获取初始化配置")
	engine.POST("/initialize", postHandler, "init", "postInitialize", "提交初始化配置")

	engine.NoRoute(func(c *core.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	port := network.GetPort(cfg.C.Web.Port + 1)
	server = &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: engine.Engine,
	}

	fmt.Println("初始化服务启动完成，请进行初始化配置")
	fmt.Printf("\thttp://%s\n", server.Addr)
	fmt.Printf("\thttp://localhost:%d\n", port)
	fmt.Printf("程序所在路径: %s\n", exePath)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("初始化服务异常退出: %v\n", err)
	}
}
