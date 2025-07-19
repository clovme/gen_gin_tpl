package api

import (
	"encoding/base64"
	"gen_gin_tpl/internal/application"
	"gen_gin_tpl/internal/form/dto"
	"gen_gin_tpl/pkg/captcha"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/enums/em_http"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/resp"
	"gen_gin_tpl/pkg/utils/u_email"
	"gen_gin_tpl/public"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type PublicHandler struct {
	PublicService *application.PublicService
}

// PostPublicKey 公钥
// @Router			/public/key [post]
// @Group 			public
func (h *PublicHandler) PostPublicKey(c *gin.Context) {
	data := base64.StdEncoding.EncodeToString(public.PublicPEM)
	for i := 0; i < 10; i++ {
		data = base64.StdEncoding.EncodeToString([]byte(data))
	}
	c.String(http.StatusOK, data)
}

// PostHttpCode 自定义Http状态码
// @Router			/public/enums [post]
// @Group 			public
func (h *PublicHandler) PostHttpCode(c *gin.Context) {
	enums, err := h.PublicService.GetEnums()
	if err != nil {
		return
	}
	resp.JsonSafe(c, em_http.Success.Desc(), enums)
}

// GetPing 自定义Http状态码
// @Router			/public/ping [get]
// @Group 			public
func (h *PublicHandler) GetPing(c *gin.Context) {
	resp.JsonUnSafe(c, em_http.Success.Desc(), nil)
}

// GetSecond 自定义Http状态码
// @Router			/public/time [get]
// @Group 			public
func (h *PublicHandler) GetSecond(c *gin.Context) {
	now := time.Now()
	// 年初
	yearTime := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	// 今日0点
	dayTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	// 当前小时
	hourTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	// 当前分钟
	minuteTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())
	// 当前秒
	secondTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, now.Location())

	resp.JsonSafe(c, em_http.Success.Desc(), gin.H{
		"year":        yearTime.Unix(),
		"day":         dayTime.Unix(),
		"hour":        hourTime.Unix(),
		"minute":      minuteTime.Unix(),
		"second":      secondTime.Unix(),
		"millisecond": now.UnixMilli(),
		"microsecond": now.UnixMicro(),
		"nanosecond":  now.UnixNano(),
		"iso8601":     now.Format(time.RFC3339Nano),
	})
}

// PostCaptcha 生成验证码
// @Router			/public/captcha [post]
// @Group 			public
func (h *PublicHandler) PostCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, _, err := captcha.NewGenerate()
	if err != nil {
		log.Error().Err(err).Msg("验证码生成失败")
		resp.JsonSafeCode(c, http.StatusInternalServerError, "验证码生成失败", nil)
		return
	}

	resp.JsonSafe(c, "验证码生成成功", gin.H{"captchaId": id, "b64s": b64s})
}

// PostSendEmailCaptcha 发送邮箱验证码
// @Router			/public/email/code [post]
// @Group 			public
func (h *PublicHandler) PostSendEmailCaptcha(c *gin.Context) {
	var data dto.EmailCode
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Error().Err(err).Msg("验证码发送失败！")
		resp.JsonSafeCode(c, em_http.VerifyError, "验证码发送失败！", c.Params)
		return
	}
	flag, status := u_email.GetEmailTitleTagName(c)
	if !status {
		return
	}
	if strings.EqualFold(data.Email, cfg.CEmail.From) {
		resp.JsonSafeCode(c, em_http.InternalServerError, em_http.InternalServerError.Desc(), c.Params)
		return
	}
	if u_email.GetEmailCodeValue(data.Email) != "" {
		resp.JsonSafeCode(c, em_http.Unknown, "验证码发送频繁，请稍后再试！", c.Params)
		return
	}
	if err := captcha.NewEmail().SendCode(data.Email, flag); err != nil {
		log.Error().Err(err).Msg("验证码发送失败！")
		resp.JsonSafeCode(c, em_http.Unknown, "验证码发送失败！", c.Params)
		return
	}
	resp.JsonSafe(c, "验证码发送成功！", nil)
}
