package boot

import (
	"fmt"
	"gen_gin_tpl/internal/bootstrap/database"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils/file"
	"gen_gin_tpl/pkg/variable"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
)

// databaseAutoMigrate 数据库自动迁移
// 参数：
//   - db: 数据库连接
//   - engine: 引擎
//
// 返回值：
//   - 无
//
// 说明：
//   - 数据库自动迁移，根据配置文件判断是否需要迁移
func databaseAutoMigrate(db *gorm.DB, engine *core.Engine) {
	if !file.IsFileExist(variable.ConfigPath) || gin.IsDebugging() || gin.TestMode == gin.Mode() {
		if err := database.AutoMigrate(db, query.Q, engine.Routes()); err != nil {
			log.Error().Err(err).Msg("[初始化]数据库迁移失败！")
			os.Exit(-1)
		}
	}
}

// databaseConnectDB 连接数据库
func databaseConnectDB(dataPath string) *gorm.DB {
	dbPath := filepath.Join(dataPath, cfg.C.SQLite.DbName)
	if !strings.HasSuffix(cfg.C.SQLite.DbName, ".db") {
		dbPath = filepath.Join(dataPath, fmt.Sprintf("%s.db", cfg.C.SQLite.DbName))
	}
	return database.OpenConnectDB(cfg.C.MySQL.Username, cfg.C.MySQL.Password, cfg.C.MySQL.Host, cfg.C.MySQL.Port, cfg.C.MySQL.DbName, cfg.C.Other.DbType, dbPath)
}
