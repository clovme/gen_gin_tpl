package views

import (
	"gen_gin_tpl/internal/schema/dto"
	"github.com/gin-gonic/gin"
)

type WebViewsService struct {
	Repo Repository
}

// CreateUser 创建用户
//
// 参数:
//   - c: gin.Context
//   - regeditDTO: dto.RegeditDTO
//
// 返回值:
//   - bool: 创建成功返回true，否则返回false
//   - string: 创建成功返回用户ID，否则返回错误信息
func (r *WebViewsService) CreateUser(c *gin.Context, regeditDTO dto.RegeditDTO) (bool, string) {
	return r.Repo.CreateUser(c, regeditDTO)
}

// UserLogin 用户登录
//
// 参数:
//   - c: gin.Context
//   - loginDTO: dto.LoginDTO
//
// 返回值:
//   - bool: 登录成功返回true，否则返回false
//   - string: 登录成功返回token，否则返回错误信息
func (r *WebViewsService) UserLogin(c *gin.Context, loginDTO dto.LoginDTO) (bool, string) {
	return r.Repo.UserLogin(c, loginDTO)
}
