package boot

import (
	"gen_gin_tpl/internal/bootstrap/boot/validate"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// initFormValidate 初始化表单验证器
func initFormValidate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		_ = v.RegisterValidation("emailValid", validate.EmailValid)
		_ = v.RegisterValidation("usernameValid", validate.UsernameValid)
		_ = v.RegisterValidation("passwordValid", validate.PasswordValid)
		_ = v.RegisterValidation("uniqueEmailValid", validate.UniqueEmailValid)
		_ = v.RegisterValidation("uniqueUsernameValid", validate.UniqueUsernameValid)
		// 注册登录验证码校验器
		_ = v.RegisterValidation("regeditCaptchaValid", validate.RegeditCaptchaValid)
		_ = v.RegisterValidation("regeditEmailCodeValid", validate.RegeditEmailCodeValid)
		// 登录验证码校验器
		_ = v.RegisterValidation("loginCaptchaValid", validate.LoginCaptchaValid)
	}
}
