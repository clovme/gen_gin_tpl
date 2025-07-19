package cache

import (
	"time"
)

var (
	cacheInstance cacheInterface
)

// cacheInterface  缓存接口定义
type cacheInterface interface {
	// Get 根据 key 获取缓存值
	// 参数:
	//   - key: 缓存键
	// 返回值:
	//   - string: 缓存值
	Get(key string) any

	// GetString 根据 key 获取缓存值
	// 参数:
	//   - key: 缓存键
	// 返回值:
	//   - string: 缓存值
	GetString(key string) string

	// Set 设置缓存值，expiration 为过期时间
	// 参数:
	//   - key: 缓存键
	//   - value: 缓存值
	//   - expiration: 过期时间，单位：秒
	// 返回值:
	//   - error: 错误信息
	Set(key string, value any, expiration time.Duration) error

	// Del 删除缓存
	// 参数:
	//   - key: 缓存键
	// 返回值:
	//   - error: 错误信息
	Del(key string) error

	// Has 判断缓存是否存在
	// 参数:
	//   - key: 缓存键
	// 返回值:
	//   - bool: 缓存是否存在，若存在且未过期则返回 true，否则返回 false
	Has(key string) bool

	// Context 设置缓存超时时间，缓存初始使用入口
	// 参数:
	//   - timeout: 超时时间，单位：秒
	// 返回值:
	//   - cacheInterface: 缓存实例
	Context(timeout time.Duration) cacheInterface

	// GetDefaultValue 获取缓存值，若不存在则返回默认值
	// 参数:
	//   - key: 缓存键
	//   - defaultValue: 默认值
	// 返回值:
	//   - string: 缓存值
	GetDefaultValue(key string, defaultValue any) any

	// MustGet 获取缓存值，若不存在则抛出异常
	// 参数:
	//   - key: 缓存键
	// 返回值:
	//   - string: 缓存值
	// 异常:
	//   - 缓存不存在时抛出异常
	MustGet(key string) any

	// GetJSON 获取缓存值并转换为对象类型
	// 参数:
	//   - key: 缓存键
	//   - object: 目标结构体指针
	// 返回值:
	//   - error: 错误信息
	GetJSON(key string, object any) error

	// SetJSON 设置缓存值为 JSON 类型
	// 参数:
	//   - key: 缓存键
	//   - value: 缓存值
	//   - expiration: 过期时间，单位：秒
	// 返回值:
	//   - error: 错误信息
	SetJSON(key string, value any, expiration time.Duration) error

	// Increment 自增缓存值
	// 参数:
	//   - key: 缓存键
	// 返回值:
	//   - int64: 自增后的值
	// 异常:
	//   - 缓存不存在时抛出异常
	Increment(key string) int64

	// Decrement 自减缓存值
	// 参数:
	//   - key: 缓存键
	// 返回值:
	//   - int64: 自减后的值
	// 异常:
	//   - 缓存不存在时抛出异常
	Decrement(key string) int64
}
