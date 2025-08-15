package with

import (
	"gen_gin_tpl/internal/models"
	"github.com/gin-gonic/gin"
)

type Context[T any] struct {
	*gin.Context
	DTOData            T
	isContextEncrypted bool
	IsLogin            bool
	Session            Session
	UserInfo           *models.User
}
