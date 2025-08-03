package auth

import (
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type RoleGroup struct {
	ID          int64         `gorm:"primaryKey;type:bigint"`
	Name        string        `gorm:"type:varchar(64);not null;unique"`
	Description string        `gorm:"type:varchar(255)"`
	Status      status.Status `gorm:"type:int;default:1"` // 状态：Enable启用，Disable禁用，其他扩展(如审核中，待发布等)
	CreatedAt   time.Time     `gorm:"autoCreateTime:nano"`
	UpdatedAt   time.Time     `gorm:"autoUpdateTime:nano"`
}

func (r *RoleGroup) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}

func (r *RoleGroup) TableName() string {
	return "sys_role_group"
}
