package initialize

import (
	"gen_gin_tpl/internal/bootstrap/initialize/validate"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// initFormValidate 初始化表单验证器
func initFormValidate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		_ = v.RegisterValidation("emailValid", validate.EmailValid)
		_ = v.RegisterValidation("passwordValid", validate.PasswordValid)
		_ = v.RegisterValidation("captchaValid", validate.CaptchaValid)
		_ = v.RegisterValidation("emailCodeValid", validate.EmailCodeValid)
		_ = v.RegisterValidation("uniqueEmailValid", validate.UniqueEmailValid)
	}
}
