package do_config

import (
	"gen_gin_tpl/pkg/enums/em_bool"
	"gen_gin_tpl/pkg/enums/em_status"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

/*
| 字段名          | 类型          | 说明                                           |
| -------------- | ----------- | -------------------------------------------- |
| `ID`           | `bigint`    | 主键，自增ID                                      |
| `Name`         | `string`    | 配置项名称，唯一，不可为空（例如："SiteName"、"MaxUploadSize"） |
| `Value`        | `string`    | 当前配置值（例如："MySite"、"50MB"）                    |
| `Default`      | `string`    | 默认配置值，用于恢复初始值或兜底显示                           |
| `Enable`       | `bool`      | 是否启用该配置，true/false                           |
| `Description`  | `string`    | 配置项的文字说明，便于后台管理者理解含义                         |
| `CreatedAt`    | `time.Time` | 创建时间，自动填充                                    |
| `UpdatedAt`    | `time.Time` | 更新时间，自动更新                                    |
*/

type Config struct {
	ID          int64            `gorm:"primaryKey;type:bigint" json:"id"`
	Name        string           `gorm:"type:varchar(50);not null;unique" json:"name"`
	Value       string           `gorm:"not null" json:"value"`
	Default     string           `gorm:"not null" json:"default"`
	Show        em_bool.Bool     `gorm:"not null" json:"-"`
	Status      em_status.Status `gorm:"type:int;default:1" json:"status"` // 状态：Enable启用，Disable禁用，其他扩展(如审核中，待发布等)
	Description string           `gorm:"type:varchar(255)" json:"description,omitempty"`
	CreatedAt   time.Time        `gorm:"autoCreateTime:nano" json:"createdAt"`
	UpdatedAt   time.Time        `gorm:"autoUpdateTime:nano" json:"updatedAt"`
}

func (config *Config) BeforeCreate(tx *gorm.DB) (err error) {
	config.ID = utils.GenerateID()
	return
}

func (config *Config) TableName() string {
	return "sys_config"
}
