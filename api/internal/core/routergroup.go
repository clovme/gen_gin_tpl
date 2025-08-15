package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RouterGroup 路由组
type RouterGroup struct {
	*gin.RouterGroup
}

type RoutesInfo struct {
	Method string
	Path   string
	Name   string
	Type   string
}

// HandlerFunc 路由处理函数
type HandlerFunc func(*Context)

var routesInfo []RoutesInfo

// wrapHandler 路由处理函数包装
func wrapHandler(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handlerFunc(NewContext(c))
	}
}

// handle 路由处理函数注册
func (group *RouterGroup) handle(httpMethod, relativePath string, handler HandlerFunc, typ, name string) {
	routesInfo = append(routesInfo, RoutesInfo{
		Method: httpMethod,
		Path:   relativePath,
		Name:   name,
		Type:   typ,
	})
	group.RouterGroup.Handle(httpMethod, relativePath, wrapHandler(handler))
}

// POST is a shortcut for router.Handle("POST", path, handler).
func (group *RouterGroup) POST(relativePath string, handler HandlerFunc, typ, name string) {
	group.handle(http.MethodPost, relativePath, handler, typ, name)
}

// GET is a shortcut for router.Handle("GET", path, handler).
func (group *RouterGroup) GET(relativePath string, handler HandlerFunc, typ, name string) {
	group.handle(http.MethodGet, relativePath, handler, typ, name)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handler).
func (group *RouterGroup) DELETE(relativePath string, handler HandlerFunc, typ, name string) {
	group.handle(http.MethodDelete, relativePath, handler, typ, name)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handler).
func (group *RouterGroup) PATCH(relativePath string, handler HandlerFunc, typ, name string) {
	group.handle(http.MethodPatch, relativePath, handler, typ, name)
}

// PUT is a shortcut for router.Handle("PUT", path, handler).
func (group *RouterGroup) PUT(relativePath string, handler HandlerFunc, typ, name string) {
	group.handle(http.MethodPut, relativePath, handler, typ, name)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handler).
func (group *RouterGroup) OPTIONS(relativePath string, handler HandlerFunc, typ, name string) {
	group.handle(http.MethodOptions, relativePath, handler, typ, name)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handler).
func (group *RouterGroup) HEAD(relativePath string, handler HandlerFunc, typ, name string) {
	group.handle(http.MethodHead, relativePath, handler, typ, name)
}
