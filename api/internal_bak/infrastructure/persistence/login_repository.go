package persistence

import (
	"gen_gin_tpl/internal/infrastructure/query"
	"gorm.io/gorm"
)

type LoginRepository struct {
	DB *gorm.DB
	Q  *query.Query
}
