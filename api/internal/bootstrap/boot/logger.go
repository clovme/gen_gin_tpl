package boot

import (
	"fmt"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/logger"
	"gen_gin_tpl/pkg/utils/file"
	"os"
)

// InitializationLogger 初始化日志
func InitializationLogger(c cfg.Logger) {
	path, err := file.GetFileAbsPath(c.LogPath)
	if err != nil {
		fmt.Println("获取日志目录失败:", err)
		os.Exit(-1)
	}
	// 初始化一次
	logger.InitLogger(logger.LoggerConfig{
		Dir:        path,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
		Compress:   c.Compress,
		Level:      c.Level,
		FormatJSON: c.FormatJSON, // true=结构化；false=文本
	})
}
