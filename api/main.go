package main

import (
	"gen_gin_tpl/internal/bootstrap/boot"
	"gen_gin_tpl/internal/bootstrap/initweb"
	"gen_gin_tpl/internal/libs"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/utils"
	"gen_gin_tpl/pkg/utils/cert"
	"gen_gin_tpl/pkg/utils/file"
	"gen_gin_tpl/pkg/variable"
	"github.com/gin-gonic/gin"
	"io"
	"time"

	"github.com/mojocn/base64Captcha"
)

func init() {
	time.Local = time.UTC

	variable.IsEnableEmail.Store(cfg.C.Other.IsEmail)
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
		gin.SetMode(gin.DebugMode)
		go initweb.StartInitializeWeb()
		for {
			if variable.IsInitialized.Load() {
				break
			}
			time.Sleep(5 * time.Second)
		}
		gin.SetMode(cfg.C.Web.Mode)
		libs.InitializeUpdateWebConfig()
	}

	// 禁用 Gin 框架的日志输出
	gin.DefaultWriter = io.Discard
	boot.Initialization().RunTLS(cfg.C.Web.IP, cfg.C.Web.Port, cfg.C.Other.DataPath)
}
