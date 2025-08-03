package persistence

import (
	"gen_gin_tpl/internal/infrastructure/query"
	viewsVO "gen_gin_tpl/internal/schema/vo/views"
	"gorm.io/gorm"
)

type WebViewsRepository struct {
	DB *gorm.DB
	Q  *query.Query
}

func (r *WebViewsRepository) FindAll() ([]*viewsVO.WebUserVO, error) {
	var results []*viewsVO.WebUserVO
	err := r.Q.User.Scan(&results)
	if err != nil {
		return nil, err
	}
	return results, err
}