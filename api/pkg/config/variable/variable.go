package variable

import (
	"sync/atomic"

	"github.com/mojocn/base64Captcha"
)

var (
	PublicPEM  []byte
	PrivatePEM []byte
	WebTitle   = "知识库"

	ConfigPath    string
	CaptchaStore  = base64Captcha.DefaultMemStore
	IsInitialized atomic.Bool
	IsEnableEmail atomic.Bool
)
