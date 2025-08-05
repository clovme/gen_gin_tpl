package variable

import (
	"github.com/mojocn/base64Captcha"
	"sync/atomic"
)

var (
	WebTitle             = "暗香浮动月黄昏" // 站点标题名称
	ConfigPath           string
	IsEnableEncryptedKey bool
	IsInitialized        atomic.Bool
	CaptchaStore         = base64Captcha.DefaultMemStore
	IsEnableEmail        atomic.Bool
)
