package application

import (
	"gen_gin_tpl/internal/domain/do_user"
)

type LoginService struct {
	Repo do_user.Repository
}
