package variable

import (
	"github.com/mojocn/base64Captcha"
	"sync/atomic"
)

var (
	WebTitle             = "知识库" // 站点标题名称
	ConfigPath           string
	IsEnableEncryptedKey bool
	IsInitialized        atomic.Bool
	CaptchaStore         = base64Captcha.DefaultMemStore
	IsEnableEmail        atomic.Bool
)
