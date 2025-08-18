package core

import (
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/variable"
	"net/http"
)

// HTMLResponse 自定义HTML响应
type viewData struct {
	Data          any
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
//
// 说明:
//   - 加载HTML模板，渲染页面数据。
func (r *Context) HTML(name string, title string, data any) {
	r.Context.HTML(http.StatusOK, name, viewData{
		Data:          data,
		PageTitle:     title,
		WebTitle:      r.Config.GetWebTitle(),
		IsLogin:       r.IsLogin,
		IsEnableEmail: variable.IsEnableEmail.Load(),
		ClientID:      r.Session.BrowserClientID(),
		UserInfo:      r.UserInfo,
	})
}
