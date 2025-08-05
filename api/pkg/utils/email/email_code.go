package email

import (
	"gen_gin_tpl/pkg/cache"
	"gen_gin_tpl/pkg/enums/code"
	"gen_gin_tpl/pkg/resp"
	"gen_gin_tpl/pkg/session"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// GetEmailId 邮箱验证码的缓存key
//
// 参数:
//   - c: *gin.Context
//   - email: 邮箱地址
//
// 返回值:
//   - string: 缓存key
func GetEmailId(c *gin.Context, email string) string {
	return session.GetCaptchaID(c, strings.ToLower(email))
}

// GetEmailCodeValue 获取邮箱验证码的缓存value
// 参数:
//   - emailId: 邮箱地址 emailId
//
// 返回值:
//   - string: 缓存key
func GetEmailCodeValue(emailId string) string {
	return cache.GetString(emailId)
}

// SetEmailCodeValue 设置邮箱验证码的缓存value
// 参数:
//   - emailId: 邮箱地址 emailId
//   - code: 验证码
//   - expiration: 过期时间
func SetEmailCodeValue(emailId string, code string, expiration time.Duration) {
	_ = cache.Set(emailId, code, expiration)
}

// IsEmailCodeValue 判断邮箱验证码是否正确
// 参数:
//   - emailId: 邮箱地址 emailId
//   - code: 验证码
//
// 返回值:
//   - bool: 是否正确
func IsEmailCodeValue(emailId string, code string) bool {
	return strings.EqualFold(GetEmailCodeValue(emailId), code)
}

// GetEmailTitleTagName 获取邮箱标题
// 参数:
//   - c: gin.Context
//
// 返回值:
//   - string: 标题
//   - bool: 是否成功
func GetEmailTitleTagName(c *gin.Context) (flag string, status bool) {
	referer := c.GetHeader("Referer")
	if referer == "" {
		resp.JsonSafeDesc(c, code.Unknown, c.Params)
		return "", false
	}

	s := strings.Split(referer, "?")
	if strings.HasSuffix(s[0], "/regedit.html") {
		return "用户注册验证码", true
	}
	resp.JsonSafeDesc(c, code.Unknown, c.Params)
	return "", false
}
