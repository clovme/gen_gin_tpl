package public

import (
	publicVO "gen_gin_tpl/internal/schema/vo/public"
)

type ApiPublicService struct {
	Repo Repository
}

// GetAllEnumsData 获取所有枚举数据
//
// 返回值:
//   - []*publicVO.ApiEnumsVO: 枚举列表
//   - error: 错误信息
func (r *ApiPublicService) GetAllEnumsData() ([]*publicVO.ApiEnumsVO, error) {
	return r.Repo.GetAllEnums()
}
