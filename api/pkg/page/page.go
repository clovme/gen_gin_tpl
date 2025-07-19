package page

import (
	"gen_gin_tpl/pkg/cfg"
	"gen_gin_tpl/pkg/variable"
)

// ViewRoot 视图数据结构
type ViewRoot[T any] struct {
	WebTitle  string
	PageTitle string
	IsEmail   bool
	Data      T
	Errors    map[string]string
}

// ViewData 模版数据结构
// 参数：
//   - title string 页面标题
//   - data T 数据
//   - err map[string]string 错误信息
//
// 返回值：
//   - ViewRoot[T] 视图数据结构
func ViewData[T any](title string, data T, err map[string]string) ViewRoot[T] {
	return ViewRoot[T]{
		WebTitle:  variable.WebTitle,
		PageTitle: title,
		IsEmail:   cfg.COther.IsEmail,
		Data:      data,
		Errors:    err,
	}
}

// ViewDataNil 模版数据结构
// 参数：
//   - title string 页面标题
//
// 返回值：
//   - ViewRoot[struct{}] 视图数据结构
func ViewDataNil(title string) ViewRoot[struct{}] {
	return ViewRoot[struct{}]{
		WebTitle:  variable.WebTitle,
		IsEmail:   cfg.COther.IsEmail,
		PageTitle: title,
	}
}
