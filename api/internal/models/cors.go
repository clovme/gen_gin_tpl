package models

import (
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

type CorsWhitelist struct {
	ID          int64          `gorm:"primaryKey;type:bigint"`
	Origin      string         `gorm:"type:varchar(255);uniqueIndex;not null"` // 白名单域名或IP
	Description string         `gorm:"type:varchar(255)"`                      // 描述备注
	Status      status.Status  `gorm:"type:int;default:1"`                     // 状态：Enable启用，Disable禁用，其他扩展(如审核中，待发布等)
	CreatedAt   time.Time      `gorm:"autoCreateTime:nano"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:nano"`
	DeletedAt   gorm.DeletedAt `gorm:"index"` // 软删除
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
