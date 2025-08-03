package views

import (
	viewsVO "gen_gin_tpl/internal/schema/vo/views"
)

type WebViewsService struct {
	Repo Repository
}

func (r *WebViewsService) FindUserAll() ([]*viewsVO.WebUserVO, error) {
	return r.Repo.FindAll()
}
