package routers

import (
	"fmt"
	"gen_gin_tpl/internal/bootstrap/middleware"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/public"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"gorm.io/gorm"
	"html/template"
	"io/fs"
	"net/http"
	"strings"
	"time"
)

func sessionStore() cookie.Store {
	store := cookie.NewStore(public.PrivatePEM)
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   0, // 3600 * 24 * 7, // 有效期 7 天
		HttpOnly: true,
		Secure:   false, // 本地调试 http 必须 false
	})
	return store
}

// regeditStaticFS 在注册中间件之前调用
func regeditStaticFS(engine *core.Engine, staticDir string) {
	loadStaticFS := func(group *core.Engine, name string, dir string) {
		if dir == "" {
			staticFS, err := fs.Sub(public.StaticFS, fmt.Sprintf("web/static/%s", name))
			if err != nil {
				panic(fmt.Errorf("failed to load static dir %s: %w", name, err))
			}
			group.StaticFS(fmt.Sprintf("/%s", name), http.FS(staticFS))
		} else {
			group.Static(fmt.Sprintf("/%s", name), dir)
		}
	}

	loadStaticFS(engine, "admin", "")
	loadStaticFS(engine, "assets", "")
	loadStaticFS(engine, "static", staticDir)
}

// regeditMiddleware 注册中间件
func regeditMiddleware(engine *core.Engine, staticDir string) {
	engine.Engine.Use(sessions.Sessions("session", sessionStore()))
	engine.Use(middleware.InitializationMiddleware())            // 初始化中间件
	engine.Engine.Use(middleware.LogMiddleware(2 * time.Second)) // 请求日志，记录全流程
	engine.Engine.Use(middleware.FaviconMiddleware())            // /favicon.ico

	regeditStaticFS(engine, staticDir)

	engine.Use(middleware.RecoveryMiddleware())     // 抓捕 panic，防止服务崩溃
	engine.Engine.Use(middleware.CorsMiddleware())  // 跨域，处理请求头
	engine.Engine.Use(middleware.EncryptResponse()) // 响应加密，最后处理输出
}

// regeditTemplate 注册模板
func regeditTemplate(engine *core.Engine) {
	loadTemplates := func(funcMap template.FuncMap) *template.Template {
		files, _ := fs.Glob(public.TemplatesFS, "web/templates/**/*.html")
		tpl := template.New("templates").Funcs(funcMap)
		for _, file := range files {
			// 去掉 "templates/" 前缀，让模板名变得更简洁
			name := strings.TrimPrefix(file, "web/templates/")
			content, _ := public.TemplatesFS.ReadFile(file)
			tpl = template.Must(tpl.New(name).Parse(string(content)))
		}
		return tpl
	}

	engine.SetHTMLTemplate(loadTemplates(template.FuncMap{
		"timeStamp":  timeStamp,
		"formatDate": formatDate,
	}))
}

// regeditRoutes 注册路由
func regeditRoutes(engine *core.Engine, db *gorm.DB) {
	v1 := engine.Group("/api/v1")
	routers := routeGroup{
		public:     v1,
		publicView: engine.Group(""),
		adminView:  engine.Group(""),
		//noAuth:     v1.Group("/", middleware.NoAuthMiddleware(), middleware.GrayTimeCheck(), middleware.RefererCheck()),
		noAuthView: engine.Group("", middleware.NoAuthMiddleware()),
	}

	// 注册路由
	routers.register(db)

	// 注册404处理
	middleware.RegisterNoRoute(engine)
}

// Initialization 初始化 web 服务
// 参数：
//   - db: 数据库连接对象
//   - staticDir: 静态文件目录
//
// 返回值：
//   - *gin.Engine: 初始化后的 Gin 引擎
func Initialization(db *gorm.DB, staticDir string) *core.Engine {
	engine := core.New()

	regeditTemplate(engine)
	regeditMiddleware(engine, staticDir)
	regeditRoutes(engine, db)

	return engine
}
