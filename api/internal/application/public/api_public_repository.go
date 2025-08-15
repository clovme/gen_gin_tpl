package public

import (
	publicVO "gen_gin_tpl/internal/schema/vo/public"
)

type Repository interface {
	// GetAllEnums 获取所有枚举
	//
	// 返回值:
	//   - []*publicVO.ApiEnumsVO: 枚举列表
	//   - error: 错误信息
	GetAllEnums() ([]*publicVO.ApiEnumsVO, error)
}
