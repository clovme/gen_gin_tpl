package email

import (
	"crypto/tls"
	"fmt"
	"gen_gin_tpl/pkg/logger/log"
	"net"
	"net/smtp"
	"time"
)

// CheckSMTPConnection SMTP 连接检测
// 参数：
//   - host: SMTP服务器地址
//   - port: SMTP服务器端口号
//
// 返回值：
//   - bool: 连接成功返回true，失败返回false
func CheckSMTPConnection(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		log.Error().Err(err).Msg("SMTP连接失败，请检查邮件服务器地址和端口号。")
		return false
	}
	defer conn.Close()
	return true
}

// CheckSMTPAuth SMTP 认证检测
// 参数：
//   - host: SMTP服务器地址
//   - port: SMTP服务器端口号
//   - username: SMTP用户名
//   - password: SMTP密码
//
// 返回值：
//   - bool: 认证成功返回true，失败返回false
func CheckSMTPAuth(host string, port int, username, password string) bool {
	// 建立TCP连接
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := tls.Dial("tcp", address, nil)
	if err != nil {
		log.Error().Err(err).Msg("无法连接到SSL")
		return false
	}
	defer conn.Close()

	// 建立SMTP客户端
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Error().Err(err).Msg("创建SMTP客户端失败")
		return false
	}
	defer c.Quit()

	// 认证
	auth := smtp.PlainAuth("", username, password, host)
	if err = c.Auth(auth); err != nil {
		log.Error().Err(err).Msg("SMTP账号认证失败")
		return false
	}

	log.Info().Msg("SMTP连接和认证成功")
	return true
}
