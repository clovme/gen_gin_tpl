package utils

import (
	"unicode"
)

// Capitalize 首字符大写
// 参数：
//   - s 字符串
//
// 返回值：
//   - string 首字符大写后的字符串
func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	// 转成 rune 切片，防止中文/多字节字符乱码
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
