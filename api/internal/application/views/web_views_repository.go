package views

import (
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/internal/schema/dto"
)

type Repository interface {
	// RegeditUser 用户注册
	//
	// 参数:
	//   - regeditDTO: 注册信息
	//   - session: 会话信息
	//
	// 返回值:
	//   - bool: 创建成功返回true，否则返回false
	//   - string: 创建成功返回用户ID，否则返回错误信息
	RegeditUser(regeditDTO dto.RegeditDTO, session core.Session) (bool, string)

	// UserLogin 用户登录
	//
	// 参数:
	//   - loginDTO: 登录信息
	//   - session: 会话信息
	//
	// 返回值:
	//   - bool: 登录成功返回true，否则返回false
	//   - string: 登录成功返回token，否则返回错误信息
	UserLogin(loginDTO dto.LoginDTO, session core.Session) (bool, string)
}
