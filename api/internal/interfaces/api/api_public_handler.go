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
	"gen_gin_tpl/pkg/variable"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type PublicHandler struct {
	Service *publicService.ApiPublicService
}

// GetPublicKey 公钥
// @Type			api
// @Group 			public
// @Router			/public/key [GET]
// @Name			publicKey
// @Summary			公钥
func (r *PublicHandler) GetPublicKey(c *core.Context) {
	data := base64.StdEncoding.EncodeToString(variable.PublicPEM)
	for i := 0; i < 10; i++ {
		data = base64.StdEncoding.EncodeToString([]byte(data))
	}
	c.String(http.StatusOK, data)
}

// GetEnumList 枚举列表
// @Type			api
// @Group 			public
// @Router			/public/enums [GET]
// @Name			enumListApi
// @Summary			枚举列表
func (r *PublicHandler) GetEnumList(c *core.Context) {
	enums, err := r.Service.GetAllEnumsData()
	if err != nil {
		c.JsonUnSafeDesc(code.ServerInternalError, err.Error())
		return
	}
	c.JsonUnSafeSuccess(enums)
}

// GetPing 心跳
// @Type			api
// @Group 			public
// @Router			/public/ping [GET]
// @Name			pingApi
// @Summary			心跳检测
func (r *PublicHandler) GetPing(c *core.Context) {
	c.JsonUnSafeSuccess(nil)
}

// GetServerTime 服务器时间
// @Type			api
// @Group 			public
// @Router			/public/time [GET]
// @Name			serverTimeApi
// @Summary			服务器时间
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
// @Type			api
// @Group 			public
// @Router			/public/email/code [POST]
// @Name			emailCodeApi
// @Summary			发送邮箱验证码
func (r *PublicHandler) PostSendEmailCaptcha(c *core.Context) {
	var emailCode dto.EmailCode
	if err := c.ShouldBindJSON(&emailCode); err != nil {
		log.Error().Err(err).Msg("验证码发送失败！")
		c.JsonSafe(code.ServiceVerifyError, "验证码发送失败！", c.Params)
		return
	}
	flag, status := email.GetEmailTitleTagName(c.Context.GetHeader("Referer"))
	if !status {
		c.JsonSafeDesc(code.RequestUnknown, c.Params)
		return
	}
	if strings.EqualFold(emailCode.Email, cfg.C.Email.From) {
		c.JsonSafeDesc(code.ServerInternalError, c.Params)
		return
	}
	if email.GetEmailValue(c.Session.GetImageCaptchaID()) != "" {
		c.JsonSafe(code.RequestUnknown, "验证码发送频繁，请稍后再试！", c.Params)
		return
	}
	if err := captcha.NewEmail().SendCode(c.Session.GetEmailCaptchaID(), emailCode.Email, flag); err != nil {
		log.Error().Err(err).Msg("验证码发送失败！")
		c.JsonSafe(code.RequestUnknown, "验证码发送失败！", c.Params)
		return
	}
	c.JsonSafe(code.Success, "验证码发送成功！", nil)
}
