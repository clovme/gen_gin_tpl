package api

import (
	"encoding/base64"
	publicService "gen_gin_tpl/internal/application/public"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/internal/schema/dto"
	"gen_gin_tpl/pkg/captcha"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/enums/code"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils/email"
	"gen_gin_tpl/public"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type PublicHandler struct {
	Service *publicService.ApiPublicService
}

// GetPublicKey 公钥
// @Router			/public/key [GET]
// @Group 			public
// @Summary			公钥
// @Type			api
func (r *PublicHandler) GetPublicKey(c *core.Context) {
	data := base64.StdEncoding.EncodeToString(public.PublicPEM)
	for i := 0; i < 10; i++ {
		data = base64.StdEncoding.EncodeToString([]byte(data))
	}
	c.String(http.StatusOK, data)
}

// GetEnumsList 枚举列表
// @Router			/public/enums [GET]
// @Group 			public
// @Summary			枚举列表
// @Type			api
func (r *PublicHandler) GetEnumsList(c *core.Context) {
	enums, err := r.Service.GetAllEnumsData()
	if err != nil {
		c.JsonUnSafeDesc(code.InternalServerError, err.Error())
		return
	}
	c.JsonUnSafeSuccess(enums)
}

// GetPing 心跳
// @Router			/public/ping [GET]
// @Group 			public
// @Summary			心跳
// @Type			api
func (r *PublicHandler) GetPing(c *core.Context) {
	c.JsonUnSafeSuccess(nil)
}

// GetServerTime 服务器时间
// @Router			/public/time [GET]
// @Group 			public
// @Summary			服务器时间
// @Type			api
func (r *PublicHandler) GetServerTime(c *core.Context) {
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

	c.JsonSafeDesc(code.Success, gin.H{
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

// PostSendEmailCaptcha 发送邮箱验证码
// @Router			/public/email/code [POST]
// @Group 			public
// @Summary			发送邮箱验证码
// @Type			api
func (r *PublicHandler) PostSendEmailCaptcha(c *core.Context) {
	var emailCode dto.EmailCode
	if err := c.ShouldBindJSON(&emailCode); err != nil {
		log.Error().Err(err).Msg("验证码发送失败！")
		c.JsonSafe(code.VerifyError, "验证码发送失败！", c.Params)
		return
	}
	flag, status := email.GetEmailTitleTagName(c.Context.GetHeader("Referer"))
	if !status {
		c.JsonSafeDesc(code.Unknown, c.Params)
		return
	}
	if strings.EqualFold(emailCode.Email, cfg.CEmail.From) {
		c.JsonSafeDesc(code.InternalServerError, c.Params)
		return
	}
	if email.GetEmailValue(c.Session.GetImageCaptchaID()) != "" {
		c.JsonSafe(code.Unknown, "验证码发送频繁，请稍后再试！", c.Params)
		return
	}
	if err := captcha.NewEmail().SendCode(c.Session.GetEmailCaptchaID(), emailCode.Email, flag); err != nil {
		log.Error().Err(err).Msg("验证码发送失败！")
		c.JsonSafe(code.Unknown, "验证码发送失败！", c.Params)
		return
	}
	c.JsonSafe(code.Success, "验证码发送成功！", nil)
}
