package dto

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type RegeditDTO struct {
	CaptchaId       string `json:"captcha_id" binding:"required"`
	Email           string `json:"email" binding:"required,emailValid,uniqueEmailValid"`
	EmailCode       string `json:"email_code" binding:"required,emailCodeValid"`
	Password        string `json:"password" binding:"required,passwordValid"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	Captcha         string `json:"captcha" binding:"required,captchaValid"`
}

func (l RegeditDTO) TranslateError(err error) map[string]string {
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
		case "CaptchaId":
			switch tag {
			case "required":
				result["captcha_id"] = "验证码ID不能为空"
			}
		case "Email":
			switch tag {
			case "required":
				result["email"] = "邮箱不能为空"
			case "emailValid":
				result["email"] = "您输入的邮箱格式不正确"
			case "uniqueEmailValid":
				result["email"] = "邮箱已存在，请检查或使用其他邮箱"
			}
		case "EmailCode":
			switch tag {
			case "required":
				result["email_code"] = "邮箱验证码不能为空"
			case "emailCodeValid":
				result["email_code"] = "邮箱验证码错误或者已过期"
			}
		case "Password":
			switch tag {
			case "required":
				result["password"] = "密码不能为空"
			case "passwordValid":
				result["password"] = "密码必须包含字母、数字和特殊字符，长度6-20"
			}
		case "ConfirmPassword":
			switch tag {
			case "required":
				result["confirm_password"] = "确认密码不能为空"
			case "eqfield":
				result["confirm_password"] = "两次输入的密码不一致"
			}
		case "Captcha":
			switch tag {
			case "required":
				result["captcha"] = "验证码不能为空"
			case "captchaValid":
				result["captcha"] = "验证码错误"
			}
		}
	}

	return result
}
