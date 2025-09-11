package database

import (
	"database/sql"
	"fmt"
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/logger"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils/file"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// CheckDbConnect 检查数据库连接是否正常, 只连server，不带库名
// 参数：
//   - username MySQL用户名
//   - password MySQL密码
//   - host MySQL主机
//   - port MySQL端口
//
// 返回值：
//   - *sql.DB 数据库连接对象
func CheckDbConnect(username, password, host string, port int) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, host, port))
	if err != nil {
		log.Error().Err(err).Msg("[数据库检测] 数据库连接失败，请检查MySQL服务是否启动，或者IP或者端口号")
		os.Exit(-1)
	}

	if err = db.Ping(); err != nil {
		log.Error().Err(err).Msg("[数据库检测] 无法建立数据库连接，请检查MySQL服务是否启动，或者IP或者端口号")
		os.Exit(-1)
	}

	// 查询数据库版本
	var version string
	if err := db.QueryRow("SELECT VERSION()").Scan(&version); err != nil {
		log.Warn().Err(err).Msg("[数据库检测] 查询数据库版本失败")
	} else {
		log.Info().Msgf("[数据库检测] 数据库连接成功，版本: %s", version)
	}

	return db, nil
}

// MySQL建库
// 参数：
//   - username MySQL用户名
//   - password MySQL密码
//   - host MySQL主机
//   - port MySQL端口
//   - dbName 数据库名
//
// 返回值：
//   - bool 是否创建成功
func checkAndCreateDatabase(username, password, host string, port int, dbName string) bool {
	// 只连server，不带库名
	db, err := CheckDbConnect(username, password, host, port)
	if err != nil {
		return false
	}
	defer db.Close()

	var count int
	if err = db.QueryRow("SELECT COUNT(*) FROM information_schema.schemata WHERE schema_name = ?", dbName).Scan(&count); err != nil {
		log.Error().Err(err).Msgf("[数据库初始化] 数据库[%s]查询失败...", dbName)
		return false
	}

	if count > 0 {
		return true
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE `%s` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';", dbName))
	if err != nil {
		log.Panic().Err(err).Msgf("[数据库初始化] 创建数据库[%s]失败", dbName)
		return false
	}
	log.Info().Msgf("[数据库初始化] 数据库[%s]初始化完成...", dbName)
	return true
}

// OpenConnectDB 打开数据库连接, 并设置连接池, 数据库链接统一入口
// 参数：
//   - username MySQL用户名
//   - password MySQL密码
//   - host MySQL主机
//   - port MySQL端口
//   - dbName 数据库名
//
// 返回值：
//   - *gorm.DB 数据库连接对象
func OpenConnectDB(username, password, host string, port int, dbName, dbType, sQLitePath string) *gorm.DB {
	var dsn gorm.Dialector

	if strings.EqualFold(constants.SQLite, dbType) {
		if !file.IsDirExist(filepath.Dir(sQLitePath)) {
			_ = os.MkdirAll(filepath.Dir(sQLitePath), os.ModePerm)
		}
		//dsn = sqlite.Open(fmt.Sprintf("file:%s?mode=&cache=shared&_foreign_keys=1", sQLitePath))
		dsn = sqlite.Open(sQLitePath)
	} else {
		if !checkAndCreateDatabase(username, password, host, port, dbName) { // 先检查并建库
			log.Error().Msg("[数据库初始化] 数据库初始化失败")
			os.Exit(-1)
		}
		dsn = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai", username, password, host, port, dbName))
	}

	db, err := gorm.Open(dsn, logger.GetGormLogger())
	if err != nil {
		log.Error().Err(err).Msg("打开SQLite失败")
		os.Exit(-1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("获取底层 sql.DB 失败")
		os.Exit(-1)
	}
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	query.SetDefault(db)
	return db
}
