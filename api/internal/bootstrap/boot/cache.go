package boot

import (
	"gen_gin_tpl/pkg/cache"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/constants"
	"strings"
	"time"
)

func initCache() {
	if strings.EqualFold(constants.Redis, cfg.C.Other.CacheType) {
		cache.RegisterCache(cache.NewRedisCache(cache.RedisConfig{
			Host:            cfg.C.Redis.Host,
			Port:            cfg.C.Redis.Port,
			Username:        cfg.C.Redis.Username,
			Password:        cfg.C.Redis.Password,
			DB:              cfg.C.Redis.DB,
			PoolSize:        50,
			MinIdleConns:    10,
			PoolTimeout:     5 * time.Second,
			MaxRetries:      3,
			MinRetryBackoff: 100 * time.Millisecond,
			MaxRetryBackoff: 1 * time.Second,
			DialTimeout:     5 * time.Second,
			ReadTimeout:     3 * time.Second,
			WriteTimeout:    3 * time.Second,
		}))
	} else {
		cache.RegisterCache(cache.NewMemoryCache())
	}
}
