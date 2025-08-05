package routers

import (
	"gen_gin_tpl/internal/bootstrap/middleware"
	"gen_gin_tpl/public"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"html/template"
	"io/fs"
	"net/http"
	"time"
)

// Initialization 初始化 web 服务
// 参数：
//   - db: 数据库连接对象
//
// 返回值：
//   - *gin.Engine: 初始化后的 Gin 引擎
func Initialization(db *gorm.DB, staticDir string) *gin.Engine {
	engine := gin.New()

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 24 * 7, // 有效期 7 天
		HttpOnly: true,
		Secure:   false, // 本地调试 http 必须 false
	})

	// 加载中间件
	engine.Use(
		sessions.Sessions("session", store),
		middleware.SessionMiddleware(),          // 设置 session
		middleware.LogMiddleware(2*time.Second), // 请求日志，记录全流程
		middleware.RecoveryMiddleware(),         // 抓捕 panic，防止服务崩溃
		middleware.CorsMiddleware(),             // 跨域，处理请求头
		middleware.EncryptResponse(),            // 响应加密，最后处理输出
		middleware.FaviconMiddleware(),          // /favicon.ico
	)

	// 加载嵌入的模板文件
	engine.SetHTMLTemplate(func(tmpl *template.Template) *template.Template {
		// 加载自定义模板函数
		tmpl.Funcs(template.FuncMap{
			"timeStamp":  timeStamp,
			"formatDate": formatDate,
		})
		// 解析嵌入的模板文件
		return template.Must(tmpl.ParseFS(public.WebFS, "web/*.html"))
	}(template.New("template")))

	// 提供嵌入的静态文件
	staticFS, _ := fs.Sub(public.WebFS, "web/assets")
	engine.StaticFS("/assets", http.FS(staticFS))
	engine.Static("/static", staticDir)

	// 注册路由组
	v1 := engine.Group("/api/v1")
	routers := routeGroup{
		public:     v1,
		publicView: engine.Group(""),
		//noAuth:     v1.Group("/", middleware.NoAuthMiddleware(), middleware.GrayTimeCheck(), middleware.RefererCheck()),
		noAuthView: engine.Group("", middleware.NoAuthMiddleware()),
	}

	// 注册路由
	routers.register(db)

	// 注册404处理
	middleware.RegisterNoRoute(engine)

	return engine
}
