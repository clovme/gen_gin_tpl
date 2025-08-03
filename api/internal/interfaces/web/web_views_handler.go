package web

import (
	viewsService "gen_gin_tpl/internal/application/views"
	"gen_gin_tpl/internal/schema/dto"
	"gen_gin_tpl/pkg/page"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ViewsHandler struct {
	Service *viewsService.WebViewsService
}

// ViewsIndexHandler
// @Router			/ [GET]
// @Group 			noAuthView
func (h *ViewsHandler) ViewsIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", page.ViewDataNil("首页"))
}

// ViewsLoginHandler
// @Router			/login.html [GET]
// @Group 			noAuthView
func (h *ViewsHandler) ViewsLoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", page.ViewData[dto.RegeditDTO]("用户登录", dto.RegeditDTO{}, nil))
}

// ViewsRegeditHandler
// @Router			/regedit.html [GET]
// @Group 			noAuthView
func (h *ViewsHandler) ViewsRegeditHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "regedit.html", page.ViewData[dto.RegeditDTO]("用户注册", dto.RegeditDTO{}, nil))
}
