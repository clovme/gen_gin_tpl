package public

import (
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/enums/gender"
	"gen_gin_tpl/pkg/enums/status"
	"gorm.io/gorm"
	"time"
)

type ApiUserVO struct {
	ID          int64          `json:"id"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	Phone       string         `json:"phone"`
	Password    string         `json:"password"`
	Nickname    string         `json:"nickname"`
	Avatar      string         `json:"avatar"`
	Gender      gender.Gender  `json:"gender"`
	Birthday    *time.Time     `json:"birthday"`
	Status      status.Status  `json:"status"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}

func ToVO(user models.User) *ApiUserVO {
	return &ApiUserVO{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Phone:       user.Phone,
		Password:    user.Password,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
		Gender:      user.Gender,
		Birthday:    user.Birthday,
		Status:      user.Status,
		Description: user.Description,
		CreatedAt:   *user.CreatedAt,
		UpdatedAt:   *user.UpdatedAt,
		DeletedAt:   user.DeletedAt,
	}
}
