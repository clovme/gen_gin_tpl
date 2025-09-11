package initweb

import (
	"fmt"
	"gen_gin_tpl/internal/bootstrap/boot"
	"gen_gin_tpl/internal/bootstrap/database"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/copyright"
	"gen_gin_tpl/pkg/enums/code"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils"
	"gen_gin_tpl/pkg/utils/email"
	"gen_gin_tpl/pkg/variable"
	"github.com/mojocn/base64Captcha"
	"strings"
)

func copyrightHandler(c *core.Context) {
	c.JsonUnSafeSuccess(copyright.NewCopyright())
}

func viewHandler(c *core.Context) {
	c.HTML("index.html", "系统初始化", nil)
}

func formHandler(c *core.Context) {
	loggerLevel := []*Options{
		{Label: "trace", Value: "trace", Desc: "细粒度最高，最大量日志"},
		{Label: "debug", Value: "debug", Desc: "调试日志"},
		{Label: "info", Value: "info", Desc: "常规运行状态日志"},
		{Label: "warn", Value: "warn", Desc: "警告，非致命异常"},
		{Label: "error", Value: "error", Desc: "错误日志，功能异常"},
		{Label: "fatal", Value: "fatal", Desc: "致命错误，程序即将终止"},
		{Label: "panic", Value: "panic", Desc: "更严重，触发 panic 行为"},
		{Label: "no", Value: "no", Desc: "没有级别，适合特殊用途"},
		{Label: "disabled", Value: "disabled", Desc: "禁止所有日志"},
	}
	Web := []*Form{
		{Field: "WebTitle", Title: "网站标题", Span: 24, ItemRender: ItemRender{Name: "VxeInput"}},
		{Field: "WebIP", Title: "监听地址", Span: 12, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Placeholder: "Web 服务监听地址，IP 地址，支持多个，逗号分隔"}}},
		{Field: "WebPort", Title: "监听端口", Span: 12, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Type: numberType, Placeholder: "Web 服务监听端口"}}},
		{Field: "WebMode", Title: "服务模式", Span: 24, ItemRender: ItemRender{Name: vxeSelect, Options: []*Options{{Label: "生产模式", Value: "release"}, {Label: "调试模式", Value: "debug"}}}},
	}
	Redis := []*Form{
		{Field: "RedisHost", Title: "主机地址", Span: 12, ItemRender: ItemRender{Name: "VxeInput"}},
		{Field: "RedisPort", Title: "端口号", Span: 12, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Type: numberType}}},
		{Field: "RedisUsername", Title: "用户名", Span: 24, ItemRender: ItemRender{Name: "VxeInput"}},
		{Field: "RedisPassword", Title: "密码", Span: 24, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Type: password}}},
		{Field: "RedisDB", Title: "选择数据库", Span: 24, ItemRender: ItemRender{Name: "VxeInput"}},
	}
	SQLite := []*Form{
		{Field: "SQLiteDbName", Title: "数据库名称", Span: 24, ItemRender: ItemRender{Name: "VxeInput"}},
	}
	MySQL := []*Form{
		{Field: "MySQLHost", Title: "主机地址", Span: 12, ItemRender: ItemRender{Name: "VxeInput"}},
		{Field: "MySQLPort", Title: "端口号", Span: 12, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Type: numberType}}},
		{Field: "MySQLUsername", Title: "用户名", Span: 24, ItemRender: ItemRender{Name: "VxeInput"}},
		{Field: "MySQLPassword", Title: "密码", Span: 24, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Type: password}}},
		{Field: "MySQLDbName", Title: "数据库名称", Span: 24, ItemRender: ItemRender{Name: "VxeInput"}},
	}
	Email := []*Form{
		{Field: "EmailSMTPHost", Title: "主机地址", Span: 12, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Placeholder: "eg:smtp.qq.com"}}},
		{Field: "EmailSMTPPort", Title: "端口号", Span: 12, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Type: numberType}}},
		{Field: "EmailUsername", Title: "用户名", Span: 24, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Placeholder: "eg:admin@qq.com"}}},
		{Field: "EmailPassword", Title: "授权码", Span: 24, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Type: password, Placeholder: "请输入邮箱授权码"}}},
		{Field: "EmailFrom", Title: "发件地址", Span: 24, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Placeholder: "eg:admin@qq.com,一般和用户名一致"}}},
	}
	Logger := []*Form{
		{Field: "LoggerLevel", Title: "日志级别", Span: 12, ItemRender: ItemRender{Name: "VxeTableSelect", Props: &Props{Columns: []*Columns{{Field: "label", Title: "日志级别"}, {Field: "desc", Title: "日志描述"}}}, Options: loggerLevel}},
		{Field: "LoggerMaxSize", Title: "分割大小(MB)", Span: 12, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Type: numberType}}},
		{Field: "LoggerLogPath", Title: "日志路径", Span: 24, ItemRender: ItemRender{Name: "VxeInput"}},
		{Field: "LoggerFormatJson", Title: "JSON/文本", Span: 8, ItemRender: ItemRender{Name: vxeSelect, Options: []*Options{{Label: "JSON", Value: true}, {Label: "文本", Value: false}}}},
		{Field: "LoggerCompress", Title: "启用压缩", Span: 16, ItemRender: ItemRender{Name: vxeSelect, Options: []*Options{{Label: "启用", Value: true}, {Label: "禁用", Value: false}}}},
		{Field: "LoggerMaxAge", Title: "保存天数(天)", Span: 12, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Type: numberType}}},
		{Field: "LoggerMaxBackups", Title: "保留数量(个)", Span: 12, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Type: numberType}}},
	}
	Captcha := []*Form{
		{Field: "CaptchaLength", Title: "验证码长度", Span: 12, ItemRender: ItemRender{Name: vxeSelect, Options: []*Options{{Label: "4位", Value: 4}, {Label: "5位", Value: 5}, {Label: "6位", Value: 6}}}},
		{Field: "CaptchaNoiseCount", Title: "噪点数量", Span: 12, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Type: numberType}}},
		{Field: "CaptchaType", Title: "验证码类型", Span: 24, ItemRender: ItemRender{Name: "VxeCheckboxGroup", Options: []*Options{
			{Label: "数字码", Value: "digit"},
			{Label: "数字字母码", Value: "alphaNum"},
			{Label: "汉字码", Value: "chinese"},
			{Label: "算术运算码", Value: "math"},
		}}},
		{Field: "CaptchaFonts", Title: "验证码字体", Span: 24, ItemRender: ItemRender{Name: "VxeCheckboxGroup", Props: &Props{ClassName: "grid-lines-5-box"}, Options: []*Options{
			{Label: "3Dumb", Value: "3Dumb"},
			{Label: "actionj", Value: "actionj"},
			{Label: "ApothecaryFont", Value: "ApothecaryFont"},
			{Label: "chromohv", Value: "chromohv"},
			{Label: "Comismsh", Value: "Comismsh"},
			{Label: "DeborahFancyDress", Value: "DeborahFancyDress"},
			{Label: "DENNEthree-dee", Value: "DENNEthree-dee"},
			{Label: "Flim-Flam", Value: "Flim-Flam"},
			{Label: "RitaSmith", Value: "RitaSmith"},
			{Label: "wqy-microhei", Value: "wqy-microhei"},
		}}},
		{Field: "CaptchaShowLine", Title: "干扰线条", Span: 24, ItemRender: ItemRender{Name: "VxeCheckboxGroup", Props: &Props{ClassName: "grid-lines-4-box"}, Options: []*Options{
			{Label: "无干扰线条", Value: 0},
			{Label: "空心线", Value: base64Captcha.OptionShowHollowLine},
			{Label: "粘稠线", Value: base64Captcha.OptionShowSlimeLine},
			{Label: "正弦波线", Value: base64Captcha.OptionShowSineLine},
			{Label: "空心+粘稠线", Value: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine},
			{Label: "空心+正弦波线", Value: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSineLine},
			{Label: "粘稠+正弦波线", Value: base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowSineLine},
			{Label: "空心+粘稠+正弦波线", Value: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowSineLine},
		}}},
	}

	OtherDbType := []*Options{{Label: constants.SQLite, Value: constants.SQLite}, {Label: constants.MySQL, Value: constants.MySQL}}
	OtherCacheType := []*Options{{Label: constants.Memory, Value: constants.Memory}, {Label: constants.Redis, Value: constants.Redis}}
	OtherIsEmail := []*Options{{Label: "启用", Value: true}, {Label: "禁用", Value: false}}

	rules := map[string][]Rules{
		"OtherDbType":    {{Required: true, Message: "必须选择数据库类型"}},
		"OtherCacheType": {{Required: true, Message: "必须选择缓存数据库"}},
		"OtherIsEmail":   {{Required: true, Message: "必须选择是否使用邮箱"}},

		"WebTitle": {{Required: true, Message: "网站标题不能为空"}},

		"SQLiteDbName": {{Required: true, Message: "SQLite 数据库名称"}},

		"MySQLHost":     {{Required: true, Message: "请输入主机地址"}},
		"MySQLPort":     {{Required: true, Type: numberType, Min: 1024, Max: 65535, Message: "请输入(1024~65535)范围内的端口号"}},
		"MySQLUsername": {{Required: true, Message: "请输入用户名"}},
		"MySQLDbName":   {{Required: true, Message: "请输入数据库名称"}},

		"WebIP":   {{Required: true, Message: "请输入IP地址"}},
		"WebPort": {{Required: true, Type: numberType, Min: 70, Max: 65535, Message: "请输入(0~65535)范围内的端口号"}},
		"WebMode": {{Required: true, Message: "请选择服务模式"}},

		"RedisHost": {{Required: true, Message: "请输入主机地址"}},
		"RedisPort": {{Required: true, Type: numberType, Min: 1024, Max: 65535, Message: "请输入(1024~65535)范围内的端口号"}},
		"RedisDB":   {{Required: true, Type: numberType, Min: 0, Max: 15, Message: "请输入(0~15)范围内的端口号"}},

		"EmailSMTPHost": {{Required: true, Message: "邮件服务器地址(smtp.qq.com)"}},
		"EmailSMTPPort": {{Required: true, Type: numberType, Min: 1, Max: 65535, Message: "请输入(0~65535)范围内的端口号"}},
		"EmailUsername": {{Required: true, Message: "请输入正确的邮箱地址", Validator: "EMAIL_ADDRESS"}},
		"EmailPassword": {{Required: true, Message: "授权码不能为空"}},
		"EmailFrom":     {{Required: true, Message: "发件地址不能为空", Validator: "EMAIL_ADDRESS"}},

		"LoggerLogPath":    {{Required: true, Message: "请输入日志存放路径"}},
		"LoggerMaxSize":    {{Required: true, Type: numberType, Min: 1, Message: "分割大小(>MB)"}},
		"LoggerMaxAge":     {{Required: true, Type: numberType, Min: 1, Message: "保存天数(天) > 1"}},
		"LoggerMaxBackups": {{Required: true, Type: numberType, Min: 1, Message: "旧日志数量 > 0"}},

		"CaptchaLength":     {{Required: true, Type: numberType, Min: 4, Max: 6, Message: "验证码长度 4~6"}},
		"CaptchaNoiseCount": {{Required: true, Type: numberType, Min: 0, Max: 100, Message: "噪点数量 0~100"}},
		"CaptchaType":       {{Required: true, Message: "至少选择一个验证码类型选项"}},
		"CaptchaFonts":      {{Required: true, Message: "至少选择一个字体选项"}},
		"CaptchaShowLine":   {{Required: true, Message: "至少选择一个干扰线选项"}},
	}

	form := FormOptions{
		Border:          true,
		TitleColon:      false,
		TitleWidth:      120,
		TitleAlign:      "right",
		TitleBackground: true,
		ValidConfig: ValidConfig{
			Theme: normalTheme,
		},
		FormData: FormData{
			WebTitle: variable.WebTitle,

			OtherIsEmail:   cfg.C.Other.IsEmail,
			OtherDbType:    cfg.C.Other.DbType,
			OtherCacheType: cfg.C.Other.CacheType,
			OtherDataPath:  cfg.C.Other.DataPath,

			SQLiteDbName: cfg.C.SQLite.DbName,

			MySQLHost:     cfg.C.MySQL.Host,
			MySQLPort:     cfg.C.MySQL.Port,
			MySQLUsername: cfg.C.MySQL.Username,
			MySQLPassword: cfg.C.MySQL.Password,
			MySQLDbName:   cfg.C.MySQL.DbName,

			WebIP:   cfg.C.Web.IP,
			WebPort: cfg.C.Web.Port,
			WebMode: cfg.C.Web.Mode,

			RedisHost:     cfg.C.Redis.Host,
			RedisPort:     cfg.C.Redis.Port,
			RedisUsername: cfg.C.Redis.Username,
			RedisPassword: cfg.C.Redis.Password,
			RedisDB:       cfg.C.Redis.DB,

			EmailSMTPHost: cfg.C.Email.SMTPHost,
			EmailSMTPPort: cfg.C.Email.SMTPPort,
			EmailUsername: cfg.C.Email.Username,
			EmailPassword: cfg.C.Email.Password,
			EmailFrom:     cfg.C.Email.From,

			LoggerLevel:      cfg.C.Logger.Level,
			LoggerLogPath:    cfg.C.Logger.LogPath,
			LoggerFormatJson: cfg.C.Logger.FormatJSON,
			LoggerCompress:   cfg.C.Logger.Compress,
			LoggerMaxSize:    cfg.C.Logger.MaxSize,
			LoggerMaxAge:     cfg.C.Logger.MaxAge,
			LoggerMaxBackups: cfg.C.Logger.MaxBackups,

			CaptchaLength:     cfg.C.Captcha.Length,
			CaptchaNoiseCount: cfg.C.Captcha.NoiseCount,
			CaptchaType:       cfg.C.Captcha.Type,
			CaptchaFonts:      cfg.C.Captcha.Fonts,
			CaptchaShowLine:   cfg.C.Captcha.ShowLine,
		},
		Rules: rules,
		FormItems: []FormItems{
			{Span: 24, Vertical: true, TitleBold: true, Title: "数据库选择", Children: []*Form{
				{Field: "OtherDbType", Title: "数据库类型", Span: 8, ItemRender: ItemRender{Name: vxeSelect, Options: OtherDbType}},
				{Field: "OtherCacheType", Title: "缓存数据库", Span: 8, ItemRender: ItemRender{Name: vxeSelect, Options: OtherCacheType}},
				{Field: "OtherIsEmail", Title: "启用邮件", Span: 8, ItemRender: ItemRender{Name: vxeSelect, Options: OtherIsEmail}},
			}},
			{Span: 24, Vertical: true, TitleBold: true, Title: "Web 配置", Children: Web},
			{Span: 24, Vertical: true, TitleBold: true, Title: "SQLite 配置", Children: SQLite, ShowWhen: &ShowWhen{Field: "OtherDbType", Value: constants.SQLite}},
			{Span: 24, Vertical: true, TitleBold: true, Title: "MySQL 配置", Children: MySQL, ShowWhen: &ShowWhen{Field: "OtherDbType", Value: constants.MySQL}},
			{Span: 24, Vertical: true, TitleBold: true, Title: "Redis 配置", Children: Redis, ShowWhen: &ShowWhen{Field: "OtherCacheType", Value: constants.Redis}},
			{Span: 24, Vertical: true, TitleBold: true, Title: "Email 配置", Children: Email, ShowWhen: &ShowWhen{Field: "OtherIsEmail", Value: true}},
			{Span: 24, Vertical: true, TitleBold: true, Title: "验证码配置", Children: Captcha},
			{Span: 24, Vertical: true, TitleBold: true, Title: "系统日志配置", Children: Logger},
			{Span: 24, Vertical: true, TitleBold: true, Title: "其他配置", Children: []*Form{
				{Field: "OtherDataPath", Title: "数据存放位置", Span: 24, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Placeholder: "数据存放位置，可以是相对路径，也可以是绝对路径"}}},
			}},
			{Span: 24, Children: []*Form{
				{Align: "center", Span: 24, ItemRender: ItemRender{Name: "VxeButtonGroup", Options: []*Options{{Type: "submit", Content: "保存配置", Status: "primary"}}}},
			}},
		},
	}
	c.JsonUnSafeSuccess(form)
}

// postHandler 验证表单，并保存配置
func postHandler(c *core.Context) {
	var formData FormData
	if err := c.ShouldBindJSON(&formData); err != nil {
		log.Error().Err(err).Msgf("参数格式错误: %+v", err)
		c.JsonUnSafe(code.RequestUnknown, fmt.Sprintf("参数格式错误: %+v", err), nil)
		return
	}

	variable.WebTitle = formData.WebTitle

	cfg.C.SQLite.DbName = formData.SQLiteDbName

	cfg.C.MySQL.Host = formData.MySQLHost
	cfg.C.MySQL.Port = formData.MySQLPort
	cfg.C.MySQL.Username = formData.MySQLUsername
	cfg.C.MySQL.Password = formData.MySQLPassword
	cfg.C.MySQL.DbName = formData.MySQLDbName

	cfg.C.Web.IP = formData.WebIP
	cfg.C.Web.Port = formData.WebPort
	cfg.C.Web.Mode = formData.WebMode

	cfg.C.Redis.Host = formData.RedisHost
	cfg.C.Redis.Port = formData.RedisPort
	cfg.C.Redis.Username = formData.RedisUsername
	cfg.C.Redis.Password = formData.RedisPassword
	cfg.C.Redis.DB = formData.RedisDB

	cfg.C.Email.SMTPHost = formData.EmailSMTPHost
	cfg.C.Email.SMTPPort = formData.EmailSMTPPort
	cfg.C.Email.Username = formData.EmailUsername
	cfg.C.Email.Password = formData.EmailPassword
	cfg.C.Email.From = formData.EmailFrom

	cfg.C.Logger.Level = formData.LoggerLevel
	cfg.C.Logger.MaxSize = formData.LoggerMaxSize
	cfg.C.Logger.LogPath = formData.LoggerLogPath
	cfg.C.Logger.FormatJSON = formData.LoggerFormatJson
	cfg.C.Logger.Compress = formData.LoggerCompress
	cfg.C.Logger.MaxAge = formData.LoggerMaxAge
	cfg.C.Logger.MaxBackups = formData.LoggerMaxBackups

	cfg.C.Captcha.Length = formData.CaptchaLength
	cfg.C.Captcha.NoiseCount = formData.CaptchaNoiseCount
	cfg.C.Captcha.Type = formData.CaptchaType
	cfg.C.Captcha.Fonts = formData.CaptchaFonts
	cfg.C.Captcha.ShowLine = formData.CaptchaShowLine

	cfg.C.Other.IsEmail = formData.OtherIsEmail
	cfg.C.Other.DbType = formData.OtherDbType
	cfg.C.Other.CacheType = formData.OtherCacheType
	cfg.C.Other.DataPath = formData.OtherDataPath

	boot.InitializationLogger(cfg.C.Logger)

	// 检测数据库链接
	if strings.EqualFold(cfg.C.Other.DbType, constants.MySQL) {
		db, err := database.CheckDbConnect(cfg.C.MySQL.Username, cfg.C.MySQL.Password, cfg.C.MySQL.Host, cfg.C.MySQL.Port)
		if err != nil {
			return
		}
		defer db.Close()
	}

	// 检测缓存链接
	if strings.EqualFold(cfg.C.Other.CacheType, constants.Redis) {
		if err := utils.CheckRedisConn(cfg.C.Redis.Host, cfg.C.Redis.Port, cfg.C.Redis.Username, cfg.C.Redis.Password, cfg.C.Redis.DB); err != nil {
			log.Error().Err(err).Msg("Redis 连接失败")
			return
		}
	}

	// 验证是否能够链接邮件服务器
	if formData.OtherIsEmail {
		if !email.CheckSMTPConnection(cfg.C.Email.SMTPHost, cfg.C.Email.SMTPPort) {
			return
		}

		if !email.CheckSMTPAuth(cfg.C.Email.SMTPHost, cfg.C.Email.SMTPPort, cfg.C.Email.Username, cfg.C.Email.Password) {
			return
		}
	}

	variable.IsInitialized.Store(true)
}
