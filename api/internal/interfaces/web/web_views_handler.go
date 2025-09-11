package web

import (
	"fmt"
	viewsService "gen_gin_tpl/internal/application/views"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/internal/schema/dto"
	"gen_gin_tpl/pkg/captcha"
	"gen_gin_tpl/pkg/enums/code"
	"gen_gin_tpl/pkg/enums/role"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils"
	"net/http"
	"strconv"
	"strings"
)

type ViewsHandler struct {
	Service *viewsService.WebViewsService
}

// GetViewsAdminHandler 系统管理首页
// @Type			web
// @Group 			adminView
// @Router			/admin.html [GET]
// @Name			adminIndexView
// @Summary			系统管理首页
func (h *ViewsHandler) GetViewsAdminHandler(c *core.Context) {
	c.HTML("admin/index.html", "后台首页", nil)
}

// GetViewsMeHandler 个人中心首页
// @Type			web
// @Group 			authView
// @Router			/me.html [GET]
// @Name			meIndexView
// @Summary			个人中心
func (h *ViewsHandler) GetViewsMeHandler(c *core.Context) {
	c.HTML("me/index.html", "个人中心", nil)
}

// GetViewsIndexHandler
// @Type			web
// @Group 			publicView
// @Router			/ [GET]
// @Name			indexView
// @Summary			首页
func (h *ViewsHandler) GetViewsIndexHandler(c *core.Context) {
	c.HTML("views/index.html", "首页", nil)
}

// GetViewsLoginHandler 用户登录页
// @Type			web
// @Group 			noAuthView
// @Router			/login.html [GET]
// @Name			loginView
// @Summary			登录
func (h *ViewsHandler) GetViewsLoginHandler(c *core.Context) {
	var loginDTO dto.LoginDTO
	loginDTO.Username = strings.ToLower(role.System.Key())
	loginDTO.Password = loginDTO.Username
	c.HTML("views/login.html", "用户登录", loginDTO)
}

// GetViewsRegeditHandler 用户注册页
// @Type			web
// @Group 			noAuthView
// @Router			/regedit.html [GET]
// @Name			regeditView
// @Summary			注册
func (h *ViewsHandler) GetViewsRegeditHandler(c *core.Context) {
	var regeditDTO dto.RegeditDTO
	regeditDTO.Email = "qingyuheji@qq.com"
	regeditDTO.Password = "silvery.0"
	regeditDTO.ConfirmPassword = "silvery.0"
	c.HTML("views/regedit.html", "用户注册", regeditDTO)
}

// GetImagesCaptcha 生成图形验证码
// @Type			web
// @Group 			publicView
// @Router			/public/captcha.png [GET]
// @Name			captchaApi
// @Summary			生成图形验证码
func (h *ViewsHandler) GetImagesCaptcha(c *core.Context) {
	// 生成验证码
	imageBytes, err := captcha.NewGenerate(c.Session.GetImageCaptchaID())
	if err != nil {
		log.Error().Err(err).Msg("验证码生成失败")
		c.JsonSafe(code.ServerInternalError, "验证码生成失败", nil)
		return
	}
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate")
	c.Data(http.StatusOK, "image/png", imageBytes)
}

// PostViewsLoginHandler 用户登录处理接口
// @Type			api
// @Group 			noAuthView
// @Router			/login.html [POST]
// @Name			loginApi
// @Summary			用户登录处理接口
func (h *ViewsHandler) PostViewsLoginHandler(c *core.Context) {
	var loginDTO dto.LoginDTO
	loginDTO.CaptchaID = c.Session.GetImageCaptchaID()
	if err := c.ShouldBind(&loginDTO); err != nil {
		log.Error().Err(err).Msg(code.ServiceVerifyError.Desc())
		c.JsonSafeDesc(code.ServiceVerifyError, loginDTO.TranslateError(err))
		return
	}
	success, msg := h.Service.UserLogin(loginDTO, c.Session)
	if !success {
		c.JsonSafe(code.ServiceVerifyError, msg, nil)
		return
	}
	c.JsonUnSafe(code.Success, "用户登录成功！", map[string]string{
		"token": msg,
	})
}

// PostViewsRegeditHandler 注册处理接口
// @Type			api
// @Group 			noAuthView
// @Router			/regedit.html [POST]
// @Name			regeditApi
// @Summary			用户注册处理接口
func (h *ViewsHandler) PostViewsRegeditHandler(c *core.Context) {
	username := strconv.FormatInt(utils.GenerateID(), 10)
	var regeditDTO dto.RegeditDTO
	regeditDTO.Phone = username[:11]
	regeditDTO.CaptchaID = c.Session.GetImageCaptchaID()
	regeditDTO.EmailID = c.Session.GetEmailCaptchaID()
	regeditDTO.Username = username
	regeditDTO.Email = fmt.Sprintf("%s@qq.com", username)

	if err := c.ShouldBind(&regeditDTO); err != nil {
		log.Error().Err(err).Msg(code.ServiceVerifyError.Desc())
		c.JsonSafeDesc(code.ServiceVerifyError, regeditDTO.TranslateError(err))
		return
	}
	exist, msg := h.Service.CreateUser(regeditDTO, c.Session)
	if !exist {
		c.JsonUnSafe(code.ServiceCreateError, msg, nil)
		return
	}

	c.JsonUnSafe(code.Success, "用户注册成功！", map[string]string{
		"token": msg,
	})
}

// GetViewsLogoutHandler 注销系统登录接口
// @Type			api
// @Group 			auth
// @Router			/logout [POST]
// @Name			logoutApi
// @Summary			注销系统登录接口
func (h *ViewsHandler) GetViewsLogoutHandler(c *core.Context) {
	ok, msg := h.Service.UserLogout(c)
	if !ok {
		c.JsonUnSafe(code.Fail, msg, nil)
		return
	}
	c.JsonUnSafe(code.Success, msg, nil)
}
