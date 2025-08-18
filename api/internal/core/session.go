package core

import (
	"fmt"
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/crypto"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"time"
)

const userSessionID = "user_session_id"     // 客户端ID
const browserClientID = "browser_client_id" // 客户端ID

type Session struct {
	session sessions.Session
}

// Set 设置会话值
//
// 参数：
//   - key: 键，用于标识会话值
//   - value: 值，要设置的会话值
func (r *Session) set(key string, value interface{}) {
	r.session.Set(r.id(key), value)
	_ = r.session.Save()
}

// Get 获取会话值
//
// 参数：
//   - key: 键，用于标识会话值
//
// 返回值：
//   - interface{}: 会话值
func (r *Session) get(key string) interface{} {
	return r.session.Get(r.id(key))
}

func (r *Session) DeleteUserSession() {
	r.session.Delete(r.id(userSessionID))
	_ = r.session.Save()
}

// BrowserClientID 获取或生成客户端ID
//
// 返回值：
//   - string: 客户端ID
func (r *Session) BrowserClientID() string {
	clientId := r.session.Get(browserClientID)
	if clientId == nil {
		clientId = base64Captcha.RandomId()
		r.session.Set(browserClientID, clientId)
		_ = r.session.Save()
	}
	return clientId.(string)
}

// id 生成唯一ID
//
// 参数：
//   - key: 键，用于生成唯一ID
//
// 返回值：
//   - string: 唯一ID
//
// 说明：
//   - 该函数用于生成唯一的ID，用于标识用户会话或请求。
//   - 生成的ID基于用户会话ID和键进行加密，确保唯一性和安全性。
func (r *Session) id(key string) string {
	return crypto.Encryption(fmt.Sprintf("%s:%s", r.BrowserClientID(), key))
}

// GetImageCaptchaID 获取图片验证码ID
//
// 返回值：
//   - string: 图片验证码ID
func (r *Session) GetImageCaptchaID() string {
	return r.id("images_captcha_suffix")
}

// GetEmailCaptchaID 获取邮箱验证码ID
//
// 返回值：
//   - string: 邮箱验证码ID
func (r *Session) GetEmailCaptchaID() string {
	return r.id("email_captcha_suffix")
}

// ClearSession 清除会话数据
//
// 说明：
//   - 该函数用于清除会话数据，包括用户会话ID和浏览器客户端ID。
//   - 清除后，会话数据将被删除，用户需要重新登录才能继续使用会话。
func (r *Session) ClearSession() {
	r.session.Clear() // 清掉 session 数据
	r.session.Options(sessions.Options{
		Path:     "/",
		MaxAge:   -1, // 关键：让 cookie 立刻过期
		HttpOnly: true,
	})
	_ = r.session.Save() // 必须 save 才会下发 Set-Cookie
}

func (r *Session) GetUserID(c *gin.Context) (uid int64, ok bool, isToken bool) {
	userID := r.get(r.id(userSessionID))
	if userID != nil {
		return userID.(int64), true, false
	}

	token := c.GetHeader("Token")
	if c.GetHeader("X-Requested-With") != "XMLHttpRequest" || token == "" {
		return 0, false, false
	}

	mapClaims, err := r.ParseUserToken(token)
	if err != nil {
		return 0, false, true
	}

	now := time.Now().Unix()
	iat, ok := mapClaims["iat"].(int64)
	if !ok {
		return 0, false, true
	}
	exp, ok := mapClaims["exp"].(int64)
	if !ok {
		return 0, false, true
	}

	if exp-now <= now-iat {
		return 0, false, true
	}

	return mapClaims["ID"].(int64), true, true
}

func (r *Session) SetUserSession(user *models.User) (userToken string, err error) {
	r.set(r.id(userSessionID), user.ID)
	userToken, err = r.GenUserToken(user)
	if err != nil {
		return "", err
	}
	return userToken, err
}
