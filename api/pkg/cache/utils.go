package cache

import (
	"gen_gin_tpl/pkg/logger/log"
	"github.com/redis/go-redis/v9"
)

// RegisterCache 注册缓存实例
// 参数:
//   - c: 缓存实例
func RegisterCache(c cacheInterface) {
	cacheInstance = c
}

// PoolStats 获取 Redis 连接池状态
//
// 返回值:
//   - *redis.PoolStats: Redis 连接池状态
func PoolStats() *redis.PoolStats {
	if redisCacheInstance == nil {
		log.Error().Msg("未注册 Redis 缓存实例")
		return nil
	}
	stats := redisCacheInstance.client.PoolStats()
	log.Debug().Msgf("当前活跃连接数: %d, 空闲连接数: %d", stats.TotalConns, stats.IdleConns)
	return stats
}

// CheckIsCache 检查是否已注册缓存实例
func CheckIsCache() bool {
	return cacheInstance != nil
}
