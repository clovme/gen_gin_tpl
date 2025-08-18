package core

import (
	"fmt"
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/variable"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// tokenHandler 生成 token
//
// 参数：
//   - user: 用户信息
//   - exp: 过期时间
//
// 返回：
//   - token: 生成的 token
//   - err: 错误信息
func tokenHandler(user *models.User, exp time.Duration) (string, error) {
	// 定义 payload（也叫 claims）
	claims := jwt.MapClaims{
		"ID":  user.ID,
		"iat": time.Now().Unix(),          // 签发时间
		"exp": time.Now().Add(exp).Unix(), // 过期时间
	}

	// 创建 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名并获得完整的 token 字符串
	return token.SignedString(variable.PrivatePEM)
}

// GenUserToken 生成新 token
//
// 参数：
//   - user: 用户信息
//
// 返回：
//   - token: 生成的 token
//   - err: 错误信息
func (r *Session) GenUserToken(user *models.User) (string, error) {
	return tokenHandler(user, time.Hour*24*1)
}

// ParseUserToken 解析 token
//
// 参数：
//   - tokenString: 待解析的 token
//
// 返回：
//   - claims: 解析后的 claims
//   - err: 错误信息
func (r *Session) ParseUserToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return variable.PrivatePEM, nil
	})

	return token.Claims.(jwt.MapClaims), err
}
