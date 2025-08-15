package models

import (
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// Token 令牌表
type Token struct {
	ID        int64     `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:令牌ID，主键"`
	UserID    int64     `gorm:"not null;index;comment:关联的用户ID，外键"`
	Token     string    `gorm:"not null;uniqueIndex;size:512;comment:令牌字符串，通常长点"`
	Type      string    `gorm:"size:50;default:'access';comment:令牌类型，比如 access、refresh、api、admin"`
	Revoked   bool      `gorm:"default:false;comment:是否被吊销"`
	CreatedAt time.Time `gorm:"autoCreateTime:nano;comment:创建时间"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano;comment:更新时间"`
}

func (r *Token) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}

func (r *Token) TableComment() string {
	return "Token令牌表"
}
