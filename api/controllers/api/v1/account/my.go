package account

import (
	"github.com/gin-gonic/gin"
	"gop/api/models/route"
)

type AccountMy struct {
	route.Route
}

func (r *AccountMy) Path() string {
	return "/my"
}

func (r *AccountMy) Get(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": r.GetFullPath(),
	})
}
