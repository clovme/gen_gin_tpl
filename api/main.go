package main

import (
	"fmt"
	"gen_gin_tpl/internal/bootstrap/initialize"
	"gen_gin_tpl/internal/bootstrap/initweb"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/crypto"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils"
	"gen_gin_tpl/pkg/utils/file"
	"gen_gin_tpl/pkg/utils/network"
	"gen_gin_tpl/pkg/variable"
	"gen_gin_tpl/public"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

func init() {
	time.Local = time.UTC
	gin.SetMode(cfg.CWeb.Mode)

	//u_file.RemoveAllData(variable.ConfigPath, true)
	//u_file.RemoveAllData(cfg.COther.DataPath, false)
	//u_file.RemoveAllData(cfg.CLogger.LogPath, false)

	//cfg.SaveToIni()

	variable.IsEnableEmail.Store(cfg.COther.IsEmail)
	variable.IsInitialized.Store(file.IsFileExist(variable.ConfigPath))
	variable.CaptchaStore = base64Captcha.NewMemoryStore(base64Captcha.GCLimitNumber, 2*time.Minute)

	utils.InitSnowflake(1)
	if err := crypto.ParseRsaKeys(public.PublicPEM, public.PrivatePEM); err != nil {
		fmt.Println("密钥初始化失败：", err)
		return
	}
}

func main() {
	exePath, err := file.GetFileAbsPath(".")

	if err != nil {
		log.Error().Err(err).Msg("获取程序所在路径失败")
		return
	}
	// 初始化配置文件
	if !variable.IsInitialized.Load() {
		gin.SetMode(gin.DebugMode)
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
	engine := initialize.Initialization()
	ip := cfg.CWeb.Host
	if ip == "0.0.0.0" {
		ip = network.GetLocalIP(cfg.CWeb.Host)
	}
	for i, route := range engine.Routes() {
		method := fmt.Sprintf("[%s]", route.Method)
		log.Info().Msgf("%03d %-6s http://%s:%d%-30s%-10s%s", i+1, method, ip, cfg.CWeb.Port, route.Path, "-->", route.Name)
	}

	log.Info().Msgf("程序所在路径 %s", exePath)
	if err := engine.Run(fmt.Sprintf("%s:%d", cfg.CWeb.Host, cfg.CWeb.Port)); err != nil {
		log.Error().Err(err).Msg("服务启动失败")
	}
}
