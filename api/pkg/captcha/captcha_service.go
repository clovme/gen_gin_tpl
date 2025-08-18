package captcha

import (
	"encoding/base64"
	"fmt"
	"gen_gin_tpl/pkg/utils/array"
	"github.com/mojocn/base64Captcha"
	"strings"
)

func base64DecodeImage(b64Str string) ([]byte, error) {
	idx := strings.Index(b64Str, ",")
	if idx == -1 {
		return nil, fmt.Errorf("invalid base64 string, no comma found")
	}

	raw := b64Str[idx+1:]
	return base64.StdEncoding.DecodeString(raw)
}

// NewGenerate 生成图形验证码。
//
// 参数：
//   - captchaID 验证码 ID
//
// 返回值：
//   - id     验证码 ID
//   - b64s   验证码图片 Base64 字符串
//   - answer 验证码答案
//   - err    错误信息
func NewGenerate(captchaID string) (imageBytes []byte, err error) {
	captcha := array.RandomArray[*base64Captcha.Captcha](captchaList)

	_, content, answer := captcha.Driver.GenerateIdQuestionAnswer()
	item, err := captcha.Driver.DrawCaptcha(content)
	if err != nil {
		return nil, err
	}

	if err = captcha.Store.Set(captchaID, answer); err != nil {
		return nil, err
	}
	imageBytes, err = base64DecodeImage(item.EncodeB64string())
	return
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
