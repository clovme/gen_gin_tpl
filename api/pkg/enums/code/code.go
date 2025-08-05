package code

import (
	"gen_gin_tpl/pkg/enums"
	"sort"
)

type ResponseCode int

const Name = "http_status_code"

const (
	// 正常返回
	Success ResponseCode = iota + 10000

	// 业务错误
	VerifyError ResponseCode = iota + 20000 - 1
	CreateError

	// 请求错误
	BadRequest ResponseCode = iota + 30000 - 1
	Unauthorized
	Forbidden
	NotFound
	Unknown

	// 服务器内部错误
	InternalServerError ResponseCode = iota + 40000 - 1
)

var (
	initiate = map[ResponseCode]enums.Enums{
		Success:             {Key: "Success", Name: "成功", Desc: "请求已成功处理！"},
		VerifyError:         {Key: "VerifyError", Name: "验证失败", Desc: "数据验证失败，请检查输入数据！"},
		CreateError:         {Key: "CreateError", Name: "创建失败", Desc: "数据创建失败，请稍后重试！"},
		BadRequest:          {Key: "BadRequest", Name: "错误请求", Desc: "请求参数格式错误或缺失，服务器无法处理！"},
		Unauthorized:        {Key: "Unauthorized", Name: "未认证", Desc: "当前请求需要用户认证或认证已失效！"},
		Forbidden:           {Key: "Forbidden", Name: "禁止访问", Desc: "当前用户无权访问此资源！"},
		NotFound:            {Key: "NotFound", Name: "资源不存在", Desc: "请求的资源不存在或已被删除！"},
		Unknown:             {Key: "Unknown", Name: "未知错误", Desc: "未知错误或异常，请检查请求参数或配置！"},
		InternalServerError: {Key: "InternalServerError", Name: "服务器内部错误", Desc: "服务器开小差了，请稍后再试！"},
	}

	enumToValue = make(map[string]ResponseCode)
)

func init() {
	for code, meta := range initiate {
		enumToValue[meta.Key] = code
	}
}

// Key 获取enums.Key
func (c ResponseCode) Key() string {
	if meta, ok := initiate[c]; ok {
		return meta.Key
	}
	return "Unknown"
}

// Name 获取枚举名称
func (c ResponseCode) Name() string {
	if meta, ok := initiate[c]; ok {
		return meta.Name
	}
	return "未知错误"
}

// Desc 获取枚举描述
func (c ResponseCode) Desc() string {
	if meta, ok := initiate[c]; ok {
		return meta.Desc
	}
	return "未知错误或异常，请检查请求参数或联系管理员"
}

// Enum 获取枚举值
func (c ResponseCode) Enum() int {
	return int(c)
}

// Is 比较枚举值
func (c ResponseCode) Is(v ResponseCode) bool {
	return v == c
}

// Code 获取Code
func Code(key string) ResponseCode {
	if enum, ok := enumToValue[key]; ok {
		return enum
	}
	return Unknown
}

// Values 获取所有枚举
func Values() []ResponseCode {
	values := make([]ResponseCode, 0, len(initiate))
	for k := range initiate {
		values = append(values, k)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return values
}
