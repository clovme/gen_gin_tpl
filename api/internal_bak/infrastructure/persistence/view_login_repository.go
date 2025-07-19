package persistence

import (
	"fmt"
	"gen_gin_tpl/internal/infrastructure/query"
	"gorm.io/gorm"
)

type ViewLoginRepository struct {
	DB *gorm.DB
	Q  *query.Query
}

func (r *ViewLoginRepository) UserRegeditRepository() {
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxx")
}
