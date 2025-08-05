package session

import (
	"fmt"
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/internal/schema/vo/public"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/pwd"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strings"
)

func Set(c *gin.Context, key string, value interface{}) {
	s := sessions.Default(c)
	s.Set(key, value)
	_ = s.Save()
}

func Get(c *gin.Context, key string) interface{} {
	s := sessions.Default(c)
	return s.Get(key)
}

// GetClientID 获取客户端 ID
//
// 参数：
//   - c: Gin 上下文对象，用于获取 session
//
// 返回值：
//   - string: 客户端 ID
func GetClientID(c *gin.Context) string {
	return Get(c, constants.ClientID).(string)
}

// GetCaptchaID 获取验证码 ID
//
// 参数：
//   - c: Gin 上下文对象，用于获取 session
//   - value: 验证码值
//
// 返回值：
//   - string: 加密后的验证码 ID
func GetCaptchaID(c *gin.Context, value string) string {
	clientID := Get(c, constants.ClientID).(string)
	return pwd.Encryption(fmt.Sprintf("%s:%s", clientID, strings.ToLower(value)))
}

// GetUserInfo 获取用户信息
//
// 参数：
//   - c: Gin 上下文对象，用于获取 session
//
// 返回值：
//   - *public.ApiUserVO: 用户信息对象指针，如果获取失败返回 nil
func GetUserInfo(c *gin.Context) *public.ApiUserVO {
	userSessionID := Get(c, constants.UserSessionID)
	if userSessionID == nil {
		return nil
	}
	user, err := query.Q.User.Where(query.User.ID.Eq(userSessionID.(int64))).First()
	if err != nil {
		return nil
	}
	return public.ToVO(*user)
}

// IsLogin 检查用户是否登录
//
// 参数：
//   - c: Gin 上下文对象，用于获取 session
//
// 返回值：
//   - bool: 如果用户已登录返回 true，否则返回 false
func IsLogin(c *gin.Context) bool {
	return Get(c, constants.UserSessionID) != nil
}
