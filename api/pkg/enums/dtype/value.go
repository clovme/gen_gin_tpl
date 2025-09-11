package dtype

import (
	"gen_gin_tpl/pkg/enums"
	"sort"
)

type DType int

const Name = "value_type"

const (
	Int DType = iota + 1
	Int64
	Float64
	String
	Bool
	Time
)

var (
	initiate = map[DType]enums.Enums{
		Int:     {ID: 1000000000000000000, Key: "int", Name: "int", Desc: "整数类型"},
		Int64:   {ID: 1000000000000000001, Key: "bigint", Name: "int64", Desc: "长整型"},
		Float64: {ID: 1000000000000000002, Key: "double", Name: "float64", Desc: "浮点型"},
		String:  {ID: 1000000000000000003, Key: "varchar", Name: "string", Desc: "字符串类型"},
		Bool:    {ID: 1000000000000000004, Key: "tinyint(1)", Name: "bool", Desc: "布尔值"},
		Time:    {ID: 1000000000000000005, Key: "datetime", Name: "time.Time", Desc: "时间类型"},
	}

	enumToValue = make(map[string]DType)
)

func init() {
	for code, meta := range initiate {
		enumToValue[meta.Key] = code
	}
}

// ID 获取enums.ID
func (c DType) ID() int64 {
	return initiate[c].ID
}

// Key 获取enums.Key
func (c DType) Key() string {
	if meta, ok := initiate[c]; ok {
		return meta.Key
	}
	return "Int"
}

// Name 获取枚举名称
func (c DType) Name() string {
	if meta, ok := initiate[c]; ok {
		return meta.Name
	}
	return "Int"
}

// Desc 获取枚举描述
func (c DType) Desc() string {
	if meta, ok := initiate[c]; ok {
		return meta.Desc
	}
	return "Int"
}

// Int 获取枚举值
func (c DType) Int() int {
	return int(c)
}

// Is 比较枚举值
func (c DType) Is(v DType) bool {
	return v == c
}

// Code 获取ValueT
func Code(key string) DType {
	if enum, ok := enumToValue[key]; ok {
		return enum
	}
	return Int
}

// Values 获取所有枚举
func Values() []DType {
	values := make([]DType, 0, len(initiate))
	for k := range initiate {
		values = append(values, k)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return values
}
