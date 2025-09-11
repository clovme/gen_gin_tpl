package middleware

import (
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/pkg/enums/code"
	httpLog "gen_gin_tpl/pkg/logger/http"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func viewError(errCode code.ResponseCode) gin.H {
	return gin.H{
		"Code": errCode.Int(),
		"Desc": errCode.Desc(),
	}
}

func responseJsonOrHtml(c *core.Context, errCode code.ResponseCode, httpCode int) {
	if c.IsAjax {
		c.JsonSafeDesc(errCode, nil)
		c.AbortWithStatus(httpCode)
		return
	}
	if code.RequestForbidden.Is(errCode) {
		c.Redirect(http.StatusFound, c.Router.Name("indexView").Path) // 302 跳转更常用
		c.Abort()                                                     // 中断后续中间件和 handler 执行！
		return
	}
	c.HTML("views/error.html", strconv.Itoa(httpCode), viewError(errCode))
	c.AbortWithStatus(httpCode)
}

// RegisterNoRoute 注册404处理
func RegisterNoRoute(engine *core.Engine) {
	engine.NoRoute(func(c *core.Context) {
		httpLog.Error(c.Context).Msg("请求地址错误")
		responseJsonOrHtml(c, code.RequestNotFound, http.StatusNotFound)
	})
}
