package web

import (
	"fmt"
	viewsService "gen_gin_tpl/internal/application/views"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/internal/schema/dto"
	"gen_gin_tpl/pkg/captcha"
	"gen_gin_tpl/pkg/enums/code"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils"
	"net/http"
	"strconv"
)

type ViewsHandler struct {
	Service *viewsService.WebViewsService
}

// GetImagesCaptcha 生成验证码
// @Router			/public/captcha.png [GET]
// @Group 			publicView
// @Summary			生成验证码
// @Type			web
func (h *ViewsHandler) GetImagesCaptcha(c *core.Context) {
	// 生成验证码
	imageBytes, err := captcha.NewGenerate(c.Session)
	if err != nil {
		log.Error().Err(err).Msg("验证码生成失败")
		c.JsonSafe(code.InternalServerError, "验证码生成失败", nil)
		return
	}
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate")
	c.Data(http.StatusOK, "image/png", imageBytes)
}

// GetViewsIndexHandler
// @Router			/ [GET]
// @Group 			publicView
// @Summary			首页
// @Type			web
func (h *ViewsHandler) GetViewsIndexHandler(c *core.Context) {
	c.HTML("views/index.html", "首页", nil)
}

// GetViewsAdminHandler 后台首页
// @Router			/admin.html [GET]
// @Group 			adminView
// @Summary			后台首页
// @Type			web
func (h *ViewsHandler) GetViewsAdminHandler(c *core.Context) {
	c.HTML("admin/index.html", "后台首页", nil)
}

// GetViewsLoginHandler 登录页
// @Router			/login.html [GET]
// @Group 			noAuthView
// @Summary			登录页面
// @Type			web
func (h *ViewsHandler) GetViewsLoginHandler(c *core.Context) {
	var loginDTO dto.LoginDTO
	loginDTO.Username = "qingyuheji@qq.com"
	loginDTO.Password = "silvery.0"
	c.HTML("views/login.html", "用户登录", loginDTO)
}

// PostViewsLoginHandler 登录出处理
// @Router			/login.html [POST]
// @Group 			noAuthView
// @Summary			登录处理接口
// @Type			api
func (h *ViewsHandler) PostViewsLoginHandler(c *core.Context) {
	var loginDTO dto.LoginDTO
	loginDTO.CaptchaID = c.Session.GetImageCaptchaID()
	if err := c.ShouldBind(&loginDTO); err != nil {
		log.Error().Err(err).Msg(code.VerifyError.Desc())
		c.JsonSafeDesc(code.VerifyError, loginDTO.TranslateError(err))
		return
	}
	success, msg := h.Service.UserLogin(loginDTO, c.Session)
	if !success {
		c.JsonSafe(code.VerifyError, msg, nil)
		return
	}
	c.JsonSafe(code.Success, msg, nil)
}

// GetViewsRegeditHandler 注册页
// @Router			/regedit.html [GET]
// @Group 			noAuthView
// @Summary			注册页面
// @Type			web
func (h *ViewsHandler) GetViewsRegeditHandler(c *core.Context) {
	var regeditDTO dto.RegeditDTO
	regeditDTO.Email = "qingyuheji@qq.com"
	regeditDTO.Password = "silvery.0"
	regeditDTO.ConfirmPassword = "silvery.0"
	c.HTML("views/regedit.html", "用户注册", regeditDTO)
}

// PostViewsRegeditHandler 注册处理接口
// @Router			/regedit.html [POST]
// @Group 			noAuthView
// @Summary			注册处理
// @Type			api
func (h *ViewsHandler) PostViewsRegeditHandler(c *core.Context) {
	username := strconv.FormatInt(utils.GenerateID(), 10)
	var regeditDTO dto.RegeditDTO
	regeditDTO.Phone = username[:11]
	regeditDTO.CaptchaID = c.Session.GetImageCaptchaID()
	regeditDTO.EmailID = c.Session.GetEmailCaptchaID()
	regeditDTO.Username = username
	regeditDTO.Email = fmt.Sprintf("%s@qq.com", username)

	if err := c.ShouldBind(&regeditDTO); err != nil {
		log.Error().Err(err).Msg(code.VerifyError.Desc())
		c.JsonSafeDesc(code.VerifyError, regeditDTO.TranslateError(err))
		return
	}
	exist, msg := h.Service.CreateUser(regeditDTO, c.Session)
	if !exist {
		c.JsonUnSafe(code.CreateError, msg, nil)
		return
	}

	c.JsonUnSafe(code.Success, "用户注册成功！", map[string]string{
		"token": msg,
	})
}
