package auth

import (
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// RoleGroup 角色组表
type RoleGroup struct {
	ID          int64          `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:角色组ID，主键"`
	Name        string         `gorm:"type:varchar(64);not null;unique;comment:角色组名称"`
	Description string         `gorm:"type:varchar(255);comment:角色组说明"`
	Status      status.Status  `gorm:"type:int;default:1;comment:状态：Enable启用，Disable禁用，其他扩展(如审核中，待发布等)"`
	CreatedAt   *time.Time     `gorm:"autoCreateTime:nano;comment:创建时间"`
	UpdatedAt   *time.Time     `gorm:"autoUpdateTime:nano;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"embedded;comment:软删除标记，空值表示未删除"`
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

func (r *RoleGroup) TableComment() string {
	return "角色组表"
}
