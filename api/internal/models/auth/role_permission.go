package auth

import (
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

/*
| ID | Name  | Code         | Type | Path           | Method |
| -- | ----- | ------------ | ---- | -------------- | ------ |
| 1  | 用户查询 | user_list   | api  | /api/user      | GET    |
| 2  | 新增用户 | user_create | api  | /api/user      | POST   |
| 3  | 删除用户 | user_delete | api  | /api/user/\:id | DELETE |
| 4  | 菜单管理 | menu_manage | menu | /menu          |        |
*/

type RolePermission struct {
	ID           int64         `gorm:"primaryKey;type:bigint"`
	RoleID       int64         `gorm:"type:bigint;not null;index"`
	PermissionID int64         `gorm:"type:bigint;not null;index"`
	CreatedAt    time.Time     `gorm:"autoCreateTime:nano"`
	Status       status.Status `gorm:"type:int;default:1"` // 状态：Enable启用，Disable禁用，其他扩展(如审核中，待发布等)
	DeletedAt    *time.Time    `gorm:"index"`
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
