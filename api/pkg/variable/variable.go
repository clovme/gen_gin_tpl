package variable

import (
	"github.com/mojocn/base64Captcha"
	"sync/atomic"
)

var (
	PublicPEM  []byte
	PrivatePEM []byte
	//WebTitle   = constants.ProjectName
	WebTitle   = "知识库"
	SessionKey []byte

	ConfigPath    string
	CaptchaStore  = base64Captcha.DefaultMemStore
	IsInitialized atomic.Bool
	IsEnableEmail atomic.Bool
)
