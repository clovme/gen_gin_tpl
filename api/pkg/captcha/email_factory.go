package captcha

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/utils/array"
	utilEmail "gen_gin_tpl/pkg/utils/email"
	"gen_gin_tpl/pkg/variable"
	"gen_gin_tpl/public"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
	"strings"
	"time"
)

// EmailTmpl 邮件模板
type emailTmpl struct{}

// clientEmail 邮件客户端
type clientEmail struct {
	tmpl string // 模板名称
}

// SendCode 生成验证码
//
// 返回值：
//   - string 验证码
func (c *emailTmpl) SendCode(emailId, to, flag string) error {
	charset := strings.Split("0123456789ABCDEFGHJKMNOQRSTUVXYZ", "")
	codeRunes := make([]string, cfg.C.Captcha.Length)
	for i := range codeRunes {
		codeRunes[i] = array.RandomArray[string](charset)
	}
	code := strings.Join(codeRunes, "")

	ce := &clientEmail{tmpl: "validate_code.html"}

	data := map[string]interface{}{
		"Code": code,
	}
	if err := ce.SendEmail([]string{to}, flag, data); err != nil {
		return err
	}
	utilEmail.SetEmailValue(emailId, code, time.Minute)
	return nil
}

// SendEmail 发送邮件
func (c *clientEmail) SendEmail(to []string, subject string, data interface{}) error {
	body, err := c.renderTemplate(data)
	if err != nil {
		return fmt.Errorf("模板渲染失败: %w", err)
	}
	return c.send(to, subject, body, cfg.C.Email)
}

// renderTemplate 渲染模板
func (c *clientEmail) renderTemplate(data interface{}) (string, error) {
	tmplContent, err := public.EmailFS.ReadFile(fmt.Sprintf("email/%s", c.tmpl))
	if err != nil {
		return "", fmt.Errorf("读取模板失败: %w", err)
	}

	tmpl, err := template.New(c.tmpl).Parse(string(tmplContent))
	if err != nil {
		return "", fmt.Errorf("模板解析失败: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("模板执行失败: %w", err)
	}
	return buf.String(), nil
}

// send 发送邮件，基于 jordan-wright/email 封装
func (c *clientEmail) send(to []string, subject, body string, cfg cfg.Email) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", variable.WebTitle, cfg.From)
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)

	addr := fmt.Sprintf("%s:%d", cfg.SMTPHost, cfg.SMTPPort)

	// TLS Config可选，一般 nil 就行，除非你要指定证书校验
	if err := e.SendWithTLS(addr, smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.SMTPHost), &tls.Config{InsecureSkipVerify: true, ServerName: cfg.SMTPHost}); err != nil {
		return fmt.Errorf("发送邮件失败: %w", err)
	}

	return nil
}
