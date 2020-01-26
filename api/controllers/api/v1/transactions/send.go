package transactions

import (
	"github.com/gin-gonic/gin"
	"source.cloud.google.com/onemo-api//api/models/route"
)

type TransactionSend struct {
	route.Route
}

func (r *TransactionSend) Path() string {
	return "/send/:id"
}

func (r *TransactionSend) Get(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": r.GetFullPath(),
	})
}
