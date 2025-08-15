package persistence

import (
	"gen_gin_tpl/internal/infrastructure/query"
	publicVO "gen_gin_tpl/internal/schema/vo/public"
	"gorm.io/gorm"
)

type ApiPublicRepository struct {
	DB *gorm.DB
	Q  *query.Query
}

// GetAllEnums 获取所有枚举
//
// 返回值:
//   - []*publicVO.ApiEnumsVO: 枚举列表
//   - error: 错误信息
func (r *ApiPublicRepository) GetAllEnums() ([]*publicVO.ApiEnumsVO, error) {
	var results []*publicVO.ApiEnumsVO
	err := r.Q.Enums.Scan(&results)
	if err != nil {
		return nil, err
	}
	return results, err
}
