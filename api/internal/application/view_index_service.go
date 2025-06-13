package application

import (
	"gen_gin_tpl/internal/domain/do_user"
)

type ViewIndexService struct {
	Repo do_user.Repository
}

func (s *ViewIndexService) GetViewIndex() ([]*do_user.User, error) {
	return s.Repo.FindAll()
}
