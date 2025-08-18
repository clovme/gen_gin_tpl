package auth

import (
	"gen_gin_tpl/pkg/enums/perm"
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// Permission 权限表
type Permission struct {
	ID          int64          `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:权限ID，主键"`
	Name        string         `gorm:"type:varchar(100);uniqueIndex;not null;comment:权限名称，必填，唯一"`
	Code        string         `gorm:"type:varchar(100);uniqueIndex;not null;comment:权限标识（唯一英文编码，建议全小写下划线）"`
	PID         int64          `gorm:"type:bigint;default:0;comment:父级权限ID，0表示顶级节点"`
	Type        perm.Perm      `gorm:"type:int;default:1;comment:权限类型：menu（菜单）/ button（按钮）/ api（接口）"`
	Uri         string         `gorm:"type:varchar(255);comment:路由路径或接口地址，菜单或接口必填"`
	Method      string         `gorm:"type:varchar(20);comment:HTTP请求方式（GET/POST/PUT/DELETE），仅api类型使用"`
	Sort        int            `gorm:"type:int;default:0;comment:排序值，值越大越靠前，默认0"`
	Status      status.Status  `gorm:"type:int;default:1;comment:状态：Enable启用，Disable禁用，其他扩展(如审核中，待发布等)"`
	Description string         `gorm:"type:varchar(255);comment:权限描述，便于备注说明"`
	CreatedAt   *time.Time     `gorm:"autoCreateTime:nano;comment:创建时间"`
	UpdatedAt   *time.Time     `gorm:"autoUpdateTime:nano;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"embedded;comment:软删除标记，空值表示未删除"`
}

func (r *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}

func (r *Permission) TableName() string {
	return "sys_permission"
}

func (r *Permission) TableComment() string {
	return "权限表"
}
