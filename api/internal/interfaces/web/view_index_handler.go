package web

import (
	"gen_gin_tpl/internal/application"
	"gen_gin_tpl/pkg/resp"
	"github.com/gin-gonic/gin"
)

type ViewIndexHandler struct {
	ViewIndexService *application.ViewIndexService
}

// IndexView
// @Router			/ [get]
// @Group 			noAuthView
func (h *ViewIndexHandler) IndexView(c *gin.Context) {
	resp.HtmlUnSafe(c, "index.html", nil)
}
