package u_email

import (
	"fmt"
	"gen_gin_tpl/pkg/cache"
	"gen_gin_tpl/pkg/enums/em_http"
	"gen_gin_tpl/pkg/resp"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// GetEmailCodeKey 邮箱验证码的缓存key
// 格式：1__email_code__:email
// 参数:
//   - email: 邮箱地址
//
// 返回值:
//   - string: 缓存key
func GetEmailCodeKey(email string) string {
	return fmt.Sprintf("1__email_code__:%s", strings.ToLower(email))

}

// GetEmailCodeValue 获取邮箱验证码的缓存value
// 参数:
//   - email: 邮箱地址
//
// 返回值:
//   - string: 缓存key
func GetEmailCodeValue(email string) string {
	return cache.GetString(GetEmailCodeKey(email))
}

// SetEmailCodeValue 设置邮箱验证码的缓存value
// 参数:
//   - email: 邮箱地址
//   - code: 验证码
//   - expiration: 过期时间
func SetEmailCodeValue(email string, code string, expiration time.Duration) {
	cache.Set(GetEmailCodeKey(email), code, expiration)
}

// IsEmailCodeValue 判断邮箱验证码是否正确
// 参数:
//   - email: 邮箱地址
//   - code: 验证码
//
// 返回值:
//   - bool: 是否正确
func IsEmailCodeValue(email string, code string) bool {
	return strings.EqualFold(GetEmailCodeValue(email), code)
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
		resp.JsonSafeCode(c, em_http.Unknown, em_http.Unknown.Desc(), c.Params)
		return "", false
	}

	s := strings.Split(referer, "?")
	if strings.HasSuffix(s[0], "/regedit.html") {
		return "用户注册验证码", true
	}
	resp.JsonSafeCode(c, em_http.Unknown, em_http.Unknown.Desc(), c.Params)
	return "", false
}
