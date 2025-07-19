package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/utils"
	"github.com/dgraph-io/ristretto"
	"sync"
	"time"
)

var (
	memoryOnce          sync.Once
	memoryCacheInstance *memoryCache
)

type memoryCache struct {
	name  string
	cache *ristretto.Cache
}

func NewMemoryCache() cacheInterface {
	memoryOnce.Do(func() {
		c, _ := ristretto.NewCache(&ristretto.Config{
			NumCounters: 1e7,
			MaxCost:     1 << 30, // 1GB
			BufferItems: 64,
		})
		memoryCacheInstance = &memoryCache{
			name:  constants.Memory,
			cache: c,
		}
	})
	return memoryCacheInstance
}

func (r *memoryCache) Get(key string) any {
	val, found := r.cache.Get(key)
	if !found {
		log.Debug().Msgf("[%s] 读取缓存失败！", r.name)
		return nil
	}
	return val
}

func (r *memoryCache) GetString(key string) string {
	val, ok := r.Get(key).(string)
	if !ok {
		return ""
	}
	return val
}

func (r *memoryCache) Set(key string, value any, expiration time.Duration) error {
	if expiration <= 0 {
		expiration = 100 * 365 * 24 * time.Hour // 100年够用了
	}

	if ok := r.cache.SetWithTTL(key, value, 1, expiration); !ok {
		return errors.New(fmt.Sprintf("[%s] 设置KeyTTL失败：%s", r.name, key))
	}
	return nil
}

func (r *memoryCache) Del(key string) error {
	r.cache.Del(key)
	return nil
}

func (r *memoryCache) Has(key string) bool {
	_, found := r.cache.Get(key)
	return found
}

func (r *memoryCache) GetDefaultValue(key string, defaultValue any) any {
	if val := r.Get(key); val != nil {
		return val
	}
	return defaultValue
}

func (r *memoryCache) MustGet(key string) any {
	val := r.Get(key)
	if val == nil {
		log.Warn().Msgf("[%s] 缓存中 key 不存在：%s", r.name, key)
	}
	return val
}

func (r *memoryCache) GetJSON(key string, object any) error {
	val := r.Get(key)
	if val == nil {
		return errors.New(fmt.Sprintf("[%s] 缓存中 key 不存在：%s", r.name, key))
	}
	bytes, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, object)
}

func (r *memoryCache) SetJSON(key string, value any, expiration time.Duration) error {
	return r.Set(key, value, expiration)
}

func (r *memoryCache) Increment(key string) int64 {
	val := r.Get(key)
	if val == nil {
		r.Set(key, int64(1), 0)
		return int64(1)
	}
	intVal, ok := utils.ToInt64(val)
	if !ok {
		log.Warn().Msgf("[%s] 缓存值类型错误：key=%s, value=%v, type=%T", r.name, key, val, val)
		return -1
	}
	intVal++
	r.Set(key, intVal, 0)
	return intVal
}

func (r *memoryCache) Decrement(key string) int64 {
	val := r.Get(key)
	if val == nil {
		r.Set(key, int64(0), 0)
		return int64(0)
	}
	intVal, ok := utils.ToInt64(val)
	if !ok {
		log.Warn().Msgf("[%s] 缓存值类型错误：key=%s, value=%v, type=%T", r.name, key, val, val)
		return -1
	}
	intVal--
	r.Set(key, intVal, 0)
	return intVal
}

func (r *memoryCache) Context(timeout time.Duration) cacheInterface {
	// 这是是为了兼容Redis的Context方法，这里不做任何操作
	return r
}
