package models

import (
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

/*
| 字段名         | 类型             | 说明                                         |
| :---------- | :------------- | :----------------------------------------- |
| `ID`        | `bigint`       | 主键，自增ID                                    |
| `UserID`    | `bigint`       | 关联的用户ID，外键，可建联合索引提高 token 查询速度             |
| `Token`     | `varchar(512)` | Token字符串，唯一索引，长度预留够（比如 JWT / API Key）      |
| `Type`      | `varchar(50)`  | Token类型，例：`access`、`refresh`、`api`、`admin` |
| `ExpiresAt` | `time.Time`    | 过期时间，做登录态管理、接口Token超时校验必备                  |
| `Revoked`   | `bool`         | 是否吊销（true=作废，false=有效），好处是软作废，安全兜底         |
| `CreatedAt` | `time.Time`    | 创建时间，自动填充                                  |
| `UpdatedAt` | `time.Time`    | 更新时间，自动更新                                  |
*/

type Token struct {
	ID        int64     `gorm:"primaryKey;type:bigint"`
	UserID    int64     `gorm:"not null;index"`                // 关联哪个用户
	Token     string    `gorm:"not null;uniqueIndex;size:512"` // 令牌字符串，通常长点
	Type      string    `gorm:"size:50;default:'access'"`      // 令牌类型，比如 access、refresh、api、admin
	ExpiresAt time.Time `gorm:"index"`                         // 过期时间
	Revoked   bool      `gorm:"default:false"`                 // 是否被吊销
	CreatedAt time.Time `gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano"`
}

func (r *Token) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}
