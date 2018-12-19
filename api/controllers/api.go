package controllers

import (
	"github.com/gin-gonic/gin"
	"gop/api/controllers/api"
	"gop/api/interfaces"
	"gop/api/models/route"
)

type ApiController struct {
	route.Route
}

func (r *ApiController) LoadRoutes() []interfaces.Routed {
	return []interfaces.Routed{
		new(api.V1Controller),
	}
}

func (r *ApiController) Path() string {
	return "/api"
}

func (r *ApiController) Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": r.GetFullPath(),
	})
}
