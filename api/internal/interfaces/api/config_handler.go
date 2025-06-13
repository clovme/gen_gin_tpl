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

// Config 列表
// @Router			/config [get]
// @Group 			public
func (h *ConfigHandler) Config(c *gin.Context) {
	config, err := h.ConfigService.GetConfig()
	if err != nil {
		resp.JsonSafeCode(c, em_http.ErrInternal, "Failed to get config", nil)
		return
	}

	resp.JsonSafe(c, em_http.Success.Desc(), config)
}
