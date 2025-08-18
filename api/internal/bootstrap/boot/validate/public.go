package validate

import (
	"errors"
	"gen_gin_tpl/internal/infrastructure/query"
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
	return false
}

// UsernameValid 用户名校验器
func UsernameValid(fl validator.FieldLevel) bool {
	value := fl.Field().String()

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

// UniqueEmailValid 邮箱唯一校验器
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

// UniqueUsernameValid 用户名唯一校验器
func UniqueUsernameValid(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	if _, err := query.User.Where(query.User.Username.Eq(value)).First(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true
		}
		return false
	}

	return true
}
