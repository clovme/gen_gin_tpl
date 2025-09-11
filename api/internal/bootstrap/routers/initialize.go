package routers

import (
	"gen_gin_tpl/internal/bootstrap/middleware"
	"gen_gin_tpl/internal/bootstrap/middleware/minit"
	"gen_gin_tpl/internal/core"
	"gorm.io/gorm"
	"html/template"
	"time"
)

// regeditMiddleware 注册中间件
func regeditMiddleware(engine *core.Engine, staticDir string) {
	engine.Engine.Use(middleware.LogMiddleware(2 * time.Second)) // 请求日志，记录全流程
	engine.Engine.Use(middleware.FaviconMiddleware())            // favicon.ico
	middleware.ResourceDirInterception(engine, staticDir)        // 资源目录拦截并加载静态资源文件，放在日志中间件后面，其他中间件前面，保证能打印日志，又不走其他中间件
	engine.Use(minit.InitializationMiddleware())                 // 初始化中间件
	engine.Use(middleware.RecoveryMiddleware())                  // 抓捕 panic，防止服务崩溃
	engine.Engine.Use(middleware.CorsMiddleware())               // 跨域，处理请求头
	//engine.Engine.Use(middleware.EncryptResponse())              // 响应加密，最后处理输出
}

// regeditTemplate 注册模板
func regeditTemplate(engine *core.Engine) {
	engine.SetHTMLTemplate("web/templates", template.FuncMap{
		"timeStamp":  timeStamp,
		"formatDate": formatDate,
	})
}

// regeditRoutes 注册路由
func regeditRoutes(engine *core.Engine, db *gorm.DB) {
	v1 := engine.Group("/api/v1")
	routers := routeGroup{
		// 接口层
		public: v1,
		auth:   v1.Group("/", middleware.AuthMiddleware()),
		//noAuth: v1.Group("/", middleware.NoAuthMiddleware()),

		// 视图层
		publicView: engine.Group("/"),
		adminView:  engine.Group("/", middleware.AdminMiddleware(), middleware.PermissionMiddleware()),
		authView:   engine.Group("/", middleware.AuthMiddleware()),
		noAuthView: engine.Group("/", middleware.NoAuthMiddleware()),
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
