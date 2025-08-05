package validate

import (
	"gen_gin_tpl/internal/schema/dto"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/session"
	"gen_gin_tpl/pkg/utils/email"
	"gen_gin_tpl/pkg/variable"
	"github.com/go-playground/validator/v10"
)

// RegeditCaptchaValid 验证码校验器
func RegeditCaptchaValid(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	// 从整个表单 struct 中拿到 CaptchaId 字段
	_form, ok := fl.Parent().Interface().(dto.RegeditDTO)

	// 校验长度
	if len(value) != cfg.CCaptcha.Length || !ok {
		return false
	}

	captchaId := session.GetCaptchaID(_form.Context, constants.CaptchaSuffix)

	// 校验验证码是否匹配
	return variable.CaptchaStore.Verify(captchaId, value, true)
}

// RegeditEmailCodeValid 邮箱验证码校验器
func RegeditEmailCodeValid(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	// 从整个表单 struct 中拿到 CaptchaId 字段
	_form, ok := fl.Parent().Interface().(dto.RegeditDTO)

	// 校验长度
	if len(value) != cfg.CCaptcha.Length || !ok {
		return false
	}

	emailId := email.GetEmailId(_form.Context, _form.Email)

	// 校验验证码是否匹配
	return email.IsEmailCodeValue(emailId, value)
}
