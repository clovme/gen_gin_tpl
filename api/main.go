package main

import (
	"gen_gin_tpl/internal/bootstrap/boot"
	"gen_gin_tpl/internal/bootstrap/initweb"
	"gen_gin_tpl/internal/libs"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils"
	"gen_gin_tpl/pkg/utils/cert"
	"gen_gin_tpl/pkg/utils/file"
	"gen_gin_tpl/pkg/variable"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

func init() {
	time.Local = time.UTC

	variable.IsEnableEmail.Store(cfg.COther.IsEmail)
	variable.IsInitialized.Store(file.IsFileExist(variable.ConfigPath))
	variable.CaptchaStore = base64Captcha.NewMemoryStore(base64Captcha.GCLimitNumber, 2*time.Minute)

	utils.InitSnowflake(1)
	// 生成Rsa密钥
	cert.InitRSAVariable()
	// 初始化系统配置
	libs.InitializeWebConfig()
}

func main() {
	// 初始化配置文件
	if !variable.IsInitialized.Load() {
		gin.SetMode(gin.ReleaseMode)
		go initweb.StartInitializeWeb()
		go initweb.StopInitializeWeb()
		for {
			if variable.IsInitialized.Load() {
				break
			}
			time.Sleep(5 * time.Second)
		}
		gin.SetMode(cfg.CWeb.Mode)
	}

	// 禁用 Gin 框架的日志输出
	gin.DefaultWriter = io.Discard
	engine := boot.Initialization()

	if err := engine.RunTLS(cfg.CWeb.Host, cfg.CWeb.Port, cfg.COther.DataPath); err != nil {
		log.Error().Err(err).Msg("服务启动失败")
	}
}
