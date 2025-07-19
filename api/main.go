package main

import (
	"fmt"
	"gen_gin_tpl/internal/bootstrap/initialize"
	"gen_gin_tpl/internal/bootstrap/initweb"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/crypto"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils"
	"gen_gin_tpl/pkg/utils/u_file"
	"gen_gin_tpl/pkg/utils/u_network"
	"gen_gin_tpl/pkg/variable"
	"gen_gin_tpl/public"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"strings"
	"time"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	time.Local = time.UTC

	//u_file.RemoveAllData(variable.ConfigPath, true)
	//u_file.RemoveAllData(cfg.COther.DataPath, false)
	//u_file.RemoveAllData(cfg.CLogger.LogPath, false)

	//cfg.SaveToIni()

	variable.CaptchaStore = base64Captcha.NewMemoryStore(base64Captcha.GCLimitNumber, 2*time.Minute)

	variable.IsInitialized.Store(u_file.IsFileExist(variable.ConfigPath))

	utils.InitSnowflake(1)
	if err := crypto.ParseRsaKeys(public.PublicPEM, public.PrivatePEM); err != nil {
		fmt.Println("密钥初始化失败：", err)
		return
	}
}

func main() {
	exePath, err := u_file.GetFileAbsPath(".", "")

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
		gin.SetMode(gin.ReleaseMode)
	}

	engine := initialize.Initialization()
	ip := cfg.CWeb.Host
	if ip == "0.0.0.0" {
		ip = u_network.GetLocalIP(cfg.CWeb.Host)
	}
	for i, route := range engine.Routes() {
		if strings.HasSuffix(route.Path, "*filepath") {
			continue
		}
		method := fmt.Sprintf("[%s]", route.Method)
		log.Info().Msgf("%03d %-6s http://%s:%d%-30s%-10s%s", i+1, method, ip, cfg.CWeb.Port, route.Path, "->", route.Handler)
	}

	log.Info().Msgf("程序所在路径 %s", exePath)
	if err := engine.Run(fmt.Sprintf("%s:%d", cfg.CWeb.Host, cfg.CWeb.Port)); err != nil {
		log.Error().Err(err).Msg("服务启动失败")
	}
}
