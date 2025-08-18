package models

import (
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// CorsWhitelist Cors跨域白名单表
type CorsWhitelist struct {
	ID          int64          `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:跨域白名单ID，主键"`
	Origin      string         `gorm:"type:varchar(255);uniqueIndex;not null;comment:跨域白名单"`
	Description string         `gorm:"type:varchar(255);comment:跨域白名单描述"`
	Status      status.Status  `gorm:"type:int;default:1;comment:状态"`
	CreatedAt   *time.Time     `gorm:"autoCreateTime:nano;comment:创建时间"`
	UpdatedAt   *time.Time     `gorm:"autoUpdateTime:nano;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"embedded;comment:软删除标记，空值表示未删除"`
}

func (r *CorsWhitelist) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}

func (r *CorsWhitelist) TableName() string {
	return "sys_cors_whitelist"
}

func (r *CorsWhitelist) TableComment() string {
	return "Cors跨域白名单表"
}
