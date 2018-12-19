package v1

import (
	"github.com/gin-gonic/gin"
	"gop/api/controllers/api/v1/transactions"
	"gop/api/interfaces"
	"gop/api/models/route"
)

type TransactionsController struct {
	route.Route
}

func (r *TransactionsController) LoadRoutes() []interfaces.Routed {
	return []interfaces.Routed{
		new(transactions.TransactionsFind),
		new(transactions.TransactionGet),
		new(transactions.TransactionSend),
	}
}

func (r *TransactionsController) Path() string {
	return "/transactions"
}

func (r *TransactionsController) Get(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": r.GetFullPath(),
	})
}
