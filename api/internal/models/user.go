package models

import (
	"gen_gin_tpl/pkg/crypto"
	"gen_gin_tpl/pkg/enums/gender"
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// User 用户表
type User struct {
	ID          int64          `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:用户ID，主键"`
	Username    string         `gorm:"type:varchar(50);uniqueIndex;not null;comment:用户名，唯一"`
	Email       string         `gorm:"type:varchar(50);uniqueIndex;not null;comment:邮箱，唯一且必须"`
	Phone       string         `gorm:"type:varchar(20);uniqueIndex;comment:电话，唯一但可以为空"`
	Password    string         `gorm:"type:varchar(50);not null;comment:密码，别json序列化"`
	Nickname    string         `gorm:"type:varchar(50);comment:昵称，非必填"`
	Avatar      string         `gorm:"type:varchar(50);comment:头像URL"`
	Gender      gender.Gender  `gorm:"type:int;default:0;comment:性别 0男 1女 2未知"`
	Birthday    *time.Time     `gorm:"comment:生日"`
	Status      status.Status  `gorm:"type:int;default:1;comment:状态：Enable启用，Disable禁用，其他扩展(如审核中，待发布等)"`
	Description string         `gorm:"type:varchar(255);comment:个人简介、备注"`
	CreatedAt   *time.Time     `gorm:"autoCreateTime:nano;comment:创建时间"`
	UpdatedAt   *time.Time     `gorm:"autoUpdateTime:nano;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"embedded;comment:软删除标记，空值表示未删除"`
}

// BeforeCreate 执行 gorm 创建前操作
func (r *User) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	r.Password = crypto.Encryption(r.Password)
	return
}

// TableComment 表注释
func (r *User) TableComment() string {
	return "用户表"
}
