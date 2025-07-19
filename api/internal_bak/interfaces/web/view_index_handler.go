package web

import (
	"gen_gin_tpl/internal/application"
	"gen_gin_tpl/pkg/page"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ViewIndexHandler struct {
	ViewIndexService *application.LoginService
}

// GetIndexView
// @Router			/ [get]
// @Group 			noAuthView
func (h *ViewIndexHandler) GetIndexView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", page.ViewDataNil("首页"))
}
