package initialize

import (
	"fmt"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/logger"
	"gen_gin_tpl/pkg/utils/u_file"
	"os"
)

// InitializationLogger 初始化日志
func InitializationLogger() {
	path, err := u_file.GetFileAbsPath(cfg.CLogger.LogPath, "")
	if err != nil {
		fmt.Println("获取日志目录失败:", err)
		os.Exit(-1)
	}
	// 初始化一次
	logger.InitLogger(logger.LoggerConfig{
		Dir:        path,
		MaxSize:    cfg.CLogger.MaxSize,
		MaxBackups: cfg.CLogger.MaxBackups,
		MaxAge:     cfg.CLogger.MaxAge,
		Compress:   cfg.CLogger.Compress,
		Level:      cfg.CLogger.Level,
		FormatJSON: cfg.CLogger.FormatJSON, // true=结构化；false=文本
	})
}
