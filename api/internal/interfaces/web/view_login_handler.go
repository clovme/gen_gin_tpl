package web

import (
	"gen_gin_tpl/internal/application"
	"gen_gin_tpl/pkg/resp"
	"github.com/gin-gonic/gin"
)

type ViewLoginHandler struct {
	ViewLoginService *application.ViewLoginService
}

// LoginView
// @Router			/login [get]
// @Group 			noAuthView
func (h *ViewLoginHandler) LoginView(c *gin.Context) {
	resp.HtmlUnSafe(c, "login.html", nil)
}

// RegeditView
// @Router			/signup [get]
// @Group 			noAuthView
func (h *ViewLoginHandler) RegeditView(c *gin.Context) {
	resp.HtmlUnSafe(c, "signup.html", nil)
}
