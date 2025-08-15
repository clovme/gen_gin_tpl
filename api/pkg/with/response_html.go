package with

import (
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/variable"
	"net/http"
)

type viewData[T any] struct {
	Data          *T
	IsEnableEmail bool
	IsLogin       bool
	WebTitle      string
	PageTitle     string
	ClientID      string
	UserInfo      *models.User
}

// HTML 加载HTML模板
//
// 参数:
//   - name: 模板名称
//   - title: 页面标题
//   - data: 页面数据
//   - err: 错误信息
func (r *Context[T]) HTML(name string, title string) {
	r.Context.HTML(http.StatusOK, name, viewData[T]{
		Data:          &r.DTOData,
		PageTitle:     title,
		WebTitle:      variable.WebTitle,
		IsLogin:       r.IsLogin,
		IsEnableEmail: variable.IsEnableEmail.Load(),
		ClientID:      r.Session.BrowserClientID(),
		UserInfo:      getUserInfo(r.Context),
	})
}
