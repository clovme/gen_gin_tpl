package validate

import (
	"errors"
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/internal/schema/dto"
	utilEmail "gen_gin_tpl/pkg/utils/email"
	"gen_gin_tpl/pkg/variable"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/mail"
	"regexp"
)

// EmailValid 邮箱校验器
func EmailValid(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	// 邮箱合法
	if _, err := mail.ParseAddress(value); err == nil {
		return true
	}

	// 字母数字组合合法
	return regexp.MustCompile(`^[a-zA-Z0-9]{5,20}$`).MatchString(value)
}

// PasswordValid 密码校验器
func PasswordValid(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	// 长度判断
	if len(value) < 6 || len(value) > 20 {
		return false
	}

	// 必须包含字母、数字、特殊字符
	hasLetter := regexp.MustCompile(`[A-Za-z]`).MatchString(value)
	hasNumber := regexp.MustCompile(`\d`).MatchString(value)
	hasSpecial := regexp.MustCompile(`[^A-Za-z\d]`).MatchString(value)

	return hasLetter && hasNumber && hasSpecial
}

// CaptchaValid 验证码校验器
func CaptchaValid(fl validator.FieldLevel) bool {
	captchaValue := fl.Field().String()

	// 从整个表单 struct 中拿到 CaptchaId 字段
	_form, ok := fl.Parent().Interface().(dto.RegeditDTO)

	// 校验长度
	if len(captchaValue) != 5 || !ok {
		return false
	}

	// 校验验证码是否匹配
	return variable.CaptchaStore.Verify(_form.CaptchaId, captchaValue, true)
}

// EmailCodeValid 邮箱验证码校验器
func EmailCodeValid(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	// 从整个表单 struct 中拿到 CaptchaId 字段
	_form, ok := fl.Parent().Interface().(dto.RegeditDTO)
	// 校验长度
	if !ok {
		return false
	}

	// 校验验证码是否匹配
	return utilEmail.IsEmailCodeValue(_form.Email, email)
}

func UniqueEmailValid(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	if _, err := query.User.Where(query.User.Email.Eq(value)).First(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true
		}
		return false
	}
	return true
}
