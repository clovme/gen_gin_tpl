package persistence

import (
	"fmt"
	"gen_gin_tpl/internal/core"
	"gen_gin_tpl/internal/infrastructure/query"
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/internal/schema/dto"
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/logger/log"
	"gen_gin_tpl/pkg/pwd"
	"gorm.io/gorm"
	"strings"
)

type WebViewsRepository struct {
	DB *gorm.DB
	Q  *query.Query
}

// RegeditUser 创建用户
//
// 参数:
//   - regeditDTO: dto数据
//   - session: 会话信息
//
// 返回值:
//   - bool: 创建成功返回true，否则返回false
//   - string: 创建成功返回用户ID，否则返回错误信息
func (r *WebViewsRepository) RegeditUser(regeditDTO dto.RegeditDTO, session core.Session) (bool, string) {
	user := models.User{
		Username: regeditDTO.Username,
		Email:    regeditDTO.Email,
		Password: regeditDTO.Password,
		Nickname: regeditDTO.Username,
		Phone:    regeditDTO.Phone,
	}

	if err := r.Q.User.Create(&user); err != nil {
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
			if strings.HasSuffix(err.Error(), "users.email") {
				return false, "邮箱已存在，请重试！"
			} else if strings.HasSuffix(err.Error(), "users.username") {
				return false, "用户名已存在，请重试！"
			} else if strings.HasSuffix(err.Error(), "users.phone") {
				return false, "手机号已存在，请重试！"
			}
		}
		if strings.HasPrefix(err.Error(), "Error 1062") {
			flag := ""
			if strings.HasSuffix(err.Error(), "_username'") {
				flag = "用户名"
			} else if strings.HasSuffix(err.Error(), "_email'") {
				flag = "邮箱"
			}
			log.Error().Err(err).Msgf("用户注册失败，%s已存在", flag)
			return false, fmt.Sprintf("%s[%s]已存在，请重试！", flag, strings.Split(strings.Split(err.Error(), "entry '")[1], "' for")[0])
		}
		log.Error().Err(err).Msg("用户注册失败")
		return false, "用户注册失败！"
	}
	userToken, err := session.SetUserSession(user)
	if err != nil {
		return false, "用户登录失败！"
	}

	if err := r.Q.Token.Create(&models.Token{UserID: user.ID, Token: userToken, Type: "access", Revoked: true}); err != nil {
		return false, "用户登录失败！"
	}

	return true, userToken
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
func (r *WebViewsRepository) UserLogin(loginDTO dto.LoginDTO, session core.Session) (bool, string) {
	user, err := r.Q.User.Where(r.Q.User.Where(r.Q.User.Username.Eq(loginDTO.Username)).Or(r.Q.User.Email.Eq(loginDTO.Username))).
		Where(r.Q.User.Password.Eq(pwd.Encryption(loginDTO.Password))).Where(r.Q.User.Status.Eq(status.Enable.Int())).First()
	if err != nil {
		return false, "用户名或密码错误！"
	}

	fmt.Println(user)
	return true, "用户登录成功！"
}
