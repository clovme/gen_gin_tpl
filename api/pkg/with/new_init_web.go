package with

import (
	"github.com/gin-gonic/gin"
)

func NewWithContextInitWeb(ctx *gin.Context, newController func(c *Context[any])) {
	newController(&Context[any]{
		Context: ctx,
	})
}

func NewWithContextInitWebData[T any](ctx *gin.Context, newController func(c *Context[T])) {
	newController(&Context[T]{
		Context: ctx,
		DTOData: *new(T),
	})
}
