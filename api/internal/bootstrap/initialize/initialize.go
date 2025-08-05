package initialize

import (
	"fmt"
	"gen_gin_tpl/internal/bootstrap/database"
	"gen_gin_tpl/internal/bootstrap/routers"
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/pkg/captcha"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils/file"
	"gen_gin_tpl/pkg/variable"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)

// Initialization 初始化系统
// 返回值：
//   - *gin.Engine 初始化后的Gin引擎
func Initialization() *gin.Engine {
	dataPath, err := file.GetFileAbsPath(cfg.COther.DataPath, "")
	if err != nil {
		fmt.Println("获取数据目录失败:", err)
		os.Exit(-1)
	}

	// 初始化系统日志
	InitializationLogger()

	// 初始化验证码
	captcha.InitImageCaptcha(cfg.CCaptcha.Length, cfg.CCaptcha.NoiseCount, cfg.CCaptcha.ShowLine, cfg.CCaptcha.Fonts, cfg.CCaptcha.Type)

	// 初始化缓存
	initCache()

	// 初始化表单验证器
	initFormValidate()

	dbPath := filepath.Join(dataPath, cfg.CSQLite.DbName)
	if !strings.HasSuffix(cfg.CSQLite.DbName, ".db") {
		dbPath = filepath.Join(dataPath, fmt.Sprintf("%s.db", cfg.CSQLite.DbName))
	}
	db := database.OpenConnectDB(cfg.CMySQL.Username, cfg.CMySQL.Password, cfg.CMySQL.Host, cfg.CMySQL.Port, cfg.CMySQL.DbName, cfg.COther.DbType, dbPath)

	static := filepath.Join(dataPath, "static")
	if !file.IsDirExist(static) {
		_ = os.MkdirAll(static, os.ModePerm)
	}
	engine := routers.Initialization(db, static)

	if !file.IsFileExist(variable.ConfigPath) {
		if err := database.AutoMigrate(db, query.Q, engine.Routes()); err != nil {
			log.Error().Err(err).Msg("[初始化]数据库迁移失败！")
			return nil
		}
		cfg.SaveToIni()
	}

	initializationSystemConfig(query.Q)

	return engine
}
