package api

import (
	"github.com/gin-gonic/gin"
	"gop/api/controllers/api/v1"
	"gop/api/interfaces"
	"gop/api/models/route"
)

type V1Controller struct {
	route.Route
}

func (r *V1Controller) LoadRoutes() []interfaces.Routed {
	return []interfaces.Routed{
		new(v1.AccountController),
		new(v1.TransactionsController),
	}
}


func (r *V1Controller) Path() string {
	return "/v1"
}

func (r *V1Controller) Get(ctx *gin.Context)   {
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": r.GetFullPath(),
	})
}
