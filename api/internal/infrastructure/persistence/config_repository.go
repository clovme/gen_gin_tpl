package persistence

import (
	"gen_gin_tpl/internal/domain/do_config"
	"gen_gin_tpl/internal/infrastructure/query"
	"gorm.io/gorm"
)

type ConfigRepository struct {
	DB *gorm.DB
	Q  *query.Query
}

func (r *ConfigRepository) FindAll() ([]*do_config.Config, error) {
	config, err := r.Q.Config.Find()
	if err != nil {
		return nil, err
	}
	return config, err
}

func (r *ConfigRepository) Save(u *do_config.Config) error {
	return r.DB.Create(u).Error
}
