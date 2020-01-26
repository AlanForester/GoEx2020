package models

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func GetJWTUser(c *gin.Context) interface{} {
	if user, ok := c.Get("id"); ok {
		return user
	}
	return nil
}