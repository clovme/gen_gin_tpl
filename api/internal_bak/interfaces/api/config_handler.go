package api

import (
	"gen_gin_tpl/internal/application"
	"gen_gin_tpl/pkg/enums/em_http"
	"gen_gin_tpl/pkg/resp"
	"github.com/gin-gonic/gin"
)

type ConfigHandler struct {
	ConfigService *application.ConfigService
}

// PostConfig 列表
// @Router			/public/config [post]
// @Group 			public
func (h *ConfigHandler) PostConfig(c *gin.Context) {
	config, err := h.ConfigService.GetConfig()
	if err != nil {
		resp.JsonSafeCode(c, em_http.InternalServerError, "Failed to get config", nil)
		return
	}

	resp.JsonSafe(c, em_http.Success.Desc(), config)
}
