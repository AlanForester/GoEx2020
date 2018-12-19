package transactions

import (
	"github.com/gin-gonic/gin"
	"gop/api/models/route"
)

type TransactionGet struct {
	route.Route
}


func (r *TransactionGet) Path() string {
	 return "view/:id"
}

func (r *TransactionGet) Get(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": r.GetFullPath(),
	})
}

