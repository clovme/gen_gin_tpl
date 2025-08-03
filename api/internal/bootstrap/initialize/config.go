package initialize

import (
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/pkg/cache"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/enums/boolean"
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/variable"
	"gen_gin_tpl/public"
)

// 设置配置
// 参数：
//   - key: 配置键
//   - value: 配置值
//   - defaultValue: 默认值
//   - s: 状态，启用或禁用
//
// 返回值：
//   - 无
func _setConfig[T comparable](key *T, value, defaultValue T, s status.Status) {
	if !s.Is(status.Enable) {
		*key = defaultValue
		return
	}
	*key = value
}

func _setConfigCache(key string, value, defaultValue any, s status.Status) {
	if !s.Is(status.Enable) {
		cache.Set(key, defaultValue, 0)
		return
	}
	cache.Set(key, value, 0)
}

// 设置配置
// 参数：
//   - key: 配置键
//   - value: 配置值
//   - defaultValue: 默认值
//   - s: 状态，启用或禁用
//
// 返回值：
//   - 无
func _setByteConfig(key *[]byte, value, defaultValue []byte, s status.Status) {
	if !s.Is(status.Enable) {
		*key = defaultValue
		return
	}
	*key = value
}

// 初始化系统配置
// 参数：
//   - query: 查询对象
//
// 返回值：
//   - 无
func initializationSystemConfig(query *query.Query) {
	configs, err := query.Config.Find()
	if err != nil {
		return
	}

	for _, cfg := range configs {
		switch cfg.Name {
		case constants.Countdown: // 统一倒计时时间，单位秒
			_setConfigCache(constants.Countdown, cfg.Value, cfg.Default, cfg.Status)
		case constants.ContextIsEncrypted: // 是否开启加密模式
			_setConfig(&variable.IsEnableEncryptedKey, cfg.Value == boolean.True.Key(), cfg.Default == boolean.True.Key(), cfg.Status)
		case constants.WebTitle: // 网站标题
			_setConfig(&variable.WebTitle, cfg.Value, cfg.Default, cfg.Status)
		case constants.PublicPEM: // 公钥
			_setByteConfig(&public.PublicPEM, []byte(cfg.Value), []byte(cfg.Default), cfg.Status)
		case constants.PrivatePEM: // 私钥
			_setByteConfig(&public.PrivatePEM, []byte(cfg.Value), []byte(cfg.Default), cfg.Status)
		case constants.ProjectName: // 项目名称
			variable.WebTitle = cfg.Value
		}
	}
}
