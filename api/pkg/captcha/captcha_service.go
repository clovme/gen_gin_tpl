package captcha

import (
	"gen_gin_tpl/pkg/utils/array"
	"github.com/mojocn/base64Captcha"
)

// NewGenerate 生成图形验证码。
//
// 参数：
//   - 无
//
// 返回值：
//   - id     验证码 ID
//   - b64s   验证码图片 Base64 字符串
//   - answer 验证码答案
//   - err    错误信息
func NewGenerate() (id string, b64s string, answer string, err error) {
	captcha := array.RandomArray[*base64Captcha.Captcha](captchaList)

	return captcha.Generate()
}

// NewEmail 创建一个新的邮件客户端。
//
// 参数：
//   - 无
//
// 返回值：
//   - emailTmpl 邮件客户端实例
func NewEmail() *emailTmpl {
	return &emailTmpl{}
}
