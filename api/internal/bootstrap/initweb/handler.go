package initweb

import (
	"fmt"
	"gen_gin_tpl/internal/bootstrap/database"
	"gen_gin_tpl/internal/bootstrap/initialize"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/copyright"
	"gen_gin_tpl/pkg/enums/code"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils"
	"gen_gin_tpl/pkg/utils/email"
	"gen_gin_tpl/pkg/variable"
	"gen_gin_tpl/pkg/with"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
	"strings"
)

func copyrightHandler(c *gin.Context) {
	with.NewWithContextInitWeb(c, func(c *with.Context[any]) {
		c.JsonUnSafeSuccess(copyright.NewCopyright())
	})
}

func viewHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": variable.WebTitle,
	})
}

func formHandler(c *gin.Context) {
	with.NewWithContextInitWeb(c, func(c *with.Context[any]) {
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
			{Field: "WebHost", Title: "监听地址", Span: 14, ItemRender: ItemRender{Name: "VxeInput"}},
			{Field: "WebPort", Title: "监听端口", Span: 10, ItemRender: ItemRender{Name: "VxeInput", Props: &Props{Type: numberType}}},
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
			"WebTitle":          {{Required: true, Message: "网站标题不能为空"}},
			"OtherDataPath":     {{Required: true, Message: "请选择数据存放位置"}},
			"SQLiteDbName":      {{Required: true, Message: "SQLite 数据库名称"}},
			"MySQLHost":         {{Required: true, Message: "请输入主机地址"}},
			"MySQLPort":         {{Required: true, Type: numberType, Min: 1024, Max: 65535, Message: "请输入(1024~65535)范围内的端口号"}},
			"MySQLUsername":     {{Required: true, Message: "请输入用户名"}},
			"MySQLDbName":       {{Required: true, Message: "请输入数据库名称"}},
			"WebHost":           {{Required: true, Message: "请输入主机地址"}},
			"WebPort":           {{Required: true, Type: numberType, Min: 1024, Max: 65535, Message: "请输入(1024~65535)范围内的端口号"}},
			"WebMode":           {{Required: true, Message: "请选择服务模式"}},
			"RedisHost":         {{Required: true, Message: "请输入主机地址"}},
			"RedisPort":         {{Required: true, Type: numberType, Min: 1024, Max: 65535, Message: "请输入(1024~65535)范围内的端口号"}},
			"RedisDB":           {{Required: true, Type: numberType, Min: 0, Max: 15, Message: "请输入(0~15)范围内的端口号"}},
			"EmailSMTPHost":     {{Required: true, Message: "邮件服务器地址(smtp.qq.com)"}},
			"EmailSMTPPort":     {{Required: true, Type: numberType, Min: 1, Max: 65535, Message: "请输入(0~65535)范围内的端口号"}},
			"EmailUsername":     {{Required: true, Message: "请输入正确的邮箱地址", Validator: "EMAIL_ADDRESS"}},
			"EmailPassword":     {{Required: true, Message: "授权码不能为空"}},
			"EmailFrom":         {{Required: true, Message: "发件地址不能为空", Validator: "EMAIL_ADDRESS"}},
			"LoggerLogPath":     {{Required: true, Message: "请输入日志存放路径"}},
			"LoggerMaxSize":     {{Required: true, Type: numberType, Min: 1, Message: "分割大小(>MB)"}},
			"LoggerMaxAge":      {{Required: true, Type: numberType, Min: 1, Message: "保存天数(天) > 1"}},
			"LoggerMaxBackups":  {{Required: true, Type: numberType, Min: 1, Message: "旧日志数量 > 0"}},
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
			ValidConfig: &ValidConfig{
				Theme: normalTheme,
			},
			FormData: &FormData{
				WebTitle:          variable.WebTitle,
				OtherIsEmail:      cfg.COther.IsEmail,
				OtherDbType:       cfg.COther.DbType,
				OtherCacheType:    cfg.COther.CacheType,
				OtherDataPath:     cfg.COther.DataPath,
				SQLiteDbName:      cfg.CSQLite.DbName,
				MySQLHost:         cfg.CMySQL.Host,
				MySQLPort:         cfg.CMySQL.Port,
				MySQLUsername:     cfg.CMySQL.Username,
				MySQLPassword:     cfg.CMySQL.Password,
				MySQLDbName:       cfg.CMySQL.DbName,
				WebHost:           cfg.CWeb.Host,
				WebPort:           cfg.CWeb.Port,
				WebMode:           cfg.CWeb.Mode,
				RedisHost:         cfg.CRedis.Host,
				RedisPort:         cfg.CRedis.Port,
				RedisUsername:     cfg.CRedis.Username,
				RedisPassword:     cfg.CRedis.Password,
				RedisDB:           cfg.CRedis.DB,
				EmailSMTPHost:     cfg.CEmail.SMTPHost,
				EmailSMTPPort:     cfg.CEmail.SMTPPort,
				EmailUsername:     cfg.CEmail.Username,
				EmailPassword:     cfg.CEmail.Password,
				EmailFrom:         cfg.CEmail.From,
				LoggerLevel:       cfg.CLogger.Level,
				LoggerLogPath:     cfg.CLogger.LogPath,
				LoggerFormatJson:  cfg.CLogger.FormatJSON,
				LoggerCompress:    cfg.CLogger.Compress,
				LoggerMaxSize:     cfg.CLogger.MaxSize,
				LoggerMaxAge:      cfg.CLogger.MaxAge,
				LoggerMaxBackups:  cfg.CLogger.MaxBackups,
				CaptchaLength:     cfg.CCaptcha.Length,
				CaptchaNoiseCount: cfg.CCaptcha.NoiseCount,
				CaptchaType:       cfg.CCaptcha.Type,
				CaptchaFonts:      cfg.CCaptcha.Fonts,
				CaptchaShowLine:   cfg.CCaptcha.ShowLine,
			},
			Rules: &rules,
			FormItems: []FormItems{
				{Span: 24, Vertical: true, TitleBold: true, Title: "数据库选择", Children: []*Form{
					{Field: "OtherDbType", Title: "数据库类型", Span: 8, ItemRender: ItemRender{Name: vxeSelect, Options: OtherDbType}},
					{Field: "OtherCacheType", Title: "缓存数据库", Span: 8, ItemRender: ItemRender{Name: vxeSelect, Options: OtherCacheType}},
					{Field: "OtherIsEmail", Title: "启用邮件", Span: 8, ItemRender: ItemRender{Name: vxeSelect, Options: OtherIsEmail}},
				}},
				{Span: 24, Vertical: true, TitleBold: true, Title: "Web= 配置", Children: Web},
				{Span: 24, Vertical: true, TitleBold: true, Title: "SQLite 配置", Children: SQLite, ShowWhen: &ShowWhen{Field: "OtherDbType", Value: constants.SQLite}},
				{Span: 24, Vertical: true, TitleBold: true, Title: "MySQL 配置", Children: MySQL, ShowWhen: &ShowWhen{Field: "OtherDbType", Value: constants.MySQL}},
				{Span: 24, Vertical: true, TitleBold: true, Title: "Redis 配置", Children: Redis, ShowWhen: &ShowWhen{Field: "OtherCacheType", Value: constants.Redis}},
				{Span: 24, Vertical: true, TitleBold: true, Title: "Email 配置", Children: Email, ShowWhen: &ShowWhen{Field: "OtherIsEmail", Value: true}},
				{Span: 24, Vertical: true, TitleBold: true, Title: "验证码配置", Children: Captcha},
				{Span: 24, Vertical: true, TitleBold: true, Title: "系统日志配置", Children: Logger},
				{Span: 24, Vertical: true, TitleBold: true, Title: "其他配置", Children: []*Form{{Field: "OtherDataPath", Title: "数据存放路径", Span: 24, ItemRender: ItemRender{Name: "VxeInput"}}}},
				{Span: 24, Children: []*Form{
					{Align: "center", Span: 24, ItemRender: ItemRender{Name: "VxeButtonGroup", Options: []*Options{{Type: "submit", Content: "保存配置", Status: "primary"}}}},
				}},
			},
		}
		c.JsonUnSafeSuccess(form)
	})
}

// postHandler 验证表单，并保存配置
func postHandler(c *gin.Context) {
	with.NewWithContextInitWebData[FormData](c, func(c *with.Context[FormData]) {
		if err := c.ShouldBindJSON(&c.DTOData); err != nil {
			log.Error().Err(err).Msgf("参数格式错误: %+v", err)
			c.JsonUnSafe(code.Unknown, fmt.Sprintf("参数格式错误: %+v", err), nil)
			return
		}

		variable.WebTitle = c.DTOData.WebTitle

		cfg.CSQLite.DbName = c.DTOData.SQLiteDbName

		cfg.CMySQL.Host = c.DTOData.MySQLHost
		cfg.CMySQL.Port = c.DTOData.MySQLPort
		cfg.CMySQL.Username = c.DTOData.MySQLUsername
		cfg.CMySQL.Password = c.DTOData.MySQLPassword
		cfg.CMySQL.DbName = c.DTOData.MySQLDbName

		cfg.CWeb.Host = c.DTOData.WebHost
		cfg.CWeb.Port = c.DTOData.WebPort
		cfg.CWeb.Mode = c.DTOData.WebMode

		cfg.CRedis.Host = c.DTOData.RedisHost
		cfg.CRedis.Port = c.DTOData.RedisPort
		cfg.CRedis.Username = c.DTOData.RedisUsername
		cfg.CRedis.Password = c.DTOData.RedisPassword
		cfg.CRedis.DB = c.DTOData.RedisDB

		cfg.CEmail.SMTPHost = c.DTOData.EmailSMTPHost
		cfg.CEmail.SMTPPort = c.DTOData.EmailSMTPPort
		cfg.CEmail.Username = c.DTOData.EmailUsername
		cfg.CEmail.Password = c.DTOData.EmailPassword
		cfg.CEmail.From = c.DTOData.EmailFrom

		cfg.CLogger.Level = c.DTOData.LoggerLevel
		cfg.CLogger.MaxSize = c.DTOData.LoggerMaxSize
		cfg.CLogger.LogPath = c.DTOData.LoggerLogPath
		cfg.CLogger.FormatJSON = c.DTOData.LoggerFormatJson
		cfg.CLogger.Compress = c.DTOData.LoggerCompress
		cfg.CLogger.MaxAge = c.DTOData.LoggerMaxAge
		cfg.CLogger.MaxBackups = c.DTOData.LoggerMaxBackups

		cfg.CCaptcha.Length = c.DTOData.CaptchaLength
		cfg.CCaptcha.NoiseCount = c.DTOData.CaptchaNoiseCount
		cfg.CCaptcha.Type = c.DTOData.CaptchaType
		cfg.CCaptcha.Fonts = c.DTOData.CaptchaFonts
		cfg.CCaptcha.ShowLine = c.DTOData.CaptchaShowLine

		cfg.COther.IsEmail = c.DTOData.OtherIsEmail
		cfg.COther.DbType = c.DTOData.OtherDbType
		cfg.COther.CacheType = c.DTOData.OtherCacheType
		cfg.COther.DataPath = c.DTOData.OtherDataPath

		initialize.InitializationLogger()

		// 检测数据库链接
		if strings.EqualFold(cfg.COther.DbType, constants.MySQL) {
			db, err := database.CheckDbConnect(cfg.CMySQL.Username, cfg.CMySQL.Password, cfg.CMySQL.Host, cfg.CMySQL.Port)
			if err != nil {
				return
			}
			defer db.Close()
		}

		// 检测缓存链接
		if strings.EqualFold(cfg.COther.CacheType, constants.Redis) {
			if err := utils.CheckRedisConn(cfg.CRedis.Host, cfg.CRedis.Port, cfg.CRedis.Username, cfg.CRedis.Password, cfg.CRedis.DB); err != nil {
				log.Error().Err(err).Msg("Redis 连接失败")
				return
			}
		}

		// 验证是否能够链接邮件服务器
		if c.DTOData.OtherIsEmail {
			if !email.CheckSMTPConnection(cfg.CEmail.SMTPHost, cfg.CEmail.SMTPPort) {
				return
			}

			if !email.CheckSMTPAuth(cfg.CEmail.SMTPHost, cfg.CEmail.SMTPPort, cfg.CEmail.Username, cfg.CEmail.Password) {
				return
			}
		}

		variable.IsInitialized.Store(true)
	})
}
