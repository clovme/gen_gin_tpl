package models

import (
	"gen_gin_tpl/pkg/enums/boolean"
	"gen_gin_tpl/pkg/enums/dtype"
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// Config 系统配置项表
type Config struct {
	ID          int64          `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:配置项ID，主键"`
	Name        string         `gorm:"type:varchar(50);not null;unique;comment:配置项名称"`
	Value       string         `gorm:"not null;comment:当前配置值"`
	Default     string         `gorm:"not null;comment:默认配置值"`
	ValueT      dtype.DType    `gorm:"default:0;comment:值类型"`
	Show        boolean.Bool   `gorm:"not null;comment:是否启用"`
	Status      status.Status  `gorm:"type:int;default:1;comment:状态：Enable启用，Disable禁用，其他扩展(如审核中，待发布等)"`
	Description string         `gorm:"type:varchar(255);comment:配置项说明"`
	CreatedAt   *time.Time     `gorm:"autoCreateTime:nano;comment:创建时间"`
	UpdatedAt   *time.Time     `gorm:"autoUpdateTime:nano;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"embedded;comment:软删除标记，空值表示未删除"`
}

func (r *Config) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = utils.GenerateID()
	return
}

func (r *Config) TableName() string {
	return "sys_config"
}

func (r *Config) TableComment() string {
	return "系统配置项表"
}

func (r *Config) Get() bool {
	if r.Value == "" {
		return r.Default == boolean.True.Key()
	}
	return r.Value == boolean.True.Key()
}
