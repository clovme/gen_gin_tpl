package role

import (
	"gen_gin_tpl/pkg/enums"
	"sort"
)

type Role int

const Name = "role"

const (
	System Role = iota
	Admin
	Vip
	SVip
	Custom
)

var (
	initiate = map[Role]enums.Enums{
		System: {ID: 1000000000000000000, Key: "System", Name: "系统管理员", Desc: "拥有系统最高权限，管理平台所有配置与用户操作。"},
		Admin:  {ID: 1000000000000000001, Key: "Admin", Name: "超级管理员", Desc: "具备高级管理权限，负责日常运营与用户管理。"},
		Vip:    {ID: 1000000000000000002, Key: "VIP", Name: "VIP用户", Desc: "尊享高级功能与专属服务的VIP用户。"},
		SVip:   {ID: 1000000000000000003, Key: "SVIP", Name: "SVIP用户", Desc: "享受全部高级特权及优先支持服务的SVIP用户。"},
		Custom: {ID: 1000000000000000004, Key: "Custom", Name: "自定义角色", Desc: "由管理员自定义权限与职责的个性化角色。"},
	}

	enumToValue = make(map[string]Role)
)

func init() {
	for code, meta := range initiate {
		enumToValue[meta.Key] = code
	}
}

// ID 获取enums.ID
func (c Role) ID() int64 {
	return initiate[c].ID
}

// Key 获取enums.Key
func (c Role) Key() string {
	if meta, ok := initiate[c]; ok {
		return meta.Key
	}
	return "Custom"
}

// Name 获取枚举名称
func (c Role) Name() string {
	if meta, ok := initiate[c]; ok {
		return meta.Name
	}
	return "Custom"
}

// Desc 获取枚举描述
func (c Role) Desc() string {
	if meta, ok := initiate[c]; ok {
		return meta.Desc
	}
	return "Custom"
}

// Int 获取枚举值
func (c Role) Int() int {
	return int(c)
}

// Is 比较枚举值
func (c Role) Is(v Role) bool {
	return v == c
}

// Code 获取Role
func Code(key string) Role {
	if enum, ok := enumToValue[key]; ok {
		return enum
	}
	return Custom
}

// Values 获取所有枚举
func Values() []Role {
	values := make([]Role, 0, len(initiate))
	for k := range initiate {
		values = append(values, k)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return values
}
