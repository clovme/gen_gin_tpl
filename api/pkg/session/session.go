package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Set(c *gin.Context, key string, value any) error {
	s := sessions.Default(c)
	s.Set(key, value)
	return s.Save()
}

func Get(c *gin.Context, key string) any {
	s := sessions.Default(c)
	return s.Get(key)
}
