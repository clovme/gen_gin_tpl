package views

import (
	"gen_gin_tpl/pkg/enums/status"
	"time"
)

type LoginUserVO struct {
	ID          int64         `json:"id"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Phone       string        `json:"phone"`
	Password    string        `json:"password"`
	Nickname    string        `json:"nickname"`
	Avatar      string        `json:"avatar"`
	Gender      int           `json:"gender"`
	Birthday    *time.Time    `json:"birthday"`
	Status      status.Status `json:"status"`
	Description string        `json:"description"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	DeletedAt   *time.Time    `json:"deletedAt"`
}
