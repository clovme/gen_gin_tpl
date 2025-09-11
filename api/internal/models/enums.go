package models

import (
	"gen_gin_tpl/pkg/enums/dtype"
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// Enums 枚举表
type Enums struct {
	ID          int64          `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:枚举项ID，主键"`
	EID         int64          `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:枚举项ID，主键"`
	Category    string         `gorm:"type:varchar(50);index;not null;comment:枚举分类"`
	Key         string         `gorm:"type:varchar(100);not null;comment:枚举键（唯一标识）"`
	Name        string         `gorm:"type:varchar(100);not null;comment:枚举名称（显示用）"`
	Value       int            `gorm:"type:int;not null;comment:枚举值（数字）"`
	ValueT      dtype.DType    `gorm:"default:0;comment:值类型"`
	Sort        int            `gorm:"default:0;comment:排序"`
	Status      status.Status  `gorm:"type:int;default:1;comment:状态"`
	Description string         `gorm:"type:varchar(255);comment:描述"`
	CreatedAt   time.Time      `gorm:"autoCreateTime:nano;comment:创建时间"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:nano;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"embedded;comment:软删除标记，空值表示未删除"`
}

func (r *Enums) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}

func (r *Enums) TableName() string {
	return "sys_enums"
}

func (r *Enums) TableComment() string {
	return "枚举表"
}
