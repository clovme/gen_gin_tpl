package views

import (
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/internal/schema/dto"
)

type WebViewsService struct {
	Repo Repository
}

// CreateUser 创建用户
//
// 参数:
//   - regeditDTO: dto数据
//   - session: 会话信息
//
// 返回值:
//   - bool: 创建成功返回true，否则返回false
//   - string: 创建成功返回用户ID，否则返回错误信息
func (r *WebViewsService) CreateUser(regeditDTO dto.RegeditDTO, session core.Session) (bool, string) {
	return r.Repo.RegeditUser(regeditDTO, session)
}

// UserLogin 用户登录
//
// 参数:
//   - loginDTO: dto数据
//   - session: 会话信息
//
// 返回值:
//   - bool: 登录成功返回true，否则返回false
//   - string: 登录成功返回token，否则返回错误信息
func (r *WebViewsService) UserLogin(loginDTO dto.LoginDTO, session core.Session) (bool, string) {
	return r.Repo.UserLogin(loginDTO, session)
}
