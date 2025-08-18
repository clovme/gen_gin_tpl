package variable

import (
	"gen_gin_tpl/pkg/constants"
	"github.com/mojocn/base64Captcha"
	"sync/atomic"
)

var (
	PublicPEM  []byte
	PrivatePEM []byte
	WebTitle   = constants.ProjectName

	ConfigPath    string
	CaptchaStore  = base64Captcha.DefaultMemStore
	IsInitialized atomic.Bool
	IsEnableEmail atomic.Bool
)
