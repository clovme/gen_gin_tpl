package with

import (
	"fmt"
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/pwd"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"time"
)

const userSessionID = "user_session_id"     // 客户端ID
const browserClientID = "browser_client_id" // 客户端ID

type Session struct {
	session     sessions.Session
	isDebugging bool
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

func (r *Session) del(key string) {
	r.session.Delete(r.id(key))
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
	return pwd.Encryption(fmt.Sprintf("%s:%s", r.BrowserClientID(), key))
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

func (r *Session) DelUserSession() {
	r.del(userSessionID)
}

func (r *Session) GetUserID(c *gin.Context) (uid int64, ok bool, isToken bool) {
	userID := r.get(r.id(userSessionID))
	if userID != nil {
		return userID.(int64), true, false
	}
	username, token, ok := c.Request.BasicAuth()
	if !ok {
		return 0, false, true
	}
	mapClaims, err := parseUserToken(token)
	if err != nil {
		return 0, false, true
	}

	now := time.Now().Unix()
	iat := mapClaims["iat"].(int64)
	exp := mapClaims["exp"].(int64)

	if exp-now <= now-iat {
		return 0, false, true
	}

	return mapClaims[username].(int64), true, true
}

func (r *Session) SetUserSession(user models.User) (userToken string, err error) {
	r.set(r.id(userSessionID), user.ID)
	if !r.isDebugging {
		return "", nil
	}
	userToken, err = genUserToken(user)
	if err != nil {
		return "", err
	}
	//c.Request.SetBasicAuth("ID", userToken)
	return userToken, err
}

func (r *Session) SetBasicAuth() {

}
