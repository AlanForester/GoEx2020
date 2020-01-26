package account

import (
	"github.com/gin-gonic/gin"
	"source.cloud.google.com/onemo-api//api/models/route"
)

type AccountBalance struct {
	route.Route
}

func (r *AccountBalance) Path() string {
	return "/balance"
}

func (r *AccountBalance) Get(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": r.GetFullPath(),
	})
}
