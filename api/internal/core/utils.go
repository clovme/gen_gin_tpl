package core

import (
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/constants"
)

func getUserInfo(ctx *Context) *models.User {
	if value, exists := ctx.Get(constants.ContextUserInfo); exists {
		return value.(*models.User)
	}
	return nil
}
