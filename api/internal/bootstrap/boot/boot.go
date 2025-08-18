package boot

import (
	"fmt"
	"gen_gin_tpl/internal/bootstrap/routers"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/internal/libs"
	"gen_gin_tpl/pkg/captcha"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/utils/cert"
	"gen_gin_tpl/pkg/utils/file"
	"gen_gin_tpl/pkg/variable"
	"os"
	"path/filepath"
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
	dataPath, err := file.GetFileAbsPath(cfg.COther.DataPath)
	if err != nil {
		fmt.Println("获取数据目录失败:", err)
		os.Exit(-1)
	}

	// 创建静态目录
	static := createStaticDir(dataPath)

	// 初始化系统日志
	InitializationLogger()

	// 初始化验证码
	captcha.InitImageCaptcha(cfg.CCaptcha.Length, cfg.CCaptcha.NoiseCount, cfg.CCaptcha.ShowLine, cfg.CCaptcha.Fonts, cfg.CCaptcha.Type)

	// 初始化缓存
	initCache()

	// 初始化表单验证器
	initFormValidate()

	// 连接数据库
	db := databaseConnectDB(dataPath)

	// 初始化配置
	if !file.IsFileExist(variable.ConfigPath) || !cfg.COther.IsInitialize {
		// 生成证书
		cert.GenCertificateFile(dataPath)
		// 初始化路由
		engine := routers.Initialization(db, static)
		// 数据库自动迁移
		databaseAutoMigrate(db, engine)
		// 初始化系统配置
		libs.WebConfig.Update()
		// 初始化标志
		cfg.COther.IsInitialize = true
		// 保存配置
		cfg.SaveToIni()
		return engine
	}
	// 初始化系统配置
	libs.WebConfig.Update()
	// 初始化路由
	return routers.Initialization(db, static)
}
