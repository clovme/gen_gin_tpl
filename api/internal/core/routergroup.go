package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// RouterGroup 路由组
type RouterGroup struct {
	*gin.RouterGroup
	uriPrefix string
}

type RoutesInfo struct {
	Method      string
	Path        string
	Name        string
	Type        string
	Description string
}

// HandlerFunc 路由处理函数
type HandlerFunc func(*Context)

var routesInfo = map[string]RoutesInfo{}

// wrapHandler 路由处理函数包装
func wrapHandler(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handlerFunc(NewContext(c))
	}
}

func groupFunc(group RouterGroup, relativePath string, handlers ...HandlerFunc) *RouterGroup {
	handlerList := make([]gin.HandlerFunc, 0)
	for _, h := range handlers {
		handlerList = append(handlerList, wrapHandler(h))
	}
	uriPrefix := group.uriPrefix + relativePath
	if strings.HasSuffix(relativePath, "/") {
		uriPrefix = strings.TrimSuffix(uriPrefix, "/")
	}
	return &RouterGroup{
		RouterGroup: group.RouterGroup.Group(relativePath, handlerList...),
		uriPrefix:   uriPrefix,
	}
}

// handle 路由处理函数注册
func (group *RouterGroup) handle(httpMethod, relativePath string, handler HandlerFunc, typ, name, description string) {
	key := fmt.Sprintf("%s:%s%s", httpMethod, group.uriPrefix, relativePath)
	routesInfo[key] = RoutesInfo{
		Method:      httpMethod,
		Path:        group.uriPrefix + relativePath,
		Name:        name,
		Type:        typ,
		Description: description,
	}
	group.RouterGroup.Handle(httpMethod, relativePath, wrapHandler(handler))
}

// POST is a shortcut for router.Handle("POST", path, handler).
func (group *RouterGroup) POST(relativePath string, handler HandlerFunc, typ, name, description string) {
	group.handle(http.MethodPost, relativePath, handler, typ, name, description)
}

// GET is a shortcut for router.Handle("GET", path, handler).
func (group *RouterGroup) GET(relativePath string, handler HandlerFunc, typ, name, description string) {
	group.handle(http.MethodGet, relativePath, handler, typ, name, description)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handler).
func (group *RouterGroup) DELETE(relativePath string, handler HandlerFunc, typ, name, description string) {
	group.handle(http.MethodDelete, relativePath, handler, typ, name, description)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handler).
func (group *RouterGroup) PATCH(relativePath string, handler HandlerFunc, typ, name, description string) {
	group.handle(http.MethodPatch, relativePath, handler, typ, name, description)
}

// PUT is a shortcut for router.Handle("PUT", path, handler).
func (group *RouterGroup) PUT(relativePath string, handler HandlerFunc, typ, name, description string) {
	group.handle(http.MethodPut, relativePath, handler, typ, name, description)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handler).
func (group *RouterGroup) OPTIONS(relativePath string, handler HandlerFunc, typ, name, description string) {
	group.handle(http.MethodOptions, relativePath, handler, typ, name, description)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handler).
func (group *RouterGroup) HEAD(relativePath string, handler HandlerFunc, typ, name, description string) {
	group.handle(http.MethodHead, relativePath, handler, typ, name, description)
}

// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
// For example, all the routes that use a common middleware for authorization could be grouped.
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return groupFunc(*group, relativePath, handlers...)
}
