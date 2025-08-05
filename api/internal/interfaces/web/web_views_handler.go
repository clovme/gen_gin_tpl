package web

import (
	"fmt"
	viewsService "gen_gin_tpl/internal/application/views"
	"gen_gin_tpl/internal/schema/dto"
	"gen_gin_tpl/pkg/captcha"
	"gen_gin_tpl/pkg/enums/code"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/page"
	"gen_gin_tpl/pkg/resp"
	"gen_gin_tpl/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ViewsHandler struct {
	Service *viewsService.WebViewsService
}

// GetImagesCaptcha 生成验证码
// @Router			/public/captcha.png [GET]
// @Group 			publicView
func (h *ViewsHandler) GetImagesCaptcha(c *gin.Context) {
	// 生成验证码
	imageBytes, err := captcha.NewGenerate(c)
	if err != nil {
		log.Error().Err(err).Msg("验证码生成失败")
		resp.JsonSafe(c, code.InternalServerError, "验证码生成失败", nil)
		return
	}

	c.Header("Cache-Control", "no-store, no-cache, must-revalidate")
	c.Data(http.StatusOK, "image/png", imageBytes)
}

// ViewsIndexHandler
// @Router			/ [GET]
// @Group 			publicView
func (h *ViewsHandler) ViewsIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", page.ViewDataNil(c, "首页"))
}

// ViewsLoginHandler 登录页
// @Router			/login.html [GET]
// @Group 			noAuthView
func (h *ViewsHandler) ViewsLoginHandler(c *gin.Context) {
	var loginDTO dto.LoginDTO
	c.HTML(http.StatusOK, "login.html", page.ViewData[dto.LoginDTO](c, "用户登录", loginDTO, nil))
}

// PostViewsLoginHandler 登录出处理
// @Router			/login.html [POST]
// @Group 			noAuthView
func (h *ViewsHandler) PostViewsLoginHandler(c *gin.Context) {
	loginDTO := dto.LoginDTO{Context: c}
	if err := c.ShouldBind(&loginDTO); err != nil {
		log.Error().Err(err).Msg(code.VerifyError.Desc())
		resp.JsonSafeDesc(c, code.VerifyError, loginDTO.TranslateError(err))
		return
	}
	success, msg := h.Service.UserLogin(c, loginDTO)
	if !success {
		resp.JsonSafe(c, code.VerifyError, msg, nil)
		return
	}
	resp.JsonSafe(c, code.Success, msg, nil)
}

// ViewsRegeditHandler 注册页
// @Router			/regedit.html [GET]
// @Group 			noAuthView
func (h *ViewsHandler) ViewsRegeditHandler(c *gin.Context) {
	var regeditDTO dto.RegeditDTO
	c.HTML(http.StatusOK, "regedit.html", page.ViewData[dto.RegeditDTO](c, "用户注册", regeditDTO, nil))
}

// PostViewsRegeditHandler 注册处理
// @Router			/regedit.html [POST]
// @Group 			noAuthView
func (h *ViewsHandler) PostViewsRegeditHandler(c *gin.Context) {
	username := strconv.FormatInt(utils.GenerateID(), 10)
	regeditDTO := dto.RegeditDTO{Context: c, Username: username, Email: fmt.Sprintf("%s@qq.com", username)}

	if err := c.ShouldBind(&regeditDTO); err != nil {
		log.Error().Err(err).Msg(code.VerifyError.Desc())
		resp.JsonSafeDesc(c, code.VerifyError, regeditDTO.TranslateError(err))
		return
	}

	exist, msg := h.Service.CreateUser(c, regeditDTO)
	if !exist {
		resp.JsonUnSafe(c, code.CreateError, msg, nil)
		return
	}
	resp.JsonUnSafe(c, code.Success, msg, nil)
}
