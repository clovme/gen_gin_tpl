package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/logger/log"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var (
	redisOnce          sync.Once
	redisRdbInstance   *redis.Client
	redisCacheInstance *redisCache
)

type RedisConfig struct {
	Host            string        // Redis 主机地址
	Port            int           // Redis 端口
	Username        string        // Redis 用户名
	Password        string        // Redis 密码
	DB              int           // Redis 数据库索引
	PoolSize        int           // Redis 连接池大小
	MinIdleConns    int           // Redis 最小空闲连接数
	PoolTimeout     time.Duration // Redis 连接池超时时间
	MaxRetries      int           // Redis 最大重试次数
	MinRetryBackoff time.Duration // Redis 最小重试间隔
	MaxRetryBackoff time.Duration // Redis 最大重试间隔
	DialTimeout     time.Duration // Redis 连接超时时间
	ReadTimeout     time.Duration // Redis 读取超时时间
	WriteTimeout    time.Duration // Redis 写入超时时间
}

// redisCache 内存缓存项
type redisCache struct {
	name    string
	timeout time.Duration
	client  *redis.Client
}

// NewRedisCache 创建 Redis 缓存实例
// 参数:
//   - host: Redis 主机地址
//   - port: Redis 端口
//   - username: Redis 用户名
//   - password: Redis 密码
//   - db: Redis 数据库索引
//
// 返回值:
//   - *redisCache: Redis 缓存实例指针
func NewRedisCache(cfg RedisConfig) *redisCache {
	redisOnce.Do(func() {
		addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
		redisRdbInstance = redis.NewClient(&redis.Options{
			Addr:     addr,
			Username: cfg.Username,
			Password: cfg.Password,
			DB:       cfg.DB,

			PoolSize:     cfg.PoolSize,
			MinIdleConns: cfg.MinIdleConns,
			PoolTimeout:  cfg.PoolTimeout,

			MaxRetries:      cfg.MaxRetries,
			MinRetryBackoff: cfg.MinRetryBackoff,
			MaxRetryBackoff: cfg.MaxRetryBackoff,

			DialTimeout:  cfg.DialTimeout,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,

			OnConnect: func(ctx context.Context, cn *redis.Conn) error {
				log.Info().Msgf("[Redis] 新连接建立到 %s", addr)
				return nil
			},
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if _, err := redisRdbInstance.Ping(ctx).Result(); err != nil {
			log.Fatal().Err(err).Msg("[Redis] 连接失败，服务终止")
		}

		redisCacheInstance = &redisCache{
			name:    constants.Redis,
			timeout: 5 * time.Second,
			client:  redisRdbInstance,
		}

		// 注册全局 cacheInstance
		RegisterCache(redisCacheInstance)
		log.Info().Msgf("[%s] 缓存初始化完成", redisCacheInstance.name)
	})

	return redisCacheInstance
}

func (r *redisCache) Get(key string) any {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	data, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// key不存在，正常情况，不记录error
		log.Debug().Msgf("[%s] key 不存在：%s", r.name, key)
		return nil
	} else if err != nil {
		log.Error().Err(err).Msgf("[%s] 读取缓存失败！", r.name)
		return nil
	}
	return data
}

func (r *redisCache) GetString(key string) string {
	val, ok := r.Get(key).(string)
	if !ok {
		log.Error().Msgf("[%s] 缓存值类型错误：key=%s, value=%v, type=%T", r.name, key, val, val)
		return ""
	}
	return val
}

func (r *redisCache) Set(key string, value any, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *redisCache) Del(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	return r.client.Del(ctx, key).Err()
}

func (r *redisCache) Has(key string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	exists, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		log.Warn().Err(err).Msgf("[%s] Has 检查失败", r.name)
		return false
	}
	return exists > 0
}

func (r *redisCache) GetDefaultValue(key string, defaultValue any) any {
	val := r.Get(key)
	if val == nil {
		return defaultValue
	}
	return val
}

func (r *redisCache) MustGet(key string) any {
	val := r.Get(key)
	if val == "" {
		log.Warn().Msgf("[%s] 缓存中 key 不存在：%s", r.name, key)
	}
	return val
}

func (r *redisCache) GetJSON(key string, object any) error {
	val := r.Get(key)
	str, ok := val.(string)
	if !ok || str == "" {
		return errors.New(fmt.Sprintf("[%s] 缓存不存在或类型错误", r.name))
	}
	return json.Unmarshal([]byte(str), object)
}

func (r *redisCache) SetJSON(key string, value any, expiration time.Duration) error {
	if value == nil {
		log.Error().Msgf("[%s] 缓存值不能为空", r.name)
		return errors.New(fmt.Sprintf("[%s] 缓存值不能为空", r.name))
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Set(key, string(data), expiration)
}

func (r *redisCache) Increment(key string) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	result, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		log.Error().Err(err).Msgf("[%s] Key:%s 自增失败", r.name, key)
		return int64(0)
	}
	return result
}

func (r *redisCache) Decrement(key string) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	result, err := r.client.Decr(ctx, key).Result()
	if err != nil {
		log.Error().Err(err).Msgf("[%s] Key:%s 自减失败", r.name, key)
		return 0
	}
	return result
}

func (r *redisCache) Context(timeout time.Duration) cacheInterface {
	r.timeout = timeout
	return r
}
