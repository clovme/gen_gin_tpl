package persistence

import (
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/internal/schema/dto"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/pwd"
	"gen_gin_tpl/pkg/session"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

type WebViewsRepository struct {
	DB *gorm.DB
	Q  *query.Query
}

// CreateUser 创建用户
//
// 参数:
//   - c: gin.Context
//   - regeditDTO: dto.RegeditDTO
//
// 返回值:
//   - bool: 创建成功返回true，否则返回false
//   - string: 返回信息
func (r *WebViewsRepository) CreateUser(c *gin.Context, regeditDTO dto.RegeditDTO) (bool, string) {
	user := &models.User{
		Username: regeditDTO.Username,
		Email:    regeditDTO.Email,
		Password: regeditDTO.Password,
		Nickname: regeditDTO.Username,
	}

	if err := r.Q.User.Create(user); err != nil {
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
			if strings.HasSuffix(err.Error(), "users.email") {
				return false, "邮箱已存在，请重试！"
			} else if strings.HasSuffix(err.Error(), "users.username") {
				return false, "用户名已存在，请重试！"
			} else if strings.HasSuffix(err.Error(), "users.phone") {
				return false, "手机号已存在，请重试！"
			}
		}
		log.Error().Err(err).Msg("用户注册失败")
		return false, "用户注册失败！"
	}
	session.Set(c, constants.UserSessionID, user.ID)
	return true, "用户注册成功！"
}

// UserLogin 用户登录
//
// 参数:
//   - c: gin.Context
//   - loginDTO: dto.LoginDTO
//
// 返回值:
//   - bool: 登录成功返回true，否则返回false
//   - string: 返回信息
func (r *WebViewsRepository) UserLogin(c *gin.Context, loginDTO dto.LoginDTO) (bool, string) {
	user, err := r.Q.User.Where(r.Q.User.Where(r.Q.User.Username.Eq(loginDTO.Username)).Or(r.Q.User.Email.Eq(loginDTO.Username))).Where(r.Q.User.Password.Eq(pwd.Encryption(loginDTO.Password))).First()
	if err != nil {
		return false, "用户名或密码错误！"
	}
	session.Set(c, constants.UserSessionID, user.ID)
	return true, "用户登录成功！"
}
