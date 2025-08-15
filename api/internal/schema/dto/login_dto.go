package dto

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type LoginDTO struct {
	CaptchaID string `json:"-"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Captcha   string `json:"captcha" binding:"required,loginCaptchaValid"`
}

func (r LoginDTO) TranslateError(err error) map[string]string {
	result := make(map[string]string)

	// 判断是不是验证错误
	var errs validator.ValidationErrors
	if !errors.As(err, &errs) {
		result["other"] = err.Error()
		return result
	}

	// 遍历所有验证错误
	for _, e := range errs {
		field := e.Field()
		tag := e.Tag()

		switch field {
		case "Username":
			switch tag {
			case "required":
				result["username"] = "用户名不能为空"
			}
		case "Password":
			switch tag {
			case "required":
				result["password"] = "密码不能为空"
			}
		case "Captcha":
			switch tag {
			case "required":
				result["captcha"] = "验证码不能为空"
			case "loginCaptchaValid":
				result["captcha"] = "验证码错误"
			}
		}
	}
	return result
}
