package application

import (
	"gen_gin_tpl/internal/domain/do_user"
)

type ViewLoginService struct {
	Repo do_user.Repository
}

func (s *ViewLoginService) UserRegeditService() ([]*do_user.User, error) {
	return s.Repo.UserRegeditRepo()
}
