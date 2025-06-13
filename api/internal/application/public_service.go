package application

import (
	"gen_gin_tpl/internal/domain/do_enums"
)

type PublicService struct {
	Repo do_enums.Repository
}

func (s *PublicService) GetEnums() ([]*do_enums.Enums, error) {
	return s.Repo.FindAll()
}
