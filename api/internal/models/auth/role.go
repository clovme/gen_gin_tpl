package auth

import (
	"gen_gin_tpl/pkg/enums/role"
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// Role 角色表
type Role struct {
	ID          int64          `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:角色ID，主键"`
	Name        string         `gorm:"type:varchar(100);uniqueIndex;not null;comment:角色名称"`
	Type        role.Role      `gorm:"type:int;comment:角色类型"`
	Code        string         `gorm:"type:varchar(64);comment:角色编码（英文唯一）"`
	CreatedBy   int64          `gorm:"not null;comment:创建人ID"`
	Description string         `gorm:"type:varchar(255);comment:角色说明"`
	Status      status.Status  `gorm:"type:int;default:1;comment:状态：Enable启用，Disable禁用，其他扩展(如审核中，待发布等)"`
	RoleGroupID int64          `gorm:"type:bigint;not null;comment:角色组ID"`
	CreatedAt   *time.Time     `gorm:"autoCreateTime:nano;comment:创建时间"`
	UpdatedAt   *time.Time     `gorm:"autoUpdateTime:nano;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"embedded;comment:软删除标记，空值表示未删除*"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}

func (r *Role) TableName() string {
	return "sys_role"
}

func (r *Role) TableComment() string {
	return "角色表"
}
