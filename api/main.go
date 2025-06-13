package main

import (
	"fmt"
	"gen_gin_tpl/internal/bootstrap/database"
	"gen_gin_tpl/internal/bootstrap/initialize"
	"gen_gin_tpl/internal/bootstrap/routers"
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/pkg/config"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/crypto"
	"gen_gin_tpl/pkg/initweb"
	"gen_gin_tpl/pkg/let"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils"
	"gen_gin_tpl/public"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strings"
	"time"
)

var cfg *config.Config

func init() {
	time.Local = time.UTC
	cfg = config.GetConfig()
	gin.SetMode(gin.ReleaseMode)

	let.IsInitialized.Store(utils.IsFileExist(let.ConfigPath))

	utils.InitSnowflake(1)
	if err := crypto.ParseRsaKeys(public.PublicPEM, public.PrivatePEM); err != nil {
		fmt.Println("密钥初始化失败：", err)
		return
	}
	let.SQLitePath = filepath.Join(cfg.Other.DataPath, fmt.Sprintf("%s.db", constants.ProjectName))
}

func main() {
	// 初始化配置文件
	if !let.IsInitialized.Load() {
		initweb.Initialization(cfg)
		for {
			if let.IsInitialized.Load() {
				break
			}
		}
		config.SaveConfig()
	}

	initialize.InitLogger(cfg.Logger)
	initialize.InitCache(*cfg)
	db := database.OpenConnectDB(*cfg)
	engine := routers.Initialization(db)

	if err := database.AutoMigrate(db, query.Q, engine.Routes()); err != nil {
		log.Error().Err(err).Msg("[初始化]数据库迁移失败！")
		return
	}
	initialize.InitSystemConfig(query.Q)

	ipPort := fmt.Sprintf("%s:%d", cfg.Web.Host, cfg.Web.Port)
	for i, route := range engine.Routes() {
		if strings.HasSuffix(route.Path, "*filepath") {
			continue
		}
		log.Info().Msgf("%03d [%s] http://%s%-30s%-10s%s", i+1, route.Method, ipPort, route.Path, "->", route.Handler)
	}
	engine.Run(fmt.Sprintf("%s:%d", cfg.Web.Host, cfg.Web.Port))
}
