package public

import (
	publicVO "gen_gin_tpl/internal/schema/vo/public"
)

type Repository interface {
	GetAllEnums() ([]*publicVO.ApiEnumsVO, error)
}
