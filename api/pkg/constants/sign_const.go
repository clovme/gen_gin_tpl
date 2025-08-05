// 常量定义，常用标志定义, 用于上下文传递，全局使用, 不建议修改
// 用于项目中常使用的一些通用键(签名)，用于设置值和取值

package constants

const (
	ContextIsEncrypted = "is_encrypted_response" // 上下文是否加密的标识
	HeaderEncrypted    = "X-Encrypted"           // 加密标识, 用于标识请求是否加密
	ProjectName        = "gen_gin_tpl"           // 项目名称
	PublicPEM          = "public_pem"            // 公钥, 用于加密解密
	PrivatePEM         = "private_pem"           // 私钥, 用于加密解密
	WebTitle           = "web_title"             // 站点标题标志
	Countdown          = "countdown"             // 倒计时标记
	ClientID           = "client_id"             // 客户端ID
	CaptchaSuffix      = "images_captcha_suffix" // 图片验证码后缀
	UserSessionID      = "user_session_id"       // 用户会话ID
)
