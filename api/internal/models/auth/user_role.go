package auth

import (
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// UserRole 用户角色表
type UserRole struct {
	ID          int64          `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:用户角色ID，主键"`
	UserID      int64          `gorm:"type:bigint;not null;index;comment:用户ID"`
	RoleID      int64          `gorm:"type:bigint;not null;index;comment:角色ID"`
	ExpireAt    *time.Time     `gorm:"comment:角色到期时间"`
	Status      status.Status  `gorm:"type:int;default:1;comment:状态：Enable启用，Disable禁用，其他扩展(如审核中，待发布等)"`
	Description string         `gorm:"type:varchar(255);comment:角色描述"`
	CreatedAt   *time.Time     `gorm:"autoCreateTime:nano;comment:创建时间"`
	DeletedAt   gorm.DeletedAt `gorm:"embedded;comment:软删除标记，空值表示未删除"`
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

func (r *UserRole) TableComment() string {
	return "用户角色表"
}
