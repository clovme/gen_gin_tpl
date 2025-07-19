package api

import (
	"gen_gin_tpl/internal/application"
	"gen_gin_tpl/pkg/enums/em_http"
	"gen_gin_tpl/pkg/resp"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService *application.LoginService
}

// GetLogin
// @Summary			HTML首页
// @Description 	HTML首页
// @Tags        	HTML首页
// @Accept       	text/html
// @Produce      	text/html
// @Success      	200  text/html  text/html
// @Router			/login [get]
// @Group 			noAuth
func (h *LoginHandler) GetLogin(c *gin.Context) {
	resp.JsonSafe(c, em_http.Success.Desc(), nil)
}
