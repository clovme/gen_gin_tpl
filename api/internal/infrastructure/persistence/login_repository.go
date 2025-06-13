package persistence

import (
	"gen_gin_tpl/internal/domain/do_user"
	"gen_gin_tpl/internal/infrastructure/query"
	"gorm.io/gorm"
)

type LoginRepository struct {
	DB *gorm.DB
    Q  *query.Query
}

func (r *LoginRepository) FindAll() ([]*do_user.User, error) {
	data, err := r.Q.User.Find()
    if err != nil {
        return nil, err
    }
    return data, err
}

func (r *LoginRepository) Save(u *do_user.User) error {
	return r.DB.Create(u).Error
}
