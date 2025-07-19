package web

import (
	"gen_gin_tpl/internal/application"
	"gen_gin_tpl/internal/form/dto"
	"gen_gin_tpl/pkg/enums/em_http"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/page"
	"gen_gin_tpl/pkg/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ViewLoginHandler struct {
	ViewLoginService *application.View_loginService
}

// GetLoginView
// @Router			/login.html [get]
// @Group 			noAuthView
func (h *ViewLoginHandler) GetLoginView(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", page.ViewDataNil("登录"))
}

// PostLoginView
// @Router			/login.html [post]
// @Group 			noAuthView
func (h *ViewLoginHandler) PostLoginView(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", page.ViewDataNil("登录"))
}

// GetRegeditView
// @Router			/regedit.html [get]
// @Group 			noAuthView
func (h *ViewLoginHandler) GetRegeditView(c *gin.Context) {
	c.HTML(http.StatusOK, "regedit.html", page.ViewData[dto.RegeditForm]("注册", dto.RegeditForm{
		CaptchaId:       "xxx",
		Email:           "clovme@qq.com",
		EmailCode:       "xxxxx",
		Password:        "silvery.0",
		ConfirmPassword: "silvery.0",
		Captcha:         "xxxcc",
	}, nil))
}

// PostRegeditView
// @Router			/regedit.html [post]
// @Group 			noAuthView
func (h *ViewLoginHandler) PostRegeditView(c *gin.Context) {
	var regForm dto.RegeditForm
	if err := c.ShouldBind(&regForm); err != nil {
		errMap := regForm.TranslateError(err)
		log.Error().Err(err).Msgf("表单验证错误：%+v", errMap)
		resp.JsonSafeCode(c, em_http.VerifyError, "表单验证错误", errMap)
		return
	}
	// 注册逻辑

	h.ViewLoginService.UserRegeditService()

	resp.JsonSafeCode(c, em_http.Success, "注册成功", nil)
}
