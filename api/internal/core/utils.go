package core

import (
	"gen_gin_tpl/internal/models"
	"gen_gin_tpl/pkg/constants"
	"github.com/gin-gonic/gin"
)

func getUserInfo(ctx *gin.Context) *models.User {
	if value, exists := ctx.Get(constants.ContextUserInfo); exists {
		return value.(*models.User)
	}
	return nil
}
