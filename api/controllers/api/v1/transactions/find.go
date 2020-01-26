package transactions

import (
	"github.com/gin-gonic/gin"
	"source.cloud.google.com/onemo-api//api/models/route"
)

type TransactionsFind struct {
	route.Route
}


func (r *TransactionsFind) Path() string {
	return "find"
}

func (r *TransactionsFind) Get(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": r.GetFullPath(),
	})
}
