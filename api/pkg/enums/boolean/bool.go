package boolean

import (
	"gen_gin_tpl/pkg/enums"
	"sort"
)

type Bool int

const Name = "bool"

const (
	False Bool = iota
	True
)

var (
	initiate = map[Bool]enums.Enums{
		True:  {ID: 1000000000000000000, Key: "true", Name: "True", Desc: "布尔值true"},
		False: {ID: 1000000000000000001, Key: "false", Name: "False", Desc: "布尔值false"},
	}

	enumToValue = make(map[string]Bool)
)

func init() {
	for code, meta := range initiate {
		enumToValue[meta.Key] = code
	}
}

// ID 获取enums.ID
func (c Bool) ID() int64 {
	return initiate[c].ID
}

// Key 获取enums.Key
func (c Bool) Key() string {
	return initiate[c].Key
}

// Name 获取枚举名称
func (c Bool) Name() string {
	return initiate[c].Name
}

// Desc 获取枚举描述
func (c Bool) Desc() string {
	return initiate[c].Desc
}

// Int 获取枚举值
func (c Bool) Int() int {
	return int(c)
}

// Is 比较枚举值
func (c Bool) Is(v Bool) bool {
	return v == c
}

// Code 获取Bool
func Code(key string) Bool {
	if enum, ok := enumToValue[key]; ok {
		return enum
	}
	return False
}

// Values 获取所有枚举
func Values() []Bool {
	values := make([]Bool, 0, len(initiate))
	for k := range initiate {
		values = append(values, k)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return values
}
