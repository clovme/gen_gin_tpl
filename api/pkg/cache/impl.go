package cache

import (
	"time"
)

// Get 根据 key 获取缓存值
// 参数:
//   - key: 缓存键
//
// 返回值:
//   - T: 缓存值
func Get[T any](key string) T {
	val := Context(5).Get(key)
	if v, ok := val.(T); ok {
		return v
	}
	var zero T
	return zero
}

// GetString 根据 key 获取缓存值
// 参数:
//   - key: 缓存键
//
// 返回值:
//   - string: 缓存值
func GetString(key string) string {
	return Context(5).GetString(key)
}

// Set 设置缓存值，expiration 为过期时间
// 参数:
//   - key: 缓存键
//   - value: 缓存值
//   - expiration: 过期时间，单位：秒
//
// 返回值:
//   - error: 错误信息
func Set(key string, value any, expiration time.Duration) {
	_ = Context(5).Set(key, value, expiration)
}

// Del 删除缓存
// 参数:
//   - key: 缓存键
//
// 返回值:
//   - error: 错误信息
func Del(key string) error {
	return Context(5).Del(key)
}

// Has 判断缓存是否存在
// 参数:
//   - key: 缓存键
//
// 返回值:
//   - bool: 缓存是否存在，若存在且未过期则返回 true，否则返回 false
func Has(key string) bool {
	return Context(5).Has(key)
}

// Context 设置缓存超时时间
// 参数:
//   - timeout: 超时时间，单位：秒
//
// 返回值:
//   - cacheInterface: 缓存实例，缓存初始使用入口
func Context(timeout time.Duration) cacheInterface {
	return cacheInstance.Context(timeout)
}

// GetDefaultValue 获取缓存值，若不存在则返回默认值
// 参数:
//   - key: 缓存键
//   - defaultValue: 默认值
//
// 返回值:
//   - string: 缓存值
func GetDefaultValue(key string, defaultValue any) any {
	return Context(5).GetDefaultValue(key, defaultValue)
}

// MustGet 获取缓存值，若不存在则抛出异常
// 参数:
//   - key: 缓存键
//
// 返回值:
//   - string: 缓存值
//
// 异常:
//   - 缓存不存在时抛出异常
func MustGet(key string) any {
	return Context(5).MustGet(key)
}

// GetJSON 获取缓存值并转换为对象类型
// 参数:
//   - key: 缓存键
//   - target: 目标结构体指针
//
// 返回值:
//   - error: 错误信息
func GetJSON[T comparable](key string, object T) error {
	return Context(5).GetJSON(key, object)
}

// SetJSON 设置缓存值为 JSON 类型
// 参数:
//   - key: 缓存键
//   - value: 缓存值
//   - expiration: 过期时间，单位：秒
//
// 返回值:
//   - error: 错误信息
func SetJSON(key string, value any, expiration time.Duration) error {
	return Context(5).SetJSON(key, value, expiration)
}

// Increment 自增缓存值
// 参数:
//   - key: 缓存键
//
// 返回值:
//   - int64: 自增后的值
//
// 异常:
//   - 缓存不存在时抛出异常
func Increment(key string) int64 {
	return Context(5).Increment(key)
}

// Decrement 自减缓存值
// 参数:
//   - key: 缓存键
//
// 返回值:
//   - int64: 自减后的值
//
// 异常:
//   - 缓存不存在时抛出异常
func Decrement(key string) int64 {
	return Context(5).Decrement(key)
}
