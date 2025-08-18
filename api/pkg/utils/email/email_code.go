package email

import (
	"gen_gin_tpl/pkg/cache"
	"gen_gin_tpl/pkg/enums/code"
	"strings"
	"time"
)

// GetEmailValue 获取邮箱验证码的缓存value
// 参数:
//   - emailId: 邮箱地址 emailId
//
// 返回值:
//   - string: 缓存key
func GetEmailValue(emailId string) string {
	return cache.GetString(emailId)
}

// SetEmailValue 设置邮箱验证码的缓存value
// 参数:
//   - emailId: 邮箱地址 emailId
//   - code: 验证码
//   - expiration: 过期时间
func SetEmailValue(emailId string, code string, expiration time.Duration) {
	cache.Set(emailId, code, expiration)
}

// IsEmailValue 判断邮箱验证码是否正确
// 参数:
//   - emailId: 邮箱地址 emailId
//   - code: 验证码
//
// 返回值:
//   - bool: 是否正确
func IsEmailValue(emailId string, code string) bool {
	return strings.EqualFold(GetEmailValue(emailId), code)
}

// GetEmailTitleTagName 获取邮箱标题
// 参数:
//   - referer: GetHeader("Referer") 请求头中的Referer防伪链接
//
// 返回值:
//   - string: 标题
//   - bool: 是否成功
func GetEmailTitleTagName(referer string) (flag string, status bool) {
	if referer == "" {
		return code.RequestUnknown.Desc(), false
	}

	s := strings.Split(referer, "?")
	if strings.HasSuffix(s[0], "/regedit.html") {
		return "用户注册验证码", true
	}
	return code.RequestUnknown.Desc(), false
}
