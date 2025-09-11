// 常量定义，常用标志定义, 用于上下文传递，全局使用, 不建议修改
// 用于项目中常使用的一些通用键(签名)，用于设置值和取值

package constants

const (
	ProjectName     = "gen_gin_tpl" // 项目名称
	HeaderEncrypted = "X-Encrypted" // 加密标识, 用于标识请求是否加密
	HttpLogKey      = "HTTP_LOG_KEY"

	WebTitle           = "WEB_TITLE"             // 站点标题标志
	PublicPEM          = "PUBLIC_PEM"            // 公钥, 用于加密解密
	Countdown          = "COUNTDOWN"             // 倒计时标记
	PrivatePEM         = "PRIVATE_PEM"           // 私钥, 用于加密解密
	SessionKey         = "SESSION_KEY"           // 会话密钥, 用于加密解密
	ContextIsEncrypted = "IS_ENCRYPTED_RESPONSE" // 上下文是否加密的标识

	IsContextAjax   = "IS_CONTEXT_AJAX"   // Ajax 请求标识
	IsContextLogin  = "IS_CONTEXT_LOGIN"  // 用户登录状态标识
	ContextUserInfo = "CONTEXT_USER_INFO" // 用户信息
)
