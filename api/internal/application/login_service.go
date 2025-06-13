package application

import (
	"gen_gin_tpl/internal/domain/do_user"
)

type LoginService struct {
	Repo do_user.Repository
}

func (s *LoginService) GetLogin() ([]*do_user.User, error) {
	return s.Repo.FindAll()
}
