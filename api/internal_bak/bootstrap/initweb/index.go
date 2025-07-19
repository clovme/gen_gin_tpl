package initweb

import (
	"context"
	"errors"
	"fmt"
	"gen_gin_tpl/internal/bootstrap/middleware"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/logger"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils/u_file"
	"gen_gin_tpl/pkg/utils/u_network"
	"gen_gin_tpl/public"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"math/rand"
	"net/http"
	"time"
)

var (
	server *http.Server
)

// getRandomPort 获取随机端口
func getRandomPort() int {
	rand.New(rand.NewSource(time.Now().UnixNano())) // 不设置种子会每次一样
	return rand.Intn(65535-50000+1) + 50000
}

// StartInitializeWeb 启动初始化服务
func StartInitializeWeb() {
	exePath, err := u_file.GetFileAbsPath(".", "")

	if err != nil {
		fmt.Printf("获取程序所在路径失败: %v\n", err)
		return
	}
	engine := gin.New()

	engine.Use(middleware.CorsMiddleware())

	// 加载嵌入模板
	tmpl := template.Must(template.New("template").Delims("[{", "}]").ParseFS(public.InitWebFS, "initweb/*.html"))
	engine.SetHTMLTemplate(tmpl)

	engine.StaticFileFS("/favicon.ico", "favicon.ico", http.FS(public.Favicon))

	staticFS, _ := fs.Sub(public.InitWebFS, "initweb/assets")
	engine.StaticFS("/assets", http.FS(staticFS))

	engine.GET("/", viewHandler)
	engine.GET("/logs", LogWebSocketHandler)
	engine.GET("/copyright", copyrightHandler)
	engine.GET("/initialize", formHandler)
	engine.POST("/initialize", postHandler)

	engine.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	port := u_network.GetPort(getRandomPort())
	server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: engine,
	}

	fmt.Printf("初始化服务启动完成，访问 http://%s:%d 进行初始化配置, 程序所在路径: %s\n", u_network.GetLanIP(), port, exePath)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("初始化服务异常退出: %v\n", err)
	}
}

// StopInitializeWeb 关闭初始化服务
func StopInitializeWeb() {
	time.Sleep(5 * time.Second) // 防空转
	for {
		if u_network.IsMainWebStart(cfg.CWeb.Host, cfg.CWeb.Port) {
			if server != nil {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				if err := server.Shutdown(ctx); err != nil {
					server.Close()
				}
				cancel()
				log.Info().Msgf("初始化服务已关闭，访问 http://%s:%d 首页", u_network.GetLocalIP(cfg.CWeb.Host), cfg.CWeb.Port)
			}
			break
		}
		time.Sleep(time.Second)
	}
	time.Sleep(100 * time.Millisecond)
	logger.LogHub.SendMessage(fmt.Sprintf("000000http://%s:%d", u_network.GetLocalIP(cfg.CWeb.Host), cfg.CWeb.Port))
	logger.LogHub.CloseAllClient()
}
