package auth

import (
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// RolePermission 角色权限表
type RolePermission struct {
	ID           int64          `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:角色权限ID，主键"`
	RoleID       int64          `gorm:"type:bigint;not null;index;comment:角色ID"`
	PermissionID int64          `gorm:"type:bigint;not null;index;comment:权限ID"`
	CreatedAt    *time.Time     `gorm:"autoCreateTime:nano;comment:创建时间"`
	Status       status.Status  `gorm:"type:int;default:1;comment:状态：Enable启用，Disable禁用，其他扩展(如审核中，待发布等)"`
	DeletedAt    gorm.DeletedAt `gorm:"embedded;comment:软删除标记，空值表示未删除"`
}

func (r *RolePermission) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}

func (r *RolePermission) TableName() string {
	return "sys_role_permission"
}

func (r *RolePermission) TableComment() string {
	return "角色权限表"
}
