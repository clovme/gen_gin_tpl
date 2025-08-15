package core

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// Engine 自定义gin.Engine
type Engine struct {
	*gin.Engine
}

// Use 注册中间件
//
// 参数:
//   - middleware: 中间件函数列表
//
// 说明:
//   - 注册中间件函数，用于在请求处理过程中进行预处理或后处理。
func (engine *Engine) Use(middleware ...HandlerFunc) {
	middlewareList := make([]gin.HandlerFunc, 0)
	for _, m := range middleware {
		middlewareList = append(middlewareList, wrapHandler(m))
	}
	engine.Engine.Use(middlewareList...)
}

// NoRoute 注册404路由
//
// 参数:
//   - handler: 404路由处理函数
//
// 说明:
//   - 注册404路由处理函数，当请求的路由不存在时调用。
func (engine *Engine) NoRoute(handler HandlerFunc) {
	engine.Engine.NoRoute(wrapHandler(handler))
}

// Group 注册路由组
//
// 参数:
//   - relativePath: 路由组的相对路径
//   - handlers: 路由组的中间件函数列表
//
// 返回值:
//   - *RouterGroup: 路由组对象
//
// 说明:
//   - 注册路由组，用于组织和管理相关的路由。
func (engine *Engine) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	handlerList := make([]gin.HandlerFunc, 0)
	for _, h := range handlers {
		handlerList = append(handlerList, wrapHandler(h))
	}
	return &RouterGroup{
		RouterGroup: engine.Engine.Group(relativePath, handlerList...),
	}
}

func (engine *Engine) Routes() (routes []RoutesInfo) {
	rInfo := make([]RoutesInfo, 0)
	for _, route := range engine.Engine.Routes() {
		if strings.HasSuffix(route.Path, "*filepath") {
			continue
		}
		for _, info := range routesInfo {
			if route.Method == info.Method && strings.HasSuffix(route.Path, info.Path) {
				rInfo = append(rInfo, RoutesInfo{
					Method: route.Method,
					Path:   route.Path,
					Name:   info.Name,
					Type:   info.Type,
				})
			}
		}
	}
	return rInfo
}

// New 创建自定义gin.Engine
//
// 参数:
//   - opts: gin.Engine的选项函数列表
//
// 返回值:
//   - *Engine: 自定义gin.Engine对象
//
// 说明:
//   - 创建自定义gin.Engine对象，用于自定义路由和中间件。
func New(opts ...gin.OptionFunc) *Engine {
	return &Engine{
		Engine: gin.New(opts...),
	}
}
