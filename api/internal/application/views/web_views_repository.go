package views

import (
	viewsVO "gen_gin_tpl/internal/schema/vo/views"
)

type Repository interface {
	FindAll() ([]*viewsVO.WebUserVO, error)
}