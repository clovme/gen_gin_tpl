package models

import (
	"gen_gin_tpl/pkg/enums/status"
	"gen_gin_tpl/pkg/pwd"
	"gen_gin_tpl/pkg/utils"
	"gorm.io/gorm"
	"time"
)

/*
| 字段名           | 类型             | 说明                            |
| :------------ | :------------- | :---------------------------- |
| `ID`          | `bigint`       | 主键，自增ID                       |
| `Username`    | `varchar(50)`  | 用户名，唯一且不能为空，常用作登录名            |
| `Email`       | `varchar(100)` | 邮箱地址，唯一且必填                    |
| `Phone`       | `varchar(20)`  | 手机号，唯一但可以为空                   |
| `Password`    | `varchar(255)` | 密码哈希值，安全起见，JSON序列化时不返回        |
| `Nickname`    | `varchar(50)`  | 昵称，非必填                        |
| `Avatar`      | `varchar(255)` | 头像图片URL                       |
| `Gender`      | `int`          | 性别，0=未知，1=男，2=女               |
| `Birthday`    | `*time.Time`   | 生日，允许空，指针类型                   |
| `Status`      | `int`          | 用户状态，1=启用，0=禁用                |
| `Description` | `varchar(255)` | 个人简介、备注                       |
| `CreatedAt`   | `time.Time`    | 创建时间，自动填充                     |
| `UpdatedAt`   | `time.Time`    | 更新时间，自动更新                     |
| `DeletedAt`   | `*time.Time`   | 软删除字段，配合 `gorm:"index"` 实现软删除 |
*/

type User struct {
	ID          int64         `gorm:"primaryKey;type:bigint"`
	Username    string        `gorm:"type:varchar(50);uniqueIndex;not null"` // 用户名，唯一
	Email       string        `gorm:"type:varchar(50);uniqueIndex;not null"` // 邮箱，唯一且必须
	Phone       string        `gorm:"type:varchar(20);uniqueIndex"`          // 电话，可以为空，唯一
	Password    string        `gorm:"type:varchar(50);not null"`             // 密码哈希，别json序列化
	Nickname    string        `gorm:"type:varchar(50)"`                      // 昵称，非必填
	Avatar      string        `gorm:"type:varchar(50)"`                      // 头像URL
	Gender      int           `gorm:"type:int;default:0"`                    // 性别 0未知 1男 2女
	Birthday    *time.Time    `json:"birthday,omitempty"`                    // 生日，指针，允许空
	Status      status.Status `gorm:"type:int;default:1"`                    // 状态：Enable启用，Disable禁用，其他扩展(如审核中，待发布等)
	Description string        `gorm:"type:varchar(255)"`
	CreatedAt   time.Time     `gorm:"autoCreateTime:nano"`
	UpdatedAt   time.Time     `gorm:"autoUpdateTime:nano"`
	DeletedAt   *time.Time    `gorm:"index"` // 软删除
}

func (r *User) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	r.Password = pwd.Encryption(r.Password)
	return
}
