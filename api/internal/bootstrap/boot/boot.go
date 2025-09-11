package boot

import (
	"fmt"
	"gen_gin_tpl/internal/bootstrap/routers"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/internal/libs"
	"gen_gin_tpl/pkg/captcha"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/enums/role"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils/cert"
	"gen_gin_tpl/pkg/utils/file"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// createStaticDir 创建静态目录
func createStaticDir(dataPath string) string {
	static := filepath.Join(dataPath, "static")
	if !file.IsDirExist(static) {
		_ = os.MkdirAll(static, os.ModePerm)
	}
	return static
}

// Initialization 初始化系统
// 返回值：
//   - *gin.Engine 初始化后的Gin引擎
func Initialization() *core.Engine {
	dataPath, err := file.GetFileAbsPath(cfg.C.Other.DataPath)
	if err != nil {
		fmt.Println("获取数据目录失败:", err)
		os.Exit(-1)
	}

	// 创建静态目录
	static := createStaticDir(dataPath)

	// 初始化系统日志
	InitializationLogger(cfg.C.Logger)

	// 初始化验证码
	captcha.InitImageCaptcha(cfg.C.Captcha.Length, cfg.C.Captcha.NoiseCount, cfg.C.Captcha.ShowLine, cfg.C.Captcha.Fonts, cfg.C.Captcha.Type)

	// 初始化缓存
	initCache()

	// 初始化表单验证器
	initFormValidate()

	// 连接数据库
	db := databaseConnectDB(dataPath)

	// 生成证书
	cert.GenCertificateFile(cfg.C.Web.IP, cfg.C.Web.IP, dataPath)

	// 初始化配置
	if !cfg.C.Other.IsInitialize {
		// 初始化路由
		engine := routers.Initialization(db, static)
		// 数据库自动迁移
		databaseAutoMigrate(db, engine)
		// 初始化标志
		cfg.C.Other.IsInitialize = true
		cfg.SaveToIni()

		admin := strings.ToLower(role.Admin.Key())
		system := strings.ToLower(role.System.Key())
		log.Info().Msgf("初始账号为：%s，密码为：%s", admin, admin)
		log.Info().Msgf("初始账号为：%s，密码为：%s", system, system)
		log.Info().Msgf("初始化程序初始化完成，3秒后程序将退出，请重新启动程序进入正式模式！")
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}
	// 初始化系统配置
	libs.WebConfig.UpdateWebConfig()
	// 初始化路由
	return routers.Initialization(db, static)
}
