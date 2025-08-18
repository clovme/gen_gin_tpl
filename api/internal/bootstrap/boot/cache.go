package boot

import (
	"gen_gin_tpl/pkg/cache"
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/constants"
	"strings"
	"time"
)

func initCache() {
	if strings.EqualFold(constants.Redis, cfg.COther.CacheType) {
		cache.RegisterCache(cache.NewRedisCache(cache.RedisConfig{
			Host:            cfg.CRedis.Host,
			Port:            cfg.CRedis.Port,
			Username:        cfg.CRedis.Username,
			Password:        cfg.CRedis.Password,
			DB:              cfg.CRedis.DB,
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
