package page

import (
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/session"
	"gen_gin_tpl/pkg/variable"
	"github.com/gin-gonic/gin"
)

// ViewRoot 视图数据结构
type ViewRoot[T any] struct {
	Data          T
	IsEnableEmail bool
	IsLogin       bool
	WebTitle      string
	PageTitle     string
	ClientID      string
	Errors        map[string]string
}

// ViewData 模版数据结构
// 参数：
//   - title string 页面标题
//   - data T 数据
//   - err map[string]string 错误信息
//
// 返回值：
//   - ViewRoot[T] 视图数据结构
func ViewData[T any](c *gin.Context, title string, data T, err map[string]string) ViewRoot[T] {
	return ViewRoot[T]{
		Errors:        err,
		Data:          data,
		PageTitle:     title,
		WebTitle:      variable.WebTitle,
		IsLogin:       session.IsLogin(c),
		IsEnableEmail: variable.IsEnableEmail.Load(),
		ClientID:      session.Get(c, constants.ClientID).(string),
	}
}

// ViewDataNil 模版数据结构
// 参数：
//   - title string 页面标题
//
// 返回值：
//   - ViewRoot[struct{}] 视图数据结构
func ViewDataNil(c *gin.Context, title string) ViewRoot[struct{}] {
	return ViewRoot[struct{}]{
		PageTitle:     title,
		WebTitle:      variable.WebTitle,
		IsLogin:       session.IsLogin(c),
		IsEnableEmail: variable.IsEnableEmail.Load(),
		ClientID:      session.Get(c, constants.ClientID).(string),
	}
}
