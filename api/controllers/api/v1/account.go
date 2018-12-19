package v1

import (
	"github.com/gin-gonic/gin"
	"gop/api/controllers/api/v1/account"
	"gop/api/interfaces"
	"gop/api/models/route"
)

type AccountController struct {
	route.Route
}

func (r *AccountController) LoadRoutes() []interfaces.Routed {
	return []interfaces.Routed{
		new(account.AccountMy),
		new(account.AccountBalance),
	}
}

func (r *AccountController) Path() string {
	return "/account"
}

func (r *AccountController) Get(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": r.GetFullPath(),
	})
}
