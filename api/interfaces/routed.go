package interfaces

import (
	"github.com/gin-gonic/gin"
)

type Routed interface {
	Path() string
	GetFullPath() string
	SetFullPath(string)
	SetApiGroup(*gin.RouterGroup)
	GetApiGroup() *gin.RouterGroup
	LoadRoutes() []Routed
	SetHandler(handler func(ctx *gin.Context))
	GetHandler() func(ctx *gin.Context)
	Get(*gin.Context)
	Post(*gin.Context)
	Put(*gin.Context)
	Patch(*gin.Context)
	Delete(*gin.Context)
}
