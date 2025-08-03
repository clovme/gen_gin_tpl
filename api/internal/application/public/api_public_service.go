package public

import (
	publicVO "gen_gin_tpl/internal/schema/vo/public"
)

type ApiPublicService struct {
	Repo Repository
}

func (r *ApiPublicService) GetAllEnumsData() ([]*publicVO.ApiEnumsVO, error) {
	return r.Repo.GetAllEnums()
}
