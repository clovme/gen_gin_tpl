package auth

import (
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

/*
| 字段         | 说明       |
| ----------- | ---------- |
| ID          | 主键   |
| UserID      | 用户ID |
| RoleID      | 角色ID |
| ExpireAt    | 角色到期时间 |
| IsActive    | 是否启用 |
| Description | 描述 |
| CreatedAt   | 备注 |
*/

type UserRole struct {
	ID          int64         `gorm:"primaryKey;type:bigint"`
	UserID      int64         `gorm:"type:bigint;not null;index"` // 用户ID
	RoleID      int64         `gorm:"type:bigint;not null;index"` // 角色ID
	ExpireAt    *time.Time    `gorm:""`                           // 角色到期时间
	Status      status.Status `gorm:"type:int;default:1"`         // 状态：Enable启用，Disable禁用，其他扩展(如审核中，待发布等)
	Description string        `gorm:"type:varchar(255)"`
	CreatedAt   time.Time     `gorm:"autoCreateTime:nano"`
	DeletedAt   *time.Time    `gorm:"index"`
}

func (r *UserRole) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}

func (r *UserRole) TableName() string {
	return "sys_user_role"
}
