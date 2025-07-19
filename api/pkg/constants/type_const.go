// 主要用于定义一些常用的，容易写错的名称常量，比如数据库类型、缓存类型、邮箱类型等。

package constants

import (
	"strings"
)

const (
	MySQL  = "MySQL"
	SQLite = "SQLite"
	Memory = "Memory"
	Redis  = "Redis"
	Email  = "Email"
)

func getConstant(name string) string {
	switch true {
	case strings.EqualFold(name, MySQL):
		return MySQL
	case strings.EqualFold(name, SQLite):
		return SQLite
	case strings.EqualFold(name, Memory):
		return Memory
	case strings.EqualFold(name, Redis):
		return Redis
	case strings.EqualFold(name, Email):
		return Email
	default:
		return ""
	}
}

func GetDbName(name string) string {
	switch true {
	case strings.EqualFold(name, MySQL):
		return MySQL
	default:
		return SQLite
	}
}

func GetCacheName(name string) string {
	switch true {
	case strings.EqualFold(name, Redis):
		return Redis
	default:
		return Memory
	}
}

func GetValue(name string) string {
	return getConstant(name)
}
